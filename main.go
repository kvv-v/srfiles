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

// Определяет параметры поиска
func detectArgs() (string, string) {
	startArgs := os.Args
	var dir, srWord string

	if len(startArgs) == 1 {
		dir = "."
		srWord = ""
	} else if len(startArgs) == 2 {
		if strings.HasPrefix(startArgs[1], "/") {
			dir = startArgs[1]
			srWord = ""
		} else {
			dir = "."
			srWord = startArgs[1]
		}
	} else if len(startArgs) == 3 {
		if strings.HasPrefix(startArgs[1], "/") {
			dir = startArgs[1]
			srWord = startArgs[2]
		} else if strings.HasPrefix(startArgs[2], "/") {
			dir = startArgs[2]
			srWord = startArgs[1]
		} else {
			fmt.Println("One argument for path, another for string")
			os.Exit(0)
		}
	} else {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}
	return dir, srWord
}

/*
	Три варианта запуска программы:
	1. С 1 аргументом - именем файла: ведет поиск в папке - месте исполнения
	2. С 1 аргументом - путем нужной папке: выводит все файлы в этой папке
	3. С 2 аргументами - обычный запуск программы
	4. Без аргументов - вывод всех файлов в этой директории
*/
func main() {
	dir, srWord := detectArgs()
	searchFiles(srWord, dir)
}
