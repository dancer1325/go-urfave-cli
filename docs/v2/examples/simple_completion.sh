#!/bin/bash

_bash_autocompletion_default() {
    local cur opts
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts="add a complete c template t help h"
    
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
}

complete -F _bash_autocompletion_default ./bash-autocompletion-default