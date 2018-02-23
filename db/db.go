package db

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Cache struct {
	BookmarksFile string
}

func userHomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

func validUserHomeDir() error {
	if _, err := os.Stat(userHomeDir()); os.IsNotExist(err) {
		return err
	}
	return nil
}

func New(file string) *Cache {
	if err := validUserHomeDir(); err != nil {
		panic(err)
	}

	bookmarks_file := filepath.Join(userHomeDir(), file)
	if _, err := os.Stat(bookmarks_file); os.IsNotExist(err) {
		_, err := os.Create(bookmarks_file) //TODO return handler to avoid multipe fopen syscalls
		if err != nil {
			panic(err)
		}
	}

	return &Cache{BookmarksFile: bookmarks_file}
}

func (c *Cache) Write(value []byte) error {
	return ioutil.WriteFile(c.BookmarksFile, value, 0644)
}

func (c *Cache) Read() ([]byte, error) {
	return ioutil.ReadFile(c.BookmarksFile)
}

func (db *Cache) EncodeAndWrite(data map[string][]string) {
	y, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}
	db.Write([]byte(y))
}

func (db *Cache) ReadAndDecode() map[string][]string {
	data, err := db.Read()
	if err != nil {
		panic(err)
	}

	bookmarksObj := map[string][]string{}
	err = yaml.Unmarshal([]byte(data), &bookmarksObj)
	if err != nil {
		panic(err)
	}

	return bookmarksObj
}
