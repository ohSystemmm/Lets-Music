package Music

import (
	"fmt"
	"github.com/hajimehoshi/oto"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetFilesWithExtensions(dir string) (map[string]string, error) {
	filesMap := make(map[string]string)
	ext := []string{".wav", ".flac", ".aac", ".ogg", ".m4a", ".wma", ".mp3"}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			for _, ext := range ext {
				if filepath.Ext(file.Name()) == ext {
					fullPath := filepath.Join(dir, file.Name())
					filesMap[file.Name()] = fullPath
					break
				}
			}
		}
	}

	return filesMap, nil
}

func Run() {
	sourceDir := "./Backend/Music/source_songs"
	destDir := "./Backend/Music/mp3_songs"

	mp3Files, err := GetMusicFiles(sourceDir, destDir)
	if err != nil {
		log.Fatalf("Error getting music files: %v", err)
	}

	if len(mp3Files) == 0 {
		fmt.Println("No files were converted.")
		return
	}

	context, err := oto.NewContext(44100, 2, 2, 65536)
	if err != nil {
		log.Fatalf("Failed to create audio context: %v", err)
	}
	defer context.Close()

	musicPlayer := NewMusicPlayer(context)
	musicPlayer.PlaySongs(mp3Files)
}

func GetMusicFiles(sourceDir, destDir string) ([]string, error) {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory: %v", err)
		}
	}

	mp3Files, err := ConvertToMP3(sourceDir, destDir)
	if err != nil {
		return nil, fmt.Errorf("error converting files: %v", err)
	}

	return mp3Files, nil
}
