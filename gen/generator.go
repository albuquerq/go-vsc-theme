package gen

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/nsalb/go-vsc-theme/manifest"
)

// GenPackage gera um pacote de tema para o Visual Studio Code.
func GenPackage(displayName, filePath, destinPath, baseColor, fileOpt string) error {

	opt, err := manifest.NewPackageOptionFromFile(fileOpt)
	if err != nil {
		log.Printf("Erro ao ler o arquivo %s. Utilizado PackageOption Default", fileOpt)
	}

	pkgManifest := manifest.NewPackageManifest(displayName, opt)
	fileName := path.Base(filePath)

	switch baseColor {
	case "dark":
		pkgManifest.Contributes.AddTheme(*manifest.NewThemeDark(
			pkgManifest.DisplayName,
			"./"+path.Join("themes", fileName),
		))
	case "light":
		pkgManifest.Contributes.AddTheme(*manifest.NewThemeLight(
			pkgManifest.DisplayName,
			"./"+path.Join("themes", fileName),
		))
	}

	log.Printf("Gerando o pacote %s...", pkgManifest.Name)

	pakagePath := path.Join(destinPath, pkgManifest.Name)

	err = os.MkdirAll(path.Join(pakagePath, "themes"), os.ModeDir|os.ModePerm)
	if err != nil {
		log.Fatalf("[falha] %s\n", err.Error())
		return err
	}

	data, err := json.MarshalIndent(pkgManifest, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(path.Join(pakagePath, "package.json"), data, os.ModePerm)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo package.json %s", err.Error())
		return err
	}

	wdir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Erro ao recuperar o work dir.")
		return err
	}

	t, err := template.ParseFiles(path.Join(wdir, "templates", "README.tmpl"))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	b := bytes.Buffer{}
	err = t.ExecuteTemplate(&b, "README.tmpl", pkgManifest)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = ioutil.WriteFile(path.Join(pakagePath, "README.md"), b.Bytes(), os.ModePerm|os.ModeDir)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = os.Rename(filePath, path.Join(pakagePath, "themes", fileName))
	if err != nil {
		log.Fatalf("Erro ao mover o arquivo %s. %s\n", fileName, err.Error())
		return err
	}

	return nil
}
