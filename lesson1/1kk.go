package lesson1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Для закрепления практических навыков программирования, напишите программу, которая создаёт один миллион пустых
// файлов в известной, пустой директории файловой системы используя вызов os.Create. Ввиду наличия определенных
// ограничений операционной системы на число открытых файлов, такая программа должна выполнять аварийную остановку.
// Запустите программу и дождитесь полученной ошибки. Используя отложенный вызов функции закрытия файла,
// стабилизируйте работу приложения. Критерием успешного выполнения программы является успешное создание миллиона
// пустых файлов в директории

// 1kk iterations. Into each iterate creating a file and open it.
// After all iterations complete, it removes the directory with 1kk files.
func IterateWithOneMillionFiles() error {
	defer func() {
		storeDir := getStoreDir()
		_ = os.RemoveAll(storeDir)
	}()

	storeDir := getStoreDir()

	if err := os.MkdirAll(storeDir, 0777); err != nil {
		return err
	}

	log.Println("Creating one million files and open each one")

	for i := 0; i < 1_000_000; i++ {
		fileName := "file" + strconv.Itoa(i) + ".txt"
		filePath := storeDir + string(os.PathSeparator) + fileName

		if err := someFileOperation(filePath); err != nil {
			return err
		}
	}

	return nil
}

// Return the store path of 1kk files
func getStoreDir() string {
	tempDir := os.TempDir()

	storeDir := strings.Join([]string{tempDir, "1kk"}, string(os.PathSeparator))

	return storeDir
}

// Some fake operation with the file
func someFileOperation(filePath string) error {
	f, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
