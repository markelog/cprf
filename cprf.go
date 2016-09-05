// Package cprf which represents `cp -Rf` logic
package cprf

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/termie/go-shutil"
)

// Copy copies yo!
func Copy(src, dst string) error {
	stat, err := os.Lstat(src)

	if err != nil {
		return err
	}

	// Simply copy the file if it is not a dir
	if stat.IsDir() == false {
		dst = filepath.Join(dst, stat.Name())

		return copy(src, dst)
	}

	// If the source_file ends in a "/", the
	// contents of the directory are copied rather than the directory itself
	if !strings.HasSuffix(src, "/") {
		dst = filepath.Join(dst, stat.Name())
	}

	walk := func(path string, info os.FileInfo, err error) error {
		dstTemp := filepath.Join(dst, strings.Replace(path, src, "", 1))

		// "Copy" directory
		if info.IsDir() {
			return mkDir(dstTemp, info.Mode())
		}

		return copy(path, dstTemp)
	}

	return filepath.Walk(src, walk)
}

func copy(src, dst string) error {
	stat, err := os.Lstat(src)

	if err != nil {
		return err
	}

	// cp -f
	os.Remove(dst)

	// "Copy" symlinks
	if shutil.IsSymlink(stat) {
		return copySymLink(src, dst)
	}

	_, err = shutil.Copy(src, dst, false)

	return err
}

// Can't use "shutil" since it absolutize path
func copySymLink(src, dst string) (err error) {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}

	return os.Symlink(link, dst)
}

func mkDir(path string, mode os.FileMode) error {
	if _, err := os.Lstat(path); err == nil {
		return nil
	}

	return os.Mkdir(path, mode)
}
