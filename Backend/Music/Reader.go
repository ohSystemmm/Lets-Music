package Music

import (
	"fmt"
	"io"
	"os"
)

func getDirectoryInfo(directory string) []string {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", directory)
		return nil
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	filenames := make([]string, 0, len(files))
	for _, file := range files {
		fmt.Println(file.Name())
		filenames = append(filenames, file.Name())
	}
	return filenames
}

func getAllSongs(directory string) [][]byte {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", directory)
		return nil
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	var allSongs [][]byte

	for _, file := range files {
		filePath := directory + "/" + file.Name()

		songFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", filePath, err)
			continue
		}
		defer songFile.Close()

		data, err := io.ReadAll(songFile)
		if err != nil {
			fmt.Println("Error reading file:", filePath, err)
			continue
		}

		allSongs = append(allSongs, data)
	}

	return allSongs
}
