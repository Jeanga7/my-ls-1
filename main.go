package main

import (
	"flag"
	"fmt"
	"my_ls/internal/display"
	"my_ls/internal/fileinfo"
	"my_ls/internal/flags"
	"my_ls/internal/sort"
	"os"
)

func main() {
	// Définir les drapeaux
	l := flag.Bool("l", false, "Utiliser un format de liste longue")
	R := flag.Bool("R", false, "Lister les sous-répertoires récursivement")
	a := flag.Bool("a", false, "Ne pas ignorer les entrées commençant par .")
	r := flag.Bool("r", false, "Inverser l'ordre du tri")
	t := flag.Bool("t", false, "Trier par date de modification")

	flag.Parse()

	// Obtenir les arguments restants (chemins de fichiers/répertoires)
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	// Créer l'objet options
	options := flags.Options{
		Long:      *l,
		Recursive: *R,
		All:       *a,
		Reverse:   *r,
		Time:      *t,
	}

	// Traiter chaque argument
	for _, path := range args {
		err := processPath(path, options)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}

func processPath(path string, options flags.Options) error {
	// Obtenir les informations sur les fichiers
	files, err := fileinfo.GetFileInfo(path, options)
	if err != nil {
		return err
	}

	// Trier les fichiers
	sort.SortFiles(files, options)

	// Afficher les fichiers
	display.DisplayFiles(files, options)

	return nil
}
