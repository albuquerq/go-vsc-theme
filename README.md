# go-vsc-theme


É uma ferramenta de linha de comando escrita em Go para criação em lote de pacotes de temas para o Visual Studio Code.
Serve como uma alternativa para o *yo code extension* para geração de pacotes de temas.
Seu uso é aconselhado quando se tem muitos arquivos de temas do TextMate (.tmTheme) e pretende-se gerar os seus
respectivos pacotes de forma automatizada.

## Getting & Install


Clone o repositório e compile. Para compilar é necessário instalar a linguagem de programação Go, [download Go](https://golang.org/dl/).

```bash
git clone https://github.com/albuquerq/go-vsc-theme.git
cd go-vsc-theme
go build .
# Depende de templates/README.tmpl
```

## Uso

Para entender o formato de pacotes de temas do Visual Studio Code veja a [documentação](https://code.visualstudio.com/docs/customization/themes).

Ajuda: go-vsc-theme --help

```
Usage of go-vsc-theme:
  -auto
    	Set to false if you want to enter the DisplayName manually. (default true)
  -dark
    	Set true for dark theme files.
  -in string
    	Directory containing files ".tmTheme". (default ".")
  -opt string
    	Path to JSON options file. (default "./opt.json")
  -out string
    	Destination directory for packages. (default ".")
``` 

Imagine a situação:
Você tem uma pasta nomeada "temas_claros" contendo vários arquivos de tema do TextMake (.tmTheme). E uma pasta que conterá os pacotes gerados, nomeda de "pacotes".
Conforme a seguir:
```
temas_claros/
    tema1.tmTheme
    tema2.tmTheme
    temaClaro3.tmTheme
pacotes/
    vazio         
```

Configure o arquivo opt.json com o seu nome de publicador.

opt.json
---------------
```JavaScript
{
    "Publisher": "SeuNomeDePublicador",
    "DescriptionTmpl": "A mensagem de descrição do pacote com o nome %s"
}
```

Se você predente rodar o gerador no modo automático, nesse modo os pacotes recebem nomes baseados nos dos arquivos .tmTheme.
Faça:
```bash
$./go-vsc-theme -in=/caminho/para/temas_claros -out=caminho/para/pacotes -dark=false -opt=opt.json
```

Se você pretende rodar o gerador no moto interativo, onde é possível atribuir os nomes dos temas dos pacotes, então set a flag -auto=false:
```bash
$./go-vsc-theme -in=/caminho/para/temas_claros -auto=false -out=caminho/para/pacotes -dark=false -opt=opt.json
```

> Observação1: go-vsc-theme roda recursivamente no caminho passado no parâmetro -in, procurando nos subdiretórios por arquivos de tema.
> Por esse motivo **não** passe como -out um subdiretório de -in!

> Observação2: Quando for escrever entra manualmente como uma nome para o tema, escreva-o entre aspas, ex.: "Meu Tema Claro"

## Personalizar o README.md do Pacote de Tema

Você pode personalizar o documento de README.md do pacote através do arquivo de template em "templates/README.tmpl".
Para mais informações sobre a sintaxe da engine de templates visite a [documentação do pacote text/templates da linguagem Go.](https://golang.org/pkg/text/template/)
