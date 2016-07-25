package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/albuquerq/go-vsc-theme/fmtname"
	"github.com/albuquerq/go-vsc-theme/gen"
)

var (
	// Diretório com arquivos .tmTheme
	filesDir string
	// Cor base para os temas
	isDarkBased bool
	// Diretório destino
	outDir string
	// False para gerar o DisplayName automaticamente
	autoNaming bool
	// Caminho para o arquivo JSON de opções
	fileOpt string
)

func init() {
	flag.StringVar(&filesDir, "in", ".", "Directory containing files \".tmTheme\".")
	flag.StringVar(&outDir, "out", ".", "Destination directory for packages.")
	flag.StringVar(&fileOpt, "opt", "./opt.json", "Path to JSON options file.")
	flag.BoolVar(&autoNaming, "auto", true, "Set to false if you want to enter the DisplayName manually.")
	flag.BoolVar(&isDarkBased, "dark", false, "Set true for dark theme files.")

	flag.Parse()
}

func main() {
	var baseColor = "light"

	if isDarkBased {
		baseColor = "dark"
	}

	var count = 0

	filepath.Walk(filesDir, func(currentDir string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Print(err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		if ok, _ := path.Match("*.tmTheme", info.Name()); ok {

			var displayName = fmtname.Normalize(strings.Replace(info.Name(), ".tmTheme", "", 1))

			if autoNaming {
				fmt.Printf("Building the package for %s. Display name \"%s\"   ", info.Name(), displayName)
			} else {
				var novoNome = ""

				fmt.Printf("Building the package for %s. Enter the displayName in quotes, or use the default (\"%s\")   ", info.Name(), displayName)
				fmt.Fscanf(os.Stdin, "%q\n", &novoNome)

				if novoNome != "" {
					displayName = novoNome
				}
			}

			err := gen.GenPackage(displayName, currentDir, outDir, baseColor, fileOpt)
			if err == nil {
				fmt.Println("Done")
				count++
			} else {
				fmt.Println("Failed")
			}
		}
		return nil
	})

	if count == 0 {
		fmt.Printf("Not found tmThemes files in \"%s\". Use --help flag for more info.\n", filesDir)
		return
	}

	fmt.Printf("Created %d theme packages.\nDone", count)
}
