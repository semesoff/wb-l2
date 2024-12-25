package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

// extractLinks извлекает все ссылки из HTML-документа
func extractLinks(url string, body io.Reader) ([]string, error) {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" || t.Data == "link" || t.Data == "img" || t.Data == "script" {
				for _, a := range t.Attr {
					if a.Key == "href" || a.Key == "src" {
						link := a.Val
						if strings.HasPrefix(link, "/") {
							link = url + link
						}
						links = append(links, link)
					}
				}
			}
		}
	}
}

// downloadFile создает файл по указанному пути и копирует в него содержимое ответа сервера
func downloadFile(url, filepath string) error {
	// получаем ответ от сервера
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// создаем файл для сохранения по указанному пути
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// копируем тело ответа в файл
	_, err = io.Copy(out, response.Body)
	return err
}

func downloadPage(url string, visited map[string]bool) error {
	// проверяем, не посещали ли мы уже эту страницу
	if visited[url] {
		return nil
	}
	// помечаем страницу как посещенную
	visited[url] = true

	// получаем ответ от сервера
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	// закрываем тело ответа при выходе из функции
	defer response.Body.Close()

	// создаем папку для сохранения файла
	dir := path.Dir(url)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// сохранение html-файла
	filepath := path.Join(dir, path.Base(url))
	if err := downloadFile(url, filepath); err != nil {
		return err
	}

	// Извлекаем ссылки и скачиваем их рекурсивно
	links, err := extractLinks(url, response.Body)
	if err != nil {
		return err
	}
	for _, link := range links {
		if err := downloadPage(link, visited); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wget <url>")
		return
	}
	url := os.Args[1]
	visited := make(map[string]bool)
	if err := downloadPage(url, visited); err != nil {
		fmt.Println(err)
	}
}
