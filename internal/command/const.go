package command

const rootCmdLong = `Copy File Watcher for IntelliJ that just copies files

To setup the watcher you need to follow these steps:

  1. File -> Settings
  2. Tools -> File Watchers
  3. Click the + at the top right
  4. Enter the following details:
     * Name: gocopyfw
     * File Type: Any
     * Scope: Project Files
     * Program: gocopyfw (or the path to the binary)
     * Arguments: -d $ProjectFileDir$ -f $FilePath$ -l [your filesets]
       * a fileset is a list of files that should remain the same separated by a ';' (semi-colon)
       * filesets can be separated by commas, i.e. each group of files you want to remain in sync
  5. Advanced Options: All Unchecked
  6. Show Console: On Error
`

const completionCmdLong = `To load completions:

#### Bash:

  $ source <(gocopyfw completion bash)

  To load completions for each session, execute once:
  ##### Linux:
  $ gocopyfw completion bash > /etc/bash_completion.d/gocopyfw
  ##### macOS:
  $ gocopyfw completion bash > /usr/local/etc/bash_completion.d/gocopyfw

#### Zsh:

  If shell completion is not already enabled in your environment,
  you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  To load completions for each session, execute once:

  $ gocopyfw completion zsh > "${fpath[1]}/_gocopyfw"

  You will need to start a new shell for this setup to take effect.

#### fish:

  $ gocopyfw completion fish | source

  To load completions for each session, execute once:

  $ gocopyfw completion fish > ~/.config/fish/completions/gocopyfw.fish

#### PowerShell:

  PS> gocopyfw completion powershell | Out-String | Invoke-Expression

  To load completions for every new session, run:

  PS> gocopyfw completion powershell > gocopyfw.ps1

  and source this file from your PowerShell profile.
`
