#!/bin/bash

# Build the example
cd docs/v2/examples
go build bash-autocompletion-default.go

# Generate and source completion
./bash-autocompletion-default --generate-bash-completion > completion.sh
source completion.sh

echo "Autocompletado instalado. Prueba:"
echo "./bash-autocompletion-default <TAB><TAB>"