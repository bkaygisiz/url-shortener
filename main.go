package main

import (
	"github.com/bkaygisiz/url-shortener/cmd"
	_ "github.com/bkaygisiz/url-shortener/cmd/cli"    // Importe le package 'cli' pour que ses init() soient exécutés
	_ "github.com/bkaygisiz/url-shortener/cmd/server" // Importe le package 'server' pour que ses init() soient exécutés
)

func main() {
	cmd.Execute()
}
