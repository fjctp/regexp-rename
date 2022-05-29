package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Setup() (string, []string, []string) {
	tmp_dir, err1 := ioutil.TempDir("", "regexp-rename-test-*")
	Check(err1)
	log.Println("Created temp directory: ", tmp_dir)

	test_dirs := []string{
		path.Join(tmp_dir, "test 01"),
		path.Join(tmp_dir, "test 02"),
	}

	test_files := []string{
		path.Join(tmp_dir, "test 01.mp4"),
		path.Join(tmp_dir, "test 02.mp4"),
	}

	for _, test_dir := range test_dirs {
		err2 := os.Mkdir(test_dir, 0755)
		Check(err2)
	}

	for _, test_file := range test_files {
		_, err3 := os.Create(test_file)
		Check(err3)
	}

	return tmp_dir, test_dirs, test_files
}

func TearDown(tmp_dir string) {
	//os.RemoveAll(tmp_dir)
	log.Println("Removed temp directory: ", tmp_dir)
}

func TestApp1(t *testing.T) {
	// Dry run in file mode
	work_dir, test_dirs, test_files := Setup()

	dir := work_dir
	mode := 0
	expr := `test (\d\d).mp4`
	name_template := `\${1}.mp4`
	dry_run := true
	app(dir,
		mode,
		expr,
		name_template,
		dry_run)

	assert.DirExists(t, test_dirs[0])
	assert.DirExists(t, test_dirs[1])
	assert.FileExists(t, test_files[0])
	assert.FileExists(t, test_files[1])

	TearDown(work_dir)
}

func TestApp2(t *testing.T) {
	// Dry run in directory mode
	work_dir, test_dirs, test_files := Setup()

	dir := work_dir
	mode := 1
	expr := `test (\d\d)`
	name_template := `${1}`
	dry_run := true
	app(dir,
		mode,
		expr,
		name_template,
		dry_run)

	assert.DirExists(t, test_dirs[0])
	assert.DirExists(t, test_dirs[1])
	assert.FileExists(t, test_files[0])
	assert.FileExists(t, test_files[1])

	TearDown(work_dir)
}

func TestApp3(t *testing.T) {
	// File mode
	work_dir, test_dirs, test_files := Setup()

	dir := work_dir
	mode := 0
	expr := `(?m)test (\d\d).mp4`
	name_template := `${1}.mp4`
	dry_run := false
	app(dir,
		mode,
		expr,
		name_template,
		dry_run)

	assert.DirExists(t, test_dirs[0])
	assert.DirExists(t, test_dirs[1])
	assert.NoFileExists(t, test_files[0])
	assert.NoFileExists(t, test_files[1])
	assert.FileExists(t, path.Join(work_dir, "01.mp4"))
	assert.FileExists(t, path.Join(work_dir, "02.mp4"))

	TearDown(work_dir)
}

func TestApp4(t *testing.T) {
	// Directory mode
	work_dir, test_dirs, test_files := Setup()

	dir := work_dir
	mode := 1
	expr := `(?m)test (\d\d)`
	name_template := `${1}`
	dry_run := false
	app(dir,
		mode,
		expr,
		name_template,
		dry_run)

	assert.NoDirExists(t, test_dirs[0])
	assert.NoDirExists(t, test_dirs[1])
	assert.DirExists(t, path.Join(work_dir, "01"))
	assert.DirExists(t, path.Join(work_dir, "02"))
	assert.FileExists(t, test_files[0])
	assert.FileExists(t, test_files[1])

	TearDown(work_dir)
}
