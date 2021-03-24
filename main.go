/*
	Поиск по файловой системе. в аргумент берет строку - путь к папке, в которой нужно искать,
	какую-либо строку и выводит все найденые файлы, содержащие эту строку. Рекурсивно ищет во всех папках в исходной папке
*/
package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

// ввод строки пользователем
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(nil)
	}

	str = strings.TrimSpace(str)

	return str
}

// принимает путь к папке, возвращает срез со списком всех файлов и папок в этой папке
func getDirList(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

// Почти главная функция, принимает поисковой запрос и выводит найденые по нему файлы
func searchFiles(str, dir string) {
	dirList := getDirList(dir)
	for _, file := range dirList {
		if file.IsDir() { // Если находит папку, то функция рекурсивно вызывается уже для этой папки
			searchFiles(str, dir+"/"+file.Name())
		}
		if strings.Contains(file.Name(), str) { // Сам поиск. Если находит файл, содержащий в имени нужную строку
			fmt.Println(dir, "\t", file.Name()) // Выводит имя этого файла и путь к нему
		}
	}
}

func main() {

	fmt.Print("Enter path to directory: ")
	dir := readLine()

	fmt.Print("Enter string: ")
	srWord := readLine()

	searchFiles(srWord, dir)
}
