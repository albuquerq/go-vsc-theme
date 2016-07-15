package fmtname

import (
	"os"
	"testing"
)

type Amostra struct {
	In  string
	Out string
}

func TestSplitCamelCase(t *testing.T) {

	amostras := []Amostra{
		{"NatanDeSouzaAlbuquerque", "Natan De Souza Albuquerque"},
		{"Active4D", "Active4D"},
		{"Chrome DevTools", "Chrome Dev Tools"},
		{"CodeWarrior", "Code Warrior"},
		{"GitHubCleanWhite", "Git Hub Clean White"},
	}

	for _, amostra := range amostras {
		if result := SplitCamelCase(amostra.In); result != amostra.Out {
			t.Errorf("Diverge %s -> %s\n", amostra.Out, result)
		}
	}
}

func TestNormalize(t *testing.T) {
	amostras := []Amostra{
		{"Active-4D", "Active 4D"},
		{"Chrome_DevTools", "Chrome Dev Tools"},
		{"Code-Warrior", "Code Warrior"},
	}

	for _, amostra := range amostras {
		if result := Normalize(amostra.In); result != amostra.Out {
			t.Errorf("Diverge %s -> %s\n", amostra.Out, result)
		}
	}
}

func TestToFolderName(t *testing.T) {
	amostras := []Amostra{
		{"Active4D", "active4d"},
		{"Chrome Dev Tools", "chromedevtools"},
		{"Code Warrior", "codewarrior"},
		{"3024 Night", "3024night"},
	}

	for _, amostra := range amostras {
		if result := ToLowerJoin(amostra.In); result != amostra.Out {
			t.Errorf("Diverge %s -> %s\n", amostra.Out, result)
		} else {
			t.Logf("theme-%s\n", result)
		}

	}
}

func TestMkdirAll(t *testing.T) {
	err := os.MkdirAll("./teste/oi/kkk", os.ModeDir|os.ModePerm)
	if err != nil {
		t.Error(err)
	}
}
