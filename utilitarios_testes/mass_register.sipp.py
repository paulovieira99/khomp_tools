#!/usr/bin/env python3
# script para registrar vários ramais via comandos sipp.
# Created by Paulo Vieira - paulo.vieira@khomp.com

import subprocess
import time
from concurrent.futures import ThreadPoolExecutor

# Função para executar o comando Sipp em um novo processo
def execute_sipp(scenario, port):
    subprocess.run(["sipp", "172.27.66.11:5060", "-p", str(port), "-sf", f"{scenario}.xml", "-m", "1"])
    time.sleep(3)

# Obter o timestamp Unix da data desejada (04 de agosto de 2023)
#data_atingir = int(time.mktime(time.strptime("04 Aug 2023", "%d %b %Y")))

# Obter o timestamp Unix da data atual
#data_atual = int(time.time())

# Verificar se a data desejada ainda não foi atingida
#if data_atual < data_atingir:
 #   while data_atual < data_atingir:
  #      time.sleep(1)
   #     data_atual = int(time.time())

# Lista dos arquivos XML de cenários
scenarios = [
    "register_201", "register_201_2", "register_202", "register_202_2", "register_203",
    "register_203_2", "register_204", "register_204_2", "register_205", "register_205_2",
    "register_206", "register_206_2", "register_207", "register_207_2", "register_208",
    "register_208_2", "register_209", "register_209_2", "register_210", "register_210_2"
]

# Porta inicial
port = 57337

# Executar o comando Sipp para cada cenário usando ThreadPoolExecutor
with ThreadPoolExecutor() as executor:
    futures = {executor.submit(execute_sipp, scenario, port + i): scenario for i, scenario in enumerate(scenarios)}
    for future in futures:
        future.result()

exit(0)
