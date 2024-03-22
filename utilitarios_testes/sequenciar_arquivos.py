import os
import shutil

        # Created by Paulo Vieira
        # paulo.vieira@khomp.com

def modify_and_save_file(input_file, output_file, old_number, new_number):
    with open(input_file, 'r') as infile:
        content = infile.read()
    modified_content = content.replace(str(old_number), str(new_number))
    with open(output_file, 'w') as outfile:
        outfile.write(modified_content)

def main():
    old_filename = "register_201.xml"
    start_number = 201
    end_number = 210

    for number in range(start_number, end_number + 1):
        new_filename = f"register_{number}.xml"
        duplicate_filename = f"register_{number}_2.xml"

        # Modifica e salva o conteúdo do novo arquivo
        modify_and_save_file(old_filename, new_filename, start_number, number)

        # Copia o conteúdo do arquivo original para o arquivo duplicado
        shutil.copy(new_filename, duplicate_filename)

        # Imprime a mensagem para o arquivo e sua duplicata
        print(f"Arquivo '{new_filename}' e sua duplicata '{duplicate_filename}' foram criados e modificados!")

if __name__ == "__main__":
    main()
