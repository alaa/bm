# BM is a cli Bookmark tool

## Usage:

```
→ bm

usage: bm [<flags>] <command> [<args> ...]

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  list
    List all available bookmarks.

  find <tag>
    Find bookmark using tag.

  find-url <url>
    Find URL using a substring in the URL.

  open <url>
    Open all URLs that match a substring.

  open-tag <tag>
    Open all URLs that match a Tag

  rename-tag <old-tag> <new-tag>
    Rename tag; applied on all bookmarks.

  add <url> <tags>
    Add new bookmark

  delete-by-url <url>
    Delete a bookmark

  delete-by-tag <tag>
    Delete a bookmark
```
