package house

import "strings"

var verb = []string{
	"",
	"that lay in ",
	"that ate ",
	"that killed ",
	"that worried ",
	"that tossed ",
	"that milked ",
	"that kissed ",
	"that married ",
	"that woke ",
	"that kept ",
	"that belonged to ",
	"",
}

var noun = []string{
	"",
	"the house that Jack built",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

func Verse(n int) string {
	if n == 1 {
		return "This is " + noun[n] + "."
	}
	verse := "This is " + noun[n] + "\n"
	verse += strings.Replace(Verse(n-1), "This is ", verb[n-1], 1)
	return verse
}

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += "\n\n" + Verse(i)
	}
	return strings.TrimSpace(song)
}
