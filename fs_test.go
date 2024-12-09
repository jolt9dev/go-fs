package fs_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/jolt9dev/go-fs"
	"github.com/stretchr/testify/assert"
)

func init() {
	fs.MkdirAllDefault("testdir")
	fs.WriteFile("testfile", []byte("test data"), 0644)
	fs.Symlink("testfile", "testsymlink")
}

func TestChown(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Chown is not supported on Windows")
	}

	err := fs.Chown("testfile", 1000, 1000)
	assert.NoError(t, err)
}

func TestChmod(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Chmod is not supported on Windows")
	}

	err := fs.Chmod("testfile", 0644)
	assert.NoError(t, err)
}

func TestCopy(t *testing.T) {
	defer fs.Remove("testfile_copy")

	err := fs.Copy("testfile", "testfile_copy", true)
	assert.NoError(t, err)
}

func TestCopyDir(t *testing.T) {
	defer fs.RemoveAll("testdir_copy")

	err := fs.CopyDir("testdir", "testdir_copy", true)
	assert.NoError(t, err)
}

func TestCopyFile(t *testing.T) {
	defer fs.Remove("testfile_copy")

	err := fs.CopyFile("testfile", "testfile_copy", true)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	defer fs.Remove("testfile2")
	file, err := fs.Create("testfile2")
	assert.NoError(t, err)
	assert.NotNil(t, file)
	file.Close()
}

func TestCreateTemp(t *testing.T) {
	file, err := fs.CreateTemp("", "testfile")
	if err == nil {
		defer fs.Remove(file.Name())
	}
	assert.NoError(t, err)
	assert.NotNil(t, file)
	file.Close()
}

func TestCwd(t *testing.T) {
	dir, err := fs.Cwd()
	assert.NoError(t, err)
	assert.NotEmpty(t, dir)
}

func TestExists(t *testing.T) {
	exists := fs.Exists("testfile")
	assert.True(t, exists)

	exists = fs.Exists("testdir")
	assert.True(t, exists)

	exists = fs.Exists("testfile999")
	assert.False(t, exists)
}

func TestEnsureDir(t *testing.T) {
	err := fs.EnsureDir("testdir", 0755)
	assert.NoError(t, err)

	defer fs.RemoveAll("testdir10")
	err = fs.EnsureDir("testdir10", 0755)
	assert.NoError(t, err)
}

func TestEnsureDirDefault(t *testing.T) {
	err := fs.EnsureDirDefault("testdir")
	assert.NoError(t, err)
}

func TestEnsureFile(t *testing.T) {
	err := fs.EnsureFile("testfile", 0644)
	assert.NoError(t, err)

	defer fs.Remove("testfile10")
	err = fs.EnsureFile("testfile10", 0644)
	assert.NoError(t, err)
}

func TestEnsureFileDefault(t *testing.T) {
	err := fs.EnsureFileDefault("testfile")
	assert.NoError(t, err)
}

func TestIsFile(t *testing.T) {
	isFile := fs.IsFile("testfile")
	assert.True(t, isFile)
}

func TestIsDir(t *testing.T) {
	isDir := fs.IsDir("testdir")
	assert.True(t, isDir)
}

func TestIsSymlink(t *testing.T) {
	defer fs.Remove("testsymlink")
	isSymlink := fs.IsSymlink("testsymlink")
	assert.True(t, isSymlink)
}

func TestLink(t *testing.T) {
	defer fs.Remove("testfile_link")
	err := fs.Link("testfile", "testfile_link")
	assert.NoError(t, err)
}

