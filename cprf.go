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

  // Simply copy the file
  if !stat.IsDir() {
    // cp -f
    os.Remove(filepath.Join(dst, stat.Name()))

    _, err = shutil.Copy(src, dst, false)

    return err
  }

  // If the source_file ends in a "/", the
  // contents of the directory are copied rather than the directory itself
  if !strings.HasSuffix(src, "/") {
    dst = filepath.Join(dst, stat.Name())
  }

  walk := func(path string, info os.FileInfo, err error) error {
    stat, err := os.Lstat(path)
    if err != nil {
      return err
    }

    dstTemp := filepath.Join(dst, strings.Replace(path, src, "", 1))

    // "Copy" directory
    if stat.IsDir() {
      return mkDir(dstTemp, stat.Mode())
    }

    // cp -f
    os.Remove(dstTemp)

    _, err = shutil.Copy(path, dstTemp, false)

    return err
  }

  return filepath.Walk(src, walk)
}

func mkDir(path string, mode os.FileMode) error {
  if _, err := os.Lstat(path); err == nil {
    return nil
  }

  return os.Mkdir(path, mode)
}
