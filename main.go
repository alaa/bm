package main

import (
	"fmt"
	"strings"

	"github.com/alaa/bookmarks/browser"
	"github.com/alaa/bookmarks/db"
	"github.com/fatih/color"
	"github.com/ryanuber/columnize"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	listCommand            = kingpin.Command("list", "List all available bookmarks")
	findCommand            = kingpin.Command("find", "Find bookmark using tags")
	findCommandTag         = findCommand.Arg("tag", "Tag keyword to find.").Required().String()
	findUrlCommand         = kingpin.Command("find-url", "Find URL using a keyword in it")
	findUrlCommandUrl      = findUrlCommand.Arg("url", "Keyword in the URL to find.").Required().String()
	openUrlCommand         = kingpin.Command("open", "Open URL by search substring in the stored URLs")
	openUrlCommandUrl      = openUrlCommand.Arg("url", "Keyword in the URL to find.").Required().String()
	openTagCommand         = kingpin.Command("open-tag", "Open URL by Tag")
	openTagCommandTags     = openTagCommand.Arg("tag", "Keyword in the URL to find.").Required().String()
	renameTagCommand       = kingpin.Command("rename-tag", "Rename tag; applied on all data")
	renameTagCommandOldTag = renameTagCommand.Arg("old-tag", "The old tag name to be replaced").Required().String()
	renameTagCommandNewTag = renameTagCommand.Arg("new-tag", "The new tag name").Required().String()
	addCommand             = kingpin.Command("add", "Add new bookmark")
	addCommandUrl          = addCommand.Arg("url", "URL to add.").Required().String()
	addCommandTags         = addCommand.Arg("tags", "Tags to add.").Required().String()
	deleteUrlCommand       = kingpin.Command("delete-by-url", "Delete a bookmark")
	deleteUrlCommandUrl    = deleteUrlCommand.Arg("url", "URL to delete.").Required().String()
	deleteTagCommand       = kingpin.Command("delete-by-tag", "Delete a bookmark")
	deleteTagCommandTag    = deleteTagCommand.Arg("tag", "URL to delete.").Required().String()
)

func main() {
	opts := kingpin.Parse()
	db := db.New(".bookmarks_file")
	data := db.ReadAndDecode()

	switch opts {

	case "list":
		_ = *listCommand
		output := []string{}
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		for url, tags := range data {
			output = append(output, fmt.Sprintf("%s | %v\n", yellow(url), green(tags)))
		}
		fmt.Println(columnize.SimpleFormat(output))

	case "find":
		_ = *findCommand
		output := []string{}
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		for url, tags := range data {
			if contains(*findCommandTag, tags) {
				output = append(output, fmt.Sprintf("%s | %v\n", yellow(url), green(tags)))
			}
		}
		fmt.Println(columnize.SimpleFormat(output))

	case "open":
		_ = *findUrlCommand
		for url, _ := range data {
			if strings.Contains(url, *openUrlCommandUrl) {
				browser.OpenURL(url)
			}
		}

	case "open-tag":
		_ = *openTagCommand
		keywords := strings.Split(*openTagCommandTags, ",")
		for url, tags := range data {
			for _, keyword := range keywords {
				if contains(keyword, tags) {
					browser.OpenURL(url)
				}
			}
		}

	case "rename-tag":
		_ = *openTagCommand
		old := *renameTagCommandOldTag
		new := *renameTagCommandNewTag
		new_tags := strings.Split(new, ",")

		for url, tags := range data {
			if contains(old, tags) {
				tags = remove(old, tags)
				for _, new_tag := range new_tags {
					tags = append_if_uniq(new_tag, tags)
					data[url] = uniqueNonEmpty(tags)
				}
			}
		}
		db.EncodeAndWrite(data)

	case "find-url":
		_ = *findUrlCommand
		output := []string{}
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		for url, tags := range data {
			if strings.Contains(url, *findUrlCommandUrl) {
				output = append(output, fmt.Sprintf("%s | %v\n", yellow(url), green(tags)))
			}
		}
		fmt.Println(columnize.SimpleFormat(output))

	case "add":
		urls := strings.Split(*addCommandUrl, "\n")
		tags := strings.Split(*addCommandTags, ",")
		for _, url := range urls {
			data[url] = uniqueNonEmpty(tags)
		}
		db.EncodeAndWrite(data)

	case "delete-by-url":
		_ = *deleteUrlCommand
		_, ok := data[*deleteUrlCommandUrl]
		if ok {
			delete(data, *deleteUrlCommandUrl)
			db.EncodeAndWrite(data)
		} else {
			red := color.New(color.FgRed).SprintFunc()
			msg := fmt.Sprintf("%s is not in the bookmarks db", *deleteUrlCommandUrl)
			fmt.Println(red(msg))
		}

	case "delete-by-tag":
		_ = *deleteTagCommand
		keywords := strings.Split(*deleteTagCommandTag, ",")
		for url, tags := range data {
			for _, keyword := range keywords {
				if contains(keyword, tags) {
					delete(data, url)
				}
			}
			db.EncodeAndWrite(data)
		}
	}
}
