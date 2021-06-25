## gocopyfw completion

Generate completion script

### Synopsis

To load completions:

Bash:

  $ source <(gocopyfw completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ gocopyfw completion bash > /etc/bash_completion.d/gocopyfw
  # macOS:
  $ gocopyfw completion bash > /usr/local/etc/bash_completion.d/gocopyfw

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ gocopyfw completion zsh > "${fpath[1]}/_gocopyfw"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ gocopyfw completion fish | source

  # To load completions for each session, execute once:
  $ gocopyfw completion fish > ~/.config/fish/completions/gocopyfw.fish

PowerShell:

  PS> gocopyfw completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> gocopyfw completion powershell > gocopyfw.ps1
  # and source this file from your PowerShell profile.


```
gocopyfw completion [bash|zsh|fish|powershell]
```

### Options

```
  -h, --help   help for completion
```

### SEE ALSO

* [gocopyfw](gocopyfw.md)	 - Copy File Watcher for IntelliJ that copies files

###### Auto generated by spf13/cobra on 25-Jun-2021