func TestLstat(t *testing.T) {
	info, err := fs.Lstat("testfile")
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestMkdir(t *testing.T) {
	defer fs.Remove("testdir700")
	err := fs.Mkdir("testdir700", 0755)
	assert.NoError(t, err)
}

func TestMkdirDefault(t *testing.T) {
	defer fs.Remove("testdir800")
	err := fs.MkdirDefault("testdir800")
	assert.NoError(t, err)
}

func TestMkdirAll(t *testing.T) {
	defer fs.RemoveAll("testdir900")
	err := fs.MkdirAll("testdir900/subdir", 0755)
	assert.NoError(t, err)
}

func TestMkdirAllDefault(t *testing.T) {
	defer fs.RemoveAll("testdir2000")
	err := fs.MkdirAllDefault("testdir2000/subdir")
	assert.NoError(t, err)
}

func TestOpen(t *testing.T) {
	file, err := fs.Open("testfile")
	assert.NoError(t, err)
	assert.NotNil(t, file)
	file.Close()
}

func TestOpenFile(t *testing.T) {
	file, err := fs.OpenFile("testfile", os.O_RDWR|os.O_CREATE, 0644)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	file.Close()
}

func TestResolve(t *testing.T) {
	path, err := fs.Resolve("testfile", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, path)
}

func TestRemove(t *testing.T) {
	fs.EnsureFile("testfile88", 0644)
	err := fs.Remove("testfile88")
	assert.NoError(t, err)
}

func TestReadFile(t *testing.T) {
	data, err := fs.ReadFile("testfile")
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, "test data", string(data))
}

func TestReadTextFile(t *testing.T) {
	data, err := fs.ReadTextFile("testfile")
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, "test data", data)
}

func TestReadFileLines(t *testing.T) {
	lines, err := fs.ReadFileLines("testfile")
	assert.NoError(t, err)
	assert.NotEmpty(t, lines)
	assert.Equal(t, []string{"test data"}, lines)
}

func TestRemoveAll(t *testing.T) {
	fs.EnsureDir("testdir9999/text", 0755)

	err := fs.RemoveAll("testdir9999")
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	fs.EnsureFile("testfilex", 0644)
	defer fs.Remove("testfile_renamed")
	err := fs.Rename("testfilex", "testfile_renamed")
	assert.NoError(t, err)
	assert.True(t, fs.Exists("testfile_renamed"))
	assert.False(t, fs.Exists("testfilex"))
}

func TestStat(t *testing.T) {
	info, err := fs.Stat("testfile")
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestSymlink(t *testing.T) {
	defer fs.Remove("testfile_symlink")
	err := fs.Symlink("testfile", "testfile_symlink")
	assert.NoError(t, err)
}

func TestWalkDir(t *testing.T) {
	err := fs.WalkDir("testdir", func(path string, d fs.DirEntry, err error) error {
		return nil
	})
	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	defer fs.Remove("testfile69")
	err := fs.WriteFile("testfile69", []byte("test data2"), 0644)
	assert.NoError(t, err)
	data, err := fs.ReadTextFile("testfile69")
	assert.NoError(t, err)
	assert.Equal(t, "test data2", data)
}

func TestWriteFileLines(t *testing.T) {
	defer fs.Remove("testfile79")
	err := fs.WriteFileLines("testfile79", []string{"line1", "line2"}, 0644)
	assert.NoError(t, err)
	data, err := fs.ReadFileLines("testfile79")
	assert.NoError(t, err)
	assert.Equal(t, []string{"line1", "line2"}, data)
}

func TestWriteFileLinesSep(t *testing.T) {
	defer fs.Remove("testfile89")
	err := fs.WriteFileLinesSep("testfile89", []string{"line1", "line2"}, "\n", 0644)
	assert.NoError(t, err)

	data, err := fs.ReadFileLines("testfile89")
	assert.NoError(t, err)
	assert.Equal(t, []string{"line1", "line2"}, data)
}

func TestWriteTextFile(t *testing.T) {
	defer fs.Remove("testfile10")
	err := fs.WriteTextFile("testfile10", "test data10", 0644)
	assert.NoError(t, err)

	data, err := fs.ReadTextFile("testfile10")
	assert.NoError(t, err)
	assert.Equal(t, "test data10", data)
}
