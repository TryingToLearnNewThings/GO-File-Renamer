package revert

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func File2() {
	dir := "./samples"

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

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
		if _, err := matching(filename); err == nil {
			// Add the filename to the list of files to rename
			renames = append(renames, filename)
		}
	}
	for _, filename := range renames {
		// Construct the original file path
		origPath := filepath.Join(dir, filename)

		// Get the new filename by matching the original filename
		newFilename, err := matching(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}

		// Construct the new file path
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("Renaming %s to %s\n", origPath, newPath)
		fmt.Printf("Renaming %s to %s\n", yellow(origPath), green(newPath))
		// Rename the file
		err = os.Rename(origPath, newPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	}

	fmt.Println("The files have been successfully renamed back.")
}

func matching(filename string) (string, error) {
	// Split the filename into base and extension
	parts := strings.SplitN(filename, ".", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("%s: invalid filename", filename)
	}

	base := parts[0]
	ext := parts[1]

	// Split the base into name and number
	pieces := strings.Split(base, "-")
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
	return fmt.Sprintf("%s.%d.%s", name, num+1, ext), nil
}
