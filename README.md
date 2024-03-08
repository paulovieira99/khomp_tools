separar_canais.go: 
Uso: go run separar_canais.go <arquivo_log> <canal_inicial> <canal_final>

Resultado: cria um diretório contendo um arquivo para cada canal no range especificado.

=================================================================================================

filtro_chamada.go:
Uso: go run filtro_chamada.go <diretorio_logs> <data_hora_inicial> <data_hora_final>

Resultado: Exporta para um arquivo de forma organizada as citações dos principais logs dentro do timestamp especificado.

O script é mais eficaz em cenários com fluxo baixo/médio de chamadas, visto que os trechos são pegos exclusivamente pelo timestamp.

Boas práticas: Ao referenciar o timestamp, colocar alguns milésimos a menos no inicial, e a mais no final, pode-se usar a hora exata de algumas mensagens
antes do EV_NEW_CALL, e algumas após o CM_DISCONNECT. O comando fica no seguinte formato:

go run filtro_chamada.go . "02/01/2006 15:04:05.999" "02/01/2006 15:05:03.932"
