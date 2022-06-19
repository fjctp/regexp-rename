package app

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Run(args AppArgs) {
	log.Println("Work dir:", args.absDir)

	// loop through items in a directory
	files, err := ioutil.ReadDir(args.absDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		doRun := (file.IsDir() && (args.mode == DirMode)) || (!file.IsDir() && (args.mode == FileMode))
		if doRun {
			// If a dir and in dir mode or a file and in file mode, then
			old_name := file.Name()
			new_name := args.expr.ReplaceAllString(old_name, args.name_template)

			is_skip := old_name == new_name
			if is_skip {
				log.Print(old_name, " -> ", "Skip")
			} else {
				log.Print(old_name, " -> ", new_name)
			}

			if !(args.dry_run) && !is_skip {
				err := os.Rename(
					filepath.Join(args.absDir, old_name),
					filepath.Join(args.absDir, new_name))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
