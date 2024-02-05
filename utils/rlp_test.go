package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeRlphex(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"f90110b842307863616262326438313236636665623633663138343431396537343438303234336638646664343135633664313234303632303732336561373133633232623461b842307861663666313032653031663162353462366131363131386466313432386361363036373962366531323237643363643765616233306466316434643731323630b842307833316432333934306165613837353662353932646637636336366138346361393132393263633362666266323261653436343830633663366231633132653662b842307833343937393237343534626139663662666464333831393866663133656338663039633263656534633231356431393937353230656436613133343336616166", []string{"0xcabb2d8126cfeb63f184419e74480243f8dfd415c6d1240620723ea713c22b4a", "0xaf6f102e01f1b54b6a16118df1428ca60679b6e1227d3cd7eab30df1d4d71260", "0x31d23940aea8756b592df7cc66a84ca91292cc3bfbf22ae46480c6c6b1c12e6b", "0x3497927454ba9f6bfdd38198ff13ec8f09c2cee4c215d1997520ed6a13436aaf"}},
		{"f8ccb842307838343465363562616434616236383761393531316431343134633332643163633339333539373565383338643033643265396335313561353066643933633763b842307833313739336437623833323131396165333732646639343437303233346266353134633635356661303739623931373037363463656665373365613635653264b842307833663337393830396135353866343233313031323466623835396166646564356133373933636430396139633235653135643938656332623738393361656332", []string{"0x844e65bad4ab687a9511d1414c32d1cc3935975e838d03d2e9c515a50fd93c7c", "0x31793d7b832119ae372df94470234bf514c655fa079b9170764cefe73ea65e2d", "0x3f379809a558f42310124fb859afded5a3793cd09a9c25e15d98ec2b7893aec2"}},
	}

	for _, test := range tests {
		result, err := DecodeRlphex(test.input)
		assert.NoError(t, err, "DecodeRlphex(%s) returned an error", test.input)
		assert.Equal(t, test.expected, result, "DecodeRlphex(%s) returned unexpected result", test.input)
	}
}

func TestDecodeRlphexError(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"qqq"},
	}

	for _, test := range tests {
		_, err := DecodeRlphex(test.input)
		assert.Error(t, err, "DecodeRlphex(%s) returned an error", test.input)
	}
}
