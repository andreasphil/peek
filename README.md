<h1 align="center">
  Peek 🌚
</h1>

<p align="center">
  <strong>Markdown preview</strong>
</p>

> [!NOTE]
>
> Work in progress. Things are most certainly incomplete and/or broken, and will
> definitely change.

- 🌎 Instantly serve an HTML preview of local markdown files
- 📤 Easily copy or share the rendered HTML
- 🚧 Live updating on changes
- 🚧 Display local images and navigate between linked Markdown files
- 🚧 Export as a static page

## Installation

```sh
go install github.com/andreasphil/peek@latest
```

## Usage

Launch a local fileserver rendering a preview of `file` at <http://localhost:8080>:

```sh
# Options:
#
# -allow-unsafe
#   	render inline HTML (default true)
# -port string
#   	the port for serving the application (default "8080")

peek <file>
```

## Development

Peek is written in Go, with some HTML, CSS and JavaScript for showing and enhancing the preview. The following commands are available:

```sh
air           # build + run in watch mode (requires @air-verse/air)
go run .      # run
go build      # build
```

## Credits

This app uses a number of open source packages listed in [go.mod](./go.mod).

Thanks 🙏
