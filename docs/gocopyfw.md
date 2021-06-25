## gocopyfw

Copy File Watcher for IntelliJ that copies files

### Synopsis

Copy File Watcher for IntelliJ that just copies files

To setup the watcher you need to follow these steps:

  1. File -> Settings
  2. Tools -> File Watchers
  3. Click the + at the top right
  4. Enter the following details:
     * Name: gocopyfw
     * File Type: Any
     * Scope: Project Files
     * Program: gocopyfw (or the path to the binary)
     * Arguments: -d $ProjectFileDir$ -f $FilePath$ -l <your filesets>
       * a fileset is a list of files that should remain the same separated by a ';' (semi-colon)
       * filesets can be separated by commas, i.e. each group of files you want to remain in sync
  5. Advanced Options: All Unchecked
  6. Show Console: On Error


```
gocopyfw [flags]
```

### Options

```
  -f, --file-path string          the file path from $FilePath$
  -h, --help                      help for gocopyfw
  -l, --linked-files strings      a list of files that are linked, each set of files is separated by commas and each file in the set is separated by a semicolon
  -d, --project-file-dir string   the project root directory from $ProjectFileDir$
```

### SEE ALSO

* [gocopyfw completion](gocopyfw_completion.md)	 - Generate completion script
* [gocopyfw markdown](gocopyfw_markdown.md)	 - Generate Markdown docs for the command

###### Auto generated by spf13/cobra on 25-Jun-2021