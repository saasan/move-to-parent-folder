package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

// 正常系
func TestMoveToParentDirNormal(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestMoveToParentDir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1  = filepath.Join(tmpDir, "dir1")
		dir3  = filepath.Join(tmpDir, "dir1", "dir2", "dir3")
		file1 = filepath.Join(tmpDir, "dir1", "file1.txt")
		file2 = filepath.Join(tmpDir, "dir1", "dir2", "file2.txt")
		file3 = filepath.Join(tmpDir, "dir1", "dir2", "dir3", "file3.txt")
		move1 = filepath.Join(tmpDir, "file1.txt")
		move2 = filepath.Join(tmpDir, "file2.txt")
		move3 = filepath.Join(tmpDir, "file3.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir3, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1, []byte("file1"), os.ModePerm)
	os.WriteFile(file2, []byte("file2"), os.ModePerm)
	os.WriteFile(file3, []byte("file3"), os.ModePerm)

	moveToParentDir(tmpDir, dir1)

	if !fileExists(move1) {
		t.Error("move1失敗")
	}
	if !fileExists(move2) {
		t.Error("move2失敗")
	}
	if !fileExists(move3) {
		t.Error("move3失敗")
	}
	if fileExists(dir1) {
		t.Error("dir1失敗")
	}
}

// 親フォルダに同名のファイルが存在する場合
func TestMoveToParentDirExist(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestMoveToParentDir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1   = filepath.Join(tmpDir, "dir1")
		dir3   = filepath.Join(tmpDir, "dir1", "dir2", "dir3")
		file1  = filepath.Join(tmpDir, "dir1", "file1.txt")
		file2  = filepath.Join(tmpDir, "dir1", "dir2", "file2.txt")
		file3  = filepath.Join(tmpDir, "dir1", "dir2", "dir3", "file3.txt")
		exist1 = filepath.Join(tmpDir, "exist.txt")
		exist2 = filepath.Join(tmpDir, "dir1", "dir2", "exist.txt")
		move1  = filepath.Join(tmpDir, "file1.txt")
		move2  = filepath.Join(tmpDir, "file2.txt")
		move3  = filepath.Join(tmpDir, "file3.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir3, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1, []byte("file1"), os.ModePerm)
	os.WriteFile(file2, []byte("file2"), os.ModePerm)
	os.WriteFile(file3, []byte("file3"), os.ModePerm)
	os.WriteFile(exist1, []byte("exist1"), os.ModePerm)
	os.WriteFile(exist2, []byte("exist2"), os.ModePerm)

	moveToParentDir(tmpDir, dir1)

	if !fileExists(move1) {
		t.Error("move1失敗")
	}
	if !fileExists(move2) {
		t.Error("move2失敗")
	}
	if !fileExists(move3) {
		t.Error("move3失敗")
	}
	if !fileExists(exist1) {
		t.Error("exist1失敗")
	}
	if !fileExists(exist2) {
		t.Error("exist2失敗")
	}
	if !fileExists(dir1) {
		t.Error("dir1失敗")
	}
	if fileExists(dir3) {
		t.Error("dir3失敗")
	}
}

// 別ディレクトリに同名のファイルが存在する場合
func TestMoveToParentDirDuplicate(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestMoveToParentDir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1  = filepath.Join(tmpDir, "dir1")
		dir2  = filepath.Join(tmpDir, "dir1", "dir2")
		dir3  = filepath.Join(tmpDir, "dir1", "dir2", "dir3")
		file1 = filepath.Join(tmpDir, "dir1", "file1.txt")
		file2 = filepath.Join(tmpDir, "dir1", "dir2", "file2.txt")
		file3 = filepath.Join(tmpDir, "dir1", "dir2", "dir3", "file3.txt")
		dup1  = filepath.Join(tmpDir, "dir1", "dup.txt")
		dup2  = filepath.Join(tmpDir, "dir1", "dir2", "dup.txt")
		move1 = filepath.Join(tmpDir, "file1.txt")
		move2 = filepath.Join(tmpDir, "file2.txt")
		move3 = filepath.Join(tmpDir, "file3.txt")
		move4 = filepath.Join(tmpDir, "dup.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir3, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file1, []byte("file1"), os.ModePerm)
	os.WriteFile(file2, []byte("file2"), os.ModePerm)
	os.WriteFile(file3, []byte("file3"), os.ModePerm)
	os.WriteFile(dup1, []byte("dup1"), os.ModePerm)
	os.WriteFile(dup2, []byte("dup2"), os.ModePerm)

	moveToParentDir(tmpDir, dir1)

	if !fileExists(move1) {
		t.Error("move1失敗")
	}
	if !fileExists(move2) {
		t.Error("move2失敗")
	}
	if !fileExists(move3) {
		t.Error("move3失敗")
	}
	if !fileExists(move4) {
		t.Error("move4失敗")
	}
	if !fileExists(dup1) {
		t.Error("dup1失敗")
	}
	if fileExists(dup2) {
		t.Error("dup2失敗")
	}
	if !fileExists(dir1) {
		t.Error("dir1失敗")
	}
	if fileExists(dir2) {
		t.Error("dir2失敗")
	}
	if fileExists(dir3) {
		t.Error("dir3失敗")
	}
}
