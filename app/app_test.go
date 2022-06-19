package app

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp1(t *testing.T) {
	// Dry run in file mode
	work_dir, test_dirs, test_files := Setup()

	dir := work_dir
	mode := 0
	expr := `test (\d\d).mp4`
	name_template := `\${1}.mp4`
	dry_run := true

	args := NewAppArgs(dir,
		mode,
		expr,
		name_template,
		dry_run)
	Run(args)

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

	args := NewAppArgs(dir,
		mode,
		expr,
		name_template,
		dry_run)
	Run(args)

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

	args := NewAppArgs(dir,
		mode,
		expr,
		name_template,
		dry_run)
	Run(args)

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

	args := NewAppArgs(dir,
		mode,
		expr,
		name_template,
		dry_run)
	Run(args)

	assert.NoDirExists(t, test_dirs[0])
	assert.NoDirExists(t, test_dirs[1])
	assert.DirExists(t, path.Join(work_dir, "01"))
	assert.DirExists(t, path.Join(work_dir, "02"))
	assert.FileExists(t, test_files[0])
	assert.FileExists(t, test_files[1])

	TearDown(work_dir)
}
