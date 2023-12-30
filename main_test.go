package main

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/longbridgeapp/opencc"
)

func TestConvertPath(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create some test files and directories
	dir1 := filepath.Join(tempDir, "微服务框架之网络编程与最简 RPC")
	dir2 := filepath.Join(tempDir, "学习手册")
	file1 := filepath.Join(dir2, "学习手册_1.txt")
	file2 := filepath.Join(dir1, "连接池：基本原理、开源实例 silenceper&pool.mp4")
	err = os.Mkdir(dir1, 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Mkdir(dir2, 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(file1, []byte("file1"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(file2, []byte("file2"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create an OpenCC converter
	converter, err := opencc.New("s2t")
	if err != nil {
		t.Fatal(err)
	}

	// Create a wait group to wait for goroutines to finish
	var wg sync.WaitGroup

	// Call the convertPath function
	convertPath(converter, tempDir, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if the files and directories have been renamed correctly
	renamedDir1 := filepath.Join(tempDir, "微服務框架之網絡編程與最簡 RPC")
	renamedDir2 := filepath.Join(tempDir, "學習手冊")
	renamedFile1 := filepath.Join(renamedDir2, "學習手冊_1.txt")
	renamedFile2 := filepath.Join(renamedDir1, "連接池：基本原理、開源實例 silenceper&pool.mp4")
	if _, err := os.Stat(renamedFile1); os.IsNotExist(err) {
		t.Errorf("File %s does not exist", renamedFile1)
	}
	if _, err := os.Stat(renamedFile2); os.IsNotExist(err) {
		t.Errorf("File %s does not exist", renamedFile2)
	}
	if _, err := os.Stat(renamedDir1); os.IsNotExist(err) {
		t.Errorf("Directory %s does not exist", renamedDir1)
	}
	if _, err := os.Stat(renamedDir2); os.IsNotExist(err) {
		t.Errorf("Directory %s does not exist", renamedDir2)
	}
}