// Package archive which represents `cp -rf` logic
package cprf

import (
  "os"
  "path/filepath"
  "strings"

  "github.com/termie/go-shutil"
)

func removeIfExist(path string) {
  if _, err := os.Stat(path); err == nil {
    os.Remove(path)
  }
}

func copyDir(path string, mode os.FileMode) error {
  if _, err := os.Stat(path); err == nil {
    return nil
  }

  return os.Mkdir(path, mode)
}

func Copy(src, dst string) error {
  stat, err := os.Stat(src)

  if err != nil {
    return err
  }

  // Simply copy the file
  if !stat.IsDir() {
    _, err = shutil.Copy(src, dst, false)
    return err
  }

  // If the source_file ends in a "/", the
  // contents of the directory are copied rather than the directory itself
  if !strings.HasSuffix(src, "/") {
    dst = filepath.Join(dst, stat.Name())
  }

  walk := func(path string, info os.FileInfo, err error) error {
    stat, err := os.Stat(path)
    if err != nil {
      return err
    }

    dstTemp := filepath.Join(dst, strings.Replace(path, src, "", 1))

    // "Copy" directory
    if stat.IsDir() {
      return copyDir(dstTemp, stat.Mode())
    }

    // cp -f
    removeIfExist(dstTemp)

    // File copy
    _, err = shutil.Copy(path, dstTemp, false)
    return err
  }

  return filepath.Walk(src, walk)
}
