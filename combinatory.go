package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt.go"
)

const version = "combinatory 1.0.0"
const usage = `combinatory.

Usage:
  combinatory worder [--entry "word" | -e "word"...] [--howmany] [--save-to-file] [--basedir path]
  combinatory -h | --help
  combinatory --version

Options:
  -e --entry <string>        Word to be combined.
  --howmany                  Just print how many combinations are there.
  --save-to-file             Will save to a file called <entry>.txt.
  --basedir <string>         Path where to save files. Defaults to current dir.
  -h --help                  Show this screen.
  --version                  Show version.
`

func dict() (base map[string][]string) {
	base = make(map[string][]string)

	base["a"] = []string{"a", "A", "4", "$", "á", "Á", "ä", "Ä", "à", "À", "ã", "Ã"}
	base["á"] = base["a"]
	base["à"] = base["a"]
	base["ä"] = base["a"]
	base["ã"] = base["a"]
	base["b"] = []string{"b", "B", "8", "*"}
	base["c"] = []string{"c", "C"}
	base["d"] = []string{"d", "D"}
	base["e"] = []string{"e", "E", "3", "#", "é", "É", "è", "È", "ë", "Ë"}
	base["é"] = base["e"]
	base["è"] = base["e"]
	base["ë"] = base["e"]
	base["f"] = []string{"f", "F"}
	base["g"] = []string{"g", "G", "9", "(", "6", "^"}
	base["h"] = []string{"h", "H"}
	base["i"] = []string{"i", "I", "l", "1", "!", "í", "Í", "ï", "Ï", "ì", "Ì", "|"}
	base["í"] = base["i"]
	base["ï"] = base["i"]
	base["ì"] = base["i"]
	base["j"] = []string{"j", "J"}
	base["k"] = []string{"k", "K"}
	base["l"] = []string{"l", "L", "1", "I", "i", "!", "|"}
	base["m"] = []string{"m", "M"}
	base["n"] = []string{"n", "N"}
	base["o"] = []string{"o", "O", "0", ")", "#", "ó", "Ó", "ò", "Ò", "ö", "Ö", "õ", "Õ"}
	base["ó"] = base["o"]
	base["ò"] = base["o"]
	base["ö"] = base["o"]
	base["õ"] = base["o"]
	base["p"] = []string{"p", "P"}
	base["q"] = []string{"q", "Q"}
	base["r"] = []string{"r", "R"}
	base["s"] = []string{"s", "S", "5", "%"}
	base["t"] = []string{"t", "T", "7", "&"}
	base["u"] = []string{"u", "U", "V", "v", "ú", "Ú", "ù", "Ù", "ü", "Ü"}
	base["ú"] = base["u"]
	base["ù"] = base["u"]
	base["ü"] = base["u"]
	base["v"] = []string{"v", "V", "u", "U"}
	base["w"] = []string{"w", "W", "VV", "vv", "Vv", "vV", "UU", "uu", "Uu", "uU"}
	base["x"] = []string{"x", "X"}
	base["y"] = []string{"y", "Y"}
	base["z"] = []string{"z", "Z"}
	base["1"] = []string{"1", "!", "l", "L", "I", "i"}
	base["2"] = []string{"2", "@"}
	base["3"] = []string{"3", "#"}
	base["4"] = []string{"4", "$"}
	base["5"] = []string{"5", "%"}
	base["6"] = []string{"6", "^"}
	base["7"] = []string{"7", "&"}
	base["8"] = []string{"8", "*"}
	base["9"] = []string{"9", "("}
	base["0"] = []string{"0", ")"}
	base["!"] = []string{"!", "1"}
	base["@"] = []string{"@", "2"}
	base["#"] = []string{"#", "3"}
	base["$"] = []string{"$", "4"}
	base["%"] = []string{"%", "5"}
	base["^"] = []string{"^", "6"}
	base["&"] = []string{"&", "7"}
	base["*"] = []string{"*", "8"}
	base["("] = []string{"(", "9"}
	base[")"] = []string{")", "0"}
	base["-"] = []string{"-", "_"}
	base["_"] = []string{"_", "-"}
	base["="] = []string{"=", "+"}
	base["+"] = []string{"+", "="}
	base["`"] = []string{"`", "~"}
	base["~"] = []string{"~", "`"}
	base["\""] = []string{"\"", "'"}
	base["'"] = base["\""]
	base["["] = []string{"[", "{"}
	base["]"] = []string{"]", "}"}
	base["{"] = base["["]
	base["}"] = base["]"]
	base[":"] = []string{":", ";"}
	base[";"] = base[":"]
	base["\\"] = []string{"\\", "|"}
	base["|"] = base["\\"]
	base["<"] = []string{"<", ","}
	base[","] = base["<"]
	base[">"] = []string{">", "."}
	base["."] = base[">"]
	base["/"] = []string{"/", "?"}
	base["?"] = base["/"]

	return base
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func combine_word(base map[string][]string, word string, sword []string, position int) (combination []string) {
	current := sword[position]

	for i := 0; i < len(base[current]); i++ {
		nrune := []rune(base[current][i])
		nword := replaceAtIndex(word, nrune[0], position)
		combination = append(combination, nword)
	}

	if position+1 < len(sword) {
		for _, dword := range combination {
			combination = append(combination, combine_word(base, dword, sword, position+1)...)
		}
	}

	return combination
}

func worder_printer(args map[string]interface{}, word string, combination []string) {
	if args["--save-to-file"].(bool) {
		baseDir, _ := os.Getwd()

		if args["--basedir"] != nil {
			baseDir = args["--basedir"].(string)
		}

		f, err := os.Create(baseDir + "/" + word + ".txt")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		for _, dword := range combination {
			f.WriteString(dword + "\n")
		}
	} else {
		for _, dword := range combination {
			fmt.Println(dword)
		}
	}
}

func worder(args map[string]interface{}) {
	base := dict()

	words := args["--entry"].([]string)
	combinations := make(map[string][]string)

	for _, word := range words {
		sword := strings.Split(word, "")
		combinations[word] = combine_word(base, word, sword, 0)
		if args["--howmany"].(bool) {
			fmt.Println(len(combinations[word]))
		} else {
			worder_printer(args, word, combinations[word])
		}
	}
}

func main() {
	args, _ := docopt.Parse(usage, nil, true, version, false)

	if args["worder"] != nil {
		worder(args)
	}
}
