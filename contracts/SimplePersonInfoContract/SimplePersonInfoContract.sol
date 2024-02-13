// SimplePersonInfoContract.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimplePersonInfoContract {
    // Structure to represent a person
    struct Person {
        string name;
        uint256 age;
    }

    // Array to store person information
    Person[] public persons;

    // Event to log when person information is updated
    event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge);

    // Function to set person information
    function setPersonInfo(string memory _name, uint256 _age) public {
        require(bytes(_name).length > 0, "Name should not be empty");
        require(_age > 0, "Age should be greater than 0");

        // Create a new person
        Person memory newPerson = Person(_name, _age);

        // Add the new person to the array
        persons.push(newPerson);

        // Emit an event to log the update
        emit PersonInfoUpdated(persons.length - 1, _name, _age);
    }

    // Function to get person information for a specific index
    function getPersonInfo(uint256 _personIndex) public view returns (string memory, uint256) {
        require(_personIndex < persons.length, "Person index out of bounds");

        Person memory person = persons[_personIndex];
        return (person.name, person.age);
    }

    // Function to get the total number of persons
    function getPersonsCount() public view returns (uint256) {
        return persons.length;
    }
}