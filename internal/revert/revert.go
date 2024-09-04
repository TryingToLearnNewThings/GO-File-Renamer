package revert

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func File2() {
	dir := "C:\\Users\\DrieMar\\Projekt_F\\GO\\cmd\\GoLearningPath\\samples"

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).Add(color.Underline).SprintFunc()
	blue := color.New(color.FgBlue).Add(color.Underline).SprintFunc()

	filenames := bufio.NewReader(os.Stdin)
	fmt.Print("Gib die Namen der Dateien an, die umbenannt werden sollen (z.B. birthday_001.txt): ")

	oldFilenames, _ := filenames.ReadString('\n')
	oldFilenames = strings.TrimSpace(oldFilenames)
	fmt.Printf("\nDer Ordner werden nach '%s' durchsucht\n\n", blue(oldFilenames))

	filenames2 := bufio.NewReader(os.Stdin)
	fmt.Print("Gib den neuen Namen an, zudem die Dateien umbenannt werden sollen (z.B. birthday-1.txt): ")

	newFilenames, _ := filenames2.ReadString('\n')
	newFilenames = strings.TrimSpace(newFilenames)
	fmt.Printf("\nDie ausgewählten Dateien werden von '%s' zu '%s' umbenannt\n\n", yellow(oldFilenames), green(newFilenames))

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Fehler beim Durchsuchen:", red(err))
			return err
		}
		if !info.IsDir() {
			log.Printf("Processing file: %s\n", info.Name())
		}
		if strings.Contains(info.Name(), oldFilenames) {
			err := os.Rename(path, filepath.Join(filepath.Dir(path), newFilenames))
			if err != nil {
				fmt.Println("Fehler beim umbenennen:", red(err))
				return err
			}
		}

		return nil
	})
	fmt.Println(blue("\n\nDie Datei wurde umbenannt"))
	if err != nil {
		log.Fatal(err)
	}
}
