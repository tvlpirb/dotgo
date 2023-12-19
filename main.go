package main

import (
	"flag"
  "bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var themeDir string = "/home/talhah/Projects/desktop-env/themes"
var theme string = ""
var finalDir string

func main(){
  flag.StringVar(&themeDir,"d",themeDir,"Path to directory containing themes") 
  flag.StringVar(&themeDir,"dir",themeDir,"Path to directory containing themes")
  flag.StringVar(&theme,"t",theme,"Name of theme")
  flag.StringVar(&theme,"theme",theme,"Name of theme")
  
  flag.Parse()

  finalDir = filepath.Join(themeDir,theme) 
  if _, err := os.Stat(finalDir); os.IsNotExist(err) {
    log.Fatal("Given theme does not exist, please try again")
  } else {
    if theme != "" && themeDir != "" {
      switchTheme(finalDir)
      return
    }
  }

  if err := os.Chdir(finalDir); err != nil {
    log.Fatal("Could not enter theme directory")
  }

  themes, err := listThemes(finalDir)
  if err != nil {
    log.Fatal("Error listing themes")
  }

  for i, theme := range themes {
    fmt.Printf("%d) %s\n",i+1,theme)
  }
  
  var input string
  var index int
  var valid bool
  for {
    fmt.Printf("Pick a theme: ")
    fmt.Scanln(&input)
    index, valid = validateChoice(input, len(themes))
    
    if valid {
      break
    }
    fmt.Println("Incorrect value passed")
  }
  theme = themes[index]
  finalDir = filepath.Join(themeDir,theme)

  switchTheme(finalDir)
}

func switchTheme(finalDir string) {
  os.Chdir(finalDir)
  userPath, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting user configuration directory:", err)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	cwd += "/"
	skipAll := false

	entries, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory entries:", err)
		return
	}

	for _, entry := range entries {
		fromPath := cwd + entry.Name()
		toPath := filepath.Join(userPath, entry.Name())
		userInput := ""

		err := os.Symlink(fromPath, toPath)
		if err != nil {
			if _, ok := err.(*os.LinkError); ok {
				if !skipAll {
					fmt.Printf("There is already a configuration for %s\n", entry.Name())
					for userInput != "y" && userInput != "n" && userInput != "all" {
						fmt.Print("Do you want to delete the current config? (y/n/all): ")
						userInput = readInput()
						if userInput == "all" {
							skipAll = true
						}
					}
				}

				if skipAll || userInput == "y" {
					deleteFile(toPath)
          os.Symlink(fromPath, toPath)
					fmt.Printf("Created %s\n", toPath)
				} else {
					fmt.Printf("Skipping config for %s\n", entry.Name())
				}
			}
		}
	}
} 

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Printf("Error deleting file %s: %v\n", path, err)
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return input[:len(input)-1] // Remove newline character
}

// Given the choice it validates it and returns the index 
func validateChoice(input string, length int) (int, bool){
  num,err := strconv.Atoi(input)
  if err != nil{
    return 0,false
  } 

  if num <= length {
   return num-1,true
  }

  return 0, false
}

// Returns a slice containing the themes 
func listThemes(dir string) ([]string, error){
  entries, err := os.ReadDir(dir)
  if err != nil {
    log.Fatal("Could not list themes")
  }
  
  var themes []string
  for _, entry := range entries {
    if entry.IsDir(){
      themes = append(themes,entry.Name())
    }
  }
 
  return themes, nil
}
