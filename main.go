package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/saasan/go-s2dir"
	"github.com/saasan/go-s2file"
	"github.com/saasan/go-term"
)

func moveToParentDir(parent string, dirname string) error {
	// サブディレクトリとファイルを列挙
	dirs, files, err := s2dir.Read(dirname)
	if err != nil {
		return err
	}

	// サブディレクトリ内のファイルを移動
	for _, dir := range dirs {
		path := filepath.Join(dirname, dir.Name())
		if err := moveToParentDir(parent, path); err != nil {
			return err
		}
	}

	// ファイルを移動
	for _, file := range files {
		oldpath := filepath.Join(dirname, file.Name())
		newpath := filepath.Join(parent, file.Name())

		if oldpath == newpath {
			continue
		}

		s2file.Rename(oldpath, newpath)
	}

	isEmpty, err := s2dir.IsEmpty(dirname)
	if err != nil {
		return err
	}

	if !isEmpty {
		return nil
	}

	// ディレクトリが空であれば削除
	fmt.Printf("空フォルダを削除: %s\n", dirname)
	if err := os.Remove(dirname); err != nil {
		return fmt.Errorf("%s を削除できません。: %w", dirname, err)
	}

	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("処理対象: %s\n", arg)

		// サブディレクトリを列挙
		dirs, _, err := s2dir.Read(arg)
		if err != nil {
			continue
		}

		// サブディレクトリ内のファイルを移動
		for _, dir := range dirs {
			path := filepath.Join(arg, dir.Name())
			if err := moveToParentDir(arg, path); err != nil {
				fmt.Println(err)
			}
		}
	}

	if runtime.GOOS == "windows" {
		fmt.Println("Press any key to continue...")
		term.WaitAnyKey()
	}
}
