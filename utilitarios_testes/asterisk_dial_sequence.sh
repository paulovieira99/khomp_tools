#!/bin/bash

# Verifica se os parâmetros foram passados
if [ $# -ne 3 ]; then
  echo "Uso: $0 <ramal_inicial> <ramal_final> <contexto>"
  exit 1
fi

# Parâmetros
ramal_inicial=$1
ramal_final=$2
contexto=$3

# Loop para executar o comando com cada ramal no intervalo
for (( ramal=$ramal_inicial; ramal<=$ramal_final; ramal++ ))
do
  echo "Executando: asterisk -rx 'console dial ${ramal}@${contexto}'"
  asterisk -rx "console dial ${ramal}@${contexto}"
done
