package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TryingToLearnNewThings/GoLearningPath/internal/revert"
	"github.com/fatih/color"
)

func main() {

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	folder := bufio.NewReader(os.Stdin)
	fmt.Println("File Renamer")
	fmt.Println("---------------------")

	fmt.Println("Gib den Pfad (z.b. c:\\Users\\Name\\samples) von den Dateien an die umbenannt werden sollen")
	fmt.Print(" -> ")
	folderName, _ := folder.ReadString('\n')
	folderName = strings.Replace(folderName, "\n", "", -1)
	fmt.Printf("Der Ordner wird nach den Dateien mit dem Namen %s durchsucht \n", folderName)

	dir := "C:\\Users\\DrieMar\\Projekt_F\\GO\\cmd\\GoLearningPath\\samples"

	// read the Name of the file that needs to be changed from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Gib den Namen der Datei an die geändert werden sollen (zahlen bitte mit * angeben)")
	fmt.Print(" -> ")
	oldName, _ := reader.ReadString('\n')
	oldName = strings.Replace(oldName, "\n", "", -1)
	fmt.Printf("Der Ordner wird nach den Dateien mit dem Namen %s durchsucht \n", yellow(oldName))

	// Collect the Name that the file should be renambed to from user
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Gib den Namen an zu dem die Datei geändert werden soll (zahlen bitte mit * angeben (anzahl der * = anzahl der Nullen))")
	fmt.Print(" -> ")
	newName, _ := reader2.ReadString('\n')
	newName = strings.Replace(newName, "\n", "", -1)
	fmt.Printf("Die gefundenen Dateien werden mit den angegebenen Namen %s geändert \n", green(newName))

	// Read all the entries in the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	// create an empty slice to store the files that need to be renamed
	var renames []string
	for _, entry := range entries {
		// Skip directories
		if entry.IsDir() {
			continue
		}
		filename := entry.Name()

		// Check if the file needs to be renamed
		if _, err := match(filename); err == nil {
			// Add the filename to the list of files to rename
			renames = append(renames, filename)
		}
	}
	for _, filename := range renames {
		// Construct the original file path
		origPath := filepath.Join(dir, filename)

		// Get the new filename by matching the original filename
		newFilename, err := match(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}

		// Construct the new file path
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("Renaming %s to %s\n", yellow(origPath), green(newPath))
		// Rename the file
		err = os.Rename(origPath, newPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	}

	fmt.Println("The files have been successfully renamed.")

	// Wait for 8 seconds and then rename the files back
	time.Sleep(8 * time.Second)
	revert.File2()
}

func match(filename string) (string, error) {
	// Split the filename into base and extension
	parts := strings.SplitN(filename, ".", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("%s: invalid filename", filename)
	}

	base := parts[0]
	ext := parts[1]

	// Split the base into name and number
	pieces := strings.Split(base, "_")
	if len(pieces) < 2 {
		return "", fmt.Errorf("%s: invalid filename", filename)
	}

	name := strings.Join(pieces[:len(pieces)-1], "_")
	numStr := pieces[len(pieces)-1]

	// Convert the number to an integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", fmt.Errorf("%s: invalid filename", filename)
	}

	// Increment the number and return the new filename
	return fmt.Sprintf("%s-%d.%s", name, num, ext), nil
}
