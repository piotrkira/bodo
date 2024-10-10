package main

import (
	"errors"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Entry struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Section struct {
	Name    string  `yaml:"name"`
	Entries []Entry `yaml:"entries"`
}

type Theme struct {
	TextColor       string `yaml:"text_color"`
	BackgroundColor string `yaml:"background_color"`
	PrimaryColor    string `yaml:"primary_color"`
}

type ThemeFile map[string]Theme

type Config struct {
	Version  int       `yaml:"version"`
	Title    string    `yaml:"title"`
	Sections []Section `yaml:"sections"`
	Columns  int       `yaml:"columns"`
	Font     string    `yaml:"font"`
}

type IndexData struct {
	Title    string
	Sections []Section
	Columns  int
	Themes   ThemeFile
	Font     string
}

var (
	configFilePath string
	themesFilePath string
	indexFilePath  string
)

func getFirstPath(paths []string) string {
	for i := 0; i < len(paths) - 1; i++ {
		if _, err := os.Stat(paths[i]); !errors.Is(err, fs.ErrNotExist) {
			return paths[i]
		}
	}
	return paths[len(paths)-1]
}

func loadThemesFile(path string) *ThemeFile {
	themesFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Theme file is not available, using default theme\n")
	}
	var themes ThemeFile
	err = yaml.Unmarshal(themesFile, &themes)
	if err != nil {
		log.Fatalf("Theme file is invalid: %v\n", err)
	}
	return &themes
}

func loadConfigFile() *Config {
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Config file is not availabe %v\n", err)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Config file is invalid: %v\n", err)
	}
	return &config
}


func getIndexHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		config := loadConfigFile()
		var themes *ThemeFile
		if themesFilePath == "" {
			themes = loadThemesFile(getFirstPath([]string{"/etc/bodo/themes.yaml", "/usr/local/share/bodo/themes.yaml", "themes.yaml"}))
		} else {
			themes = loadThemesFile(themesFilePath)
		}
		indexData := IndexData{Title: config.Title, Columns: config.Columns, Sections: config.Sections, Themes: (*themes), Font: config.Font}
		t, err := template.ParseFiles(getFirstPath([]string{"/etc/bodo/index.html", "/usr/local/share/bodo/index.html"}))
		if err != nil {
			log.Fatalf("Error rendering template %v\n", err)
		}
		t.Execute(w, indexData)
	}
}

func main() {
	flag.StringVar(&configFilePath, "config", "/etc/bodo/config.yaml", "Config path")
	flag.StringVar(&themesFilePath, "themes", "", "Themes path")
	flag.StringVar(&indexFilePath, "index", "", "Index path")
	flag.Parse()
	http.HandleFunc("/", getIndexHandler())
	log.Println("Started bodo service")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
