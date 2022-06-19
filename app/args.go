package app

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type ModeEnum int

const (
	FileMode ModeEnum = 0
	DirMode  ModeEnum = 1
)

type AppArgs struct {
	absDir        string
	mode          ModeEnum
	expr          *regexp.Regexp
	name_template string
	dry_run       bool
}

func NewAppArgs(dir string, mode int, expr string, name_template string, dry_run bool) AppArgs {
	var args AppArgs
	args.validateInputs(dir, mode, expr, name_template, dry_run)

	return args
}

func (args *AppArgs) validateInputs(dir string, mode int, expr string, name_template string, dry_run bool) {
	args.absDir = validateDir(dir)
	args.mode = validateMode(mode)
	args.expr = validateExpr(expr)
	args.name_template = validateNameTemplate(name_template)
	args.dry_run = dry_run
}

func validateDir(dir string) string {
	dir_abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(dir_abs); os.IsNotExist(err) {
		log.Fatal("Path does not exist")
	}

	return dir_abs
}

func validateMode(mode int) ModeEnum {
	if mode == 0 {
		return FileMode
	} else if mode == 1 {
		return DirMode
	} else {
		log.Fatal("Invalid mode")
		return -1
	}
}

func validateExpr(expr string) *regexp.Regexp {
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}

	return re
}

func validateNameTemplate(name_template string) string {
	return name_template
}
