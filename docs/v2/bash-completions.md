---
tags:
  - v2
search:
  boost: 2
---

* goal
  * enable completion commands 

* steps
  * `App.EnableBashCompletion: true`

#### Default auto-completion

* by default, enabled | app's subcommands
* _Example:_ [here](examples/bash-autocompletion-default.go)
  * `go build examples/bash-autocompletion-default.go`
  * `./bash-autocompletion-default` + tab key
    * Problems:
      * Problem1: Why does bash autocompletion NOT work?
        * Attempt1: 
          * `./bash-autocompletion-default --generate-bash-completion > defaultcompletion.sh`
          * `source defaultcompletion.sh`
        * Attempt2:
          * create manually ["simple_completion.sh"](examples/simple_completion.sh)
          * `source examples/simple_completion.sh`
          * `./bash-autocompletion-default` + tab key
            * NOT work
        * Solution: TODO: ❓

          ![](images/default-bash-autocomplete.gif)

#### Custom auto-completion

* `BashComplete: func ()`
  * write your OWN completion methods -- for --
    * the App OR
    * its subcommands

* _Example:_ [here](examples/bash-autocompletion-custom.go)
  * `go build bash-autocompletion-custom.go`
  * `./bash-autocompletion-custom` + + tab key
    * ❌NOT work❌
  
      ![](images/custom-bash-autocomplete.gif)

#### Enabling

* ["autocomplete/bash_autocomplete"](/autocomplete/bash_autocomplete)
  * bash script /
    * ⚠️| CURRENT (ONLY) shell session,⚠️
      * enable auto-completion 
  * steps to use | SAME shell
    * `export PROG=nameOfYourProgram`
      * _Example:_ `export PROG=bash-autocompletion-default`
    * `source ../../../autocomplete/bash_autocomplete`
      * Problems:
        * Problem1: "bash_autocomplete:7: no matches found: __%[1]s_init_completion"
          * Attempt1: `PROG=bash-autocompletion-default && source ../../../autocomplete/bash_autocomplete`
          * Attempt2: `export PROG=bash-autocompletion-default && source ../../../autocomplete/bash_autocomplete`
          * Solution: TODO: ❓

#### Distribution and Persistent Autocompletion

Copy `autocomplete/bash_autocomplete` into `/etc/bash_completion.d/` and rename
it to the name of the program you wish to add autocomplete support for (or
automatically install it there if you are distributing a package). Don't forget
to source the file or restart your shell to activate the auto-completion.

```sh-session
$ sudo cp path/to/autocomplete/bash_autocomplete /etc/bash_completion.d/<myprogram>
$ source /etc/bash_completion.d/<myprogram>
```

Alternatively, you can just document that users should `source` the generic
`autocomplete/bash_autocomplete` and set `$PROG` within their bash configuration
file, adding these lines:

```sh-session
$ PROG=<myprogram>
$ source path/to/cli/autocomplete/bash_autocomplete
```

Keep in mind that if they are enabling auto-completion for more than one
program, they will need to set `PROG` and source
`autocomplete/bash_autocomplete` for each program, like so:

```sh-session
$ PROG=<program1>
$ source path/to/cli/autocomplete/bash_autocomplete

$ PROG=<program2>
$ source path/to/cli/autocomplete/bash_autocomplete
```

#### Customization

The default shell completion flag (`--generate-bash-completion`) is defined as
`cli.EnableBashCompletion`, and may be redefined if desired, e.g.:

<!-- {
  "args": ["&#45;&#45;generate&#45;bash&#45;completion"],
  "output": "wat\nhelp\nh"
} -->
```go
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "wat",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

#### ZSH Support

Auto-completion for ZSH is also supported using the
`autocomplete/zsh_autocomplete` file included in this repo. One environment
variable is used, `PROG`.  Set `PROG` to the program name as before, and then
`source path/to/autocomplete/zsh_autocomplete`.  Adding the following lines to
your ZSH configuration file (usually `.zshrc`) will allow the auto-completion to
persist across new shells:

```sh-session
$ PROG=<myprogram>
$ source path/to/autocomplete/zsh_autocomplete
```

#### ZSH default auto-complete example
![](images/default-zsh-autocomplete.gif)

#### ZSH custom auto-complete example
![](images/custom-zsh-autocomplete.gif)

#### PowerShell Support

Auto-completion for PowerShell is also supported using the
`autocomplete/powershell_autocomplete.ps1` file included in this repo.

Rename the script to `<my program>.ps1` and move it anywhere in your file
system.  The location of script does not matter, only the file name of the
script has to match the your program's binary name.

To activate it, enter:

```powershell
& path/to/autocomplete/<my program>.ps1
```

To persist across new shells, open the PowerShell profile (with `code $profile`
or `notepad $profile`) and add the line:

```powershell
& path/to/autocomplete/<my program>.ps1
```
