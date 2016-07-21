package manifest

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nsalb/go-vsc-theme/fmtname"
)

// category tipo categria
type category string

const (
	// Themes Categoria para Temas
	Themes category = "Themes"
	// Languages Categoria para Linguagens
	Languages category = "Languages"
	// Snippets Categoria para Snippets
	Snippets category = "Snippets"
	// Linters Categoria para Linters
	Linters category = "Linters"
	// Debuggers Categoria para Debuggers
	Debuggers category = "Debuggers"
	// Other Putros tipos de categorias
	Other category = "Other"
)

type PackageOption struct {
	Publisher       string
	DescriptionTmpl string
}

// NewPackageOptionFromFile cria um PackageOption a partir de um arquivo JSON.
// Caso ocorra algum erro é retornado PackageOption default, no modelo a seguir:
// 	PackageOption {
// 		Publisher: "Undefined",
// 		DescriptionTmpl: "%s theme for Visual Studio Code",
// 	}
// E o erro ocorrido.
func NewPackageOptionFromFile(filename string) (PackageOption, error) {
	optDefault := PackageOption{"Undefined", "%s theme for Visual Studio Code"}

	file, err := os.Open(filename)
	if err != nil {
		return optDefault, err
	}
	err = json.NewDecoder(file).Decode(&optDefault)
	if err != nil {
		return optDefault, err
	}
	return optDefault, nil
}

// PackageManifest representa a entidade do package.json.
// São omitidos vários campos, para saber mais
// visite <https://code.visualstudio.com/docs/extensionAPI/extension-manifest>.
type PackageManifest struct {
	Name        string      `json:"name"`
	DisplayName string      `json:"displayName"`
	Description string      `json:"description"`
	Version     string      `json:"version"`
	Publisher   string      `json:"publisher"`
	Engines     Engine      `json:"engines"`
	Categories  []category  `json:"categories"`
	Contributes Contributes `json:"contributes"`
}

// NewPackageManifest retorna uma referência para um novo PackageManifest
func NewPackageManifest(displayName string, opt PackageOption) *PackageManifest {
	nameFolder := fmtname.ToLowerJoin(displayName)

	p := &PackageManifest{
		Name:        fmt.Sprintf("theme-%s", nameFolder),
		DisplayName: displayName,
		Description: fmt.Sprintf(opt.DescriptionTmpl, displayName),
		Version:     "0.0.1",
		Publisher:   opt.Publisher,
		Engines:     *NewEngine("^1.0.0"),
		Contributes: Contributes{},
	}
	p.AddCategory(Themes)

	return p
}

// AddCategory adiciona uma categoria a lista de categorias
func (p *PackageManifest) AddCategory(c category) {
	p.Categories = append(p.Categories, c)
}

// Engine contém a chave vscode que casa com a última versões compatível com a extensão
type Engine struct {
	Vscode string `json:"vscode"`
}

// NewEngine retorna uma referência para uma nova Engine
func NewEngine(version string) *Engine {
	return &Engine{Vscode: version}
}

// Contributes corresponde aos pontos de contribuição.
// Nessa implementação suporta somente Temas, para mais informações
// visite <https://code.visualstudio.com/docs/extensionAPI/extension-points>.
type Contributes struct {
	Themes []Theme `json:"themes"`
}

// AddTheme adiciona um Theme no objeto Contributes
func (c *Contributes) AddTheme(t Theme) {
	c.Themes = append(c.Themes, t)
}

// Theme representa um tema do TextMate para o VS Code.
// Para saber mais visite <https://code.visualstudio.com/docs/extensionAPI/extension-points#_contributesthemes>
type Theme struct {
	Label   string `json:"label"`
	UiTheme string `json:"uiTheme"`
	Path    string `json:"path"`
}

func newTheme(label, path, uiTheme string) *Theme {
	return &Theme{
		Label:   label,
		UiTheme: uiTheme,
		Path:    path,
	}
}

// NewThemeLight retorna uma referência para um novo Thema de base Light.
func NewThemeLight(label, path string) *Theme {
	return newTheme(label, path, "vs")
}

// NewThemeDark retorna uma referência para um novo Theme de base Dark.
func NewThemeDark(label, path string) *Theme {
	return newTheme(label, path, "vs-dark")
}
