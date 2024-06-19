package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	if err := os.MkdirAll("downloaded", os.ModePerm); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.Chdir("downloaded"); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: go run task.go 'URL'")
		os.Exit(1)
	}

	baseURL := args[0]
	baseDir := ""

	if len(args) >= 2 {
		baseDir = filepath.Join(baseDir, args[1])

		if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
		}
	}

	outputDir := filepath.Join(baseDir, getDirectoryName(baseURL))

	var wg sync.WaitGroup
	var mu sync.Mutex

	mapLinks := make(map[string]bool, 1000)
	mapLinks[baseURL] = true

	file, err := wget(baseURL, baseURL, baseDir, outputDir, &wg, &mu, mapLinks, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Complete downloading URL - %v\nCreated file: %v\n", baseURL, file)
	}
}

func getDirectoryName(URL string) string {
	URL = strings.TrimSuffix(URL, `/`)

	absURL, err := url.Parse(URL)
	if err != nil {
		fmt.Printf("Error rapsing URL %v: %v\n", URL, err)
	}

	directoryName := absURL.Host + strings.ReplaceAll(absURL.Path, `/`, "_")

	return directoryName
}

func wget(baseURL, URL, baseDir, outputDir string, wg *sync.WaitGroup, mu *sync.Mutex, mapLinks map[string]bool, depth int) (string, error) {
	if depth > 10 {
		return "", nil
	}

	file, err := downloadFile(baseURL, URL, baseDir, outputDir, mu)
	if err != nil {
		return "", fmt.Errorf("error downloanig %v: %v", URL, err)
	}

	if file == "" {
		return "", nil
	}

	if filepath.Ext(file) == ".html" && depth <= 0 {
		parsedBaseURL, err := url.Parse(baseURL)
		if err != nil {
			fmt.Printf("Error parsing baseURL: %v\n", err)
		}

		parsedURL, err := url.Parse(URL)
		if err != nil {
			fmt.Printf("Error parsing URL: %v\n", err)
		}

		if parsedBaseURL.Host != parsedURL.Host {
			depth += 10
		}

		if baseURL == URL {
			if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
				fmt.Printf("Error creating folder: %v\n", err)
			}
		}

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %v: %v\n", file, err)
		}

		newData := string(data)

		newURL := "./" + strings.Replace(file, baseDir+`\`, `/`, -1)

		if baseDir != "" {
			newURL = "./" + strings.Replace(file, baseDir+`\`, "", 1)
		}

		urlQuote := `"` + URL + `"`
		newURLQuote := `"` + newURL + `"`

		if strings.Contains(newData, urlQuote) {
			mu.Lock()
			newData = strings.ReplaceAll(newData, urlQuote, newURLQuote)
			mu.Unlock()
		}

		resp, err := http.Get(URL)
		if err != nil {
			return "", fmt.Errorf("error getting response from URL: %v", err)
		}
		defer resp.Body.Close()

		links, err := parseHTML(resp.Body)
		if err != nil {
			fmt.Printf("Error parsing URL: %v\n", err)
		}

		for _, link := range links {
			oldLink := link

			if strings.HasPrefix(link, `#`) {
				continue
			}

			if len(link) > 2 {
				baseURL = strings.TrimSuffix(baseURL, `/`)
				link = strings.TrimSuffix(link, `/`)
			}

			parsedBaseURL, err := url.Parse(baseURL)
			if err != nil {
				fmt.Printf("Error parsing baseURL: %v\n", err)
			}

			parsedURL, err := url.Parse(link)
			if err != nil {
				fmt.Printf("Error parsing URL: %v\n", err)
			}

			if parsedURL.Scheme == "" {
				parsedURL = parsedBaseURL.ResolveReference(&url.URL{Path: link})

				tempURL := strings.ReplaceAll(parsedURL.String(), `%3F`, `?`)

				parsedURL, err = url.Parse(tempURL)
				if err != nil {
					fmt.Printf("Error parsing URL: %v\n", err)
				}
			}

			if strings.HasPrefix(parsedURL.String(), "mailto") ||
				strings.HasPrefix(parsedURL.String(), "tel") ||
				strings.HasPrefix(parsedURL.String(), "tg") {
				continue
			}

			if parsedBaseURL.String() == parsedURL.String() && len(link) > 2 {
				continue
			}

			if (mapLinks[parsedURL.String()] || mapLinks[oldLink]) && len(link) > 2 {
				continue
			}

			if !strings.HasSuffix(baseURL, "/") {
				baseURL = baseURL + "/"
			}

			wg.Add(1)

			mapLinks[oldLink] = true

			go func(link, oldLink string) {
				defer wg.Done()

				newFile, err := wget(baseURL, link, baseDir, outputDir, wg, mu, mapLinks, depth+1)
				if err != nil {
					fmt.Println(err)
				}

				if newFile != "" {
					newLink := "./" + strings.Replace(newFile, baseDir+`\`, `/`, -1)

					if baseDir != "" {
						newLink = "./" + strings.Replace(newFile, baseDir+`\`, "", 1)
					}

					linkQuote := `"` + oldLink + `"`
					newLinkQuote := `"` + newLink + `"`

					if strings.Contains(newData, linkQuote) {
						mu.Lock()
						newData = strings.ReplaceAll(newData, linkQuote, newLinkQuote)
						mu.Unlock()
					}
				}
			}(parsedURL.String(), oldLink)
		}
		wg.Wait()

		if newData != "" {
			if err := os.WriteFile(file, []byte(newData), os.ModePerm); err != nil {
				fmt.Printf("Error writing data in file %v: %v\n", file, err)
			}
		}
	}

	return file, nil
}

func downloadFile(baseURL, URL, baseDir, outputDir string, mu *sync.Mutex) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fileName, err := getFileName(baseURL, URL, baseDir, outputDir, resp)
	if err != nil {
		return "", err
	}

	fileName, oldFileNames := getUniqueFileName(fileName)

	if fileName == "" {
		return "", nil
	}

	err = saveFile(fileName, resp)
	if err != nil {
		return "", err
	}

	for _, oldFileName := range oldFileNames {
		if oldFileName == "" {
			continue
		}

		equal, err := compareFiles(fileName, oldFileName)
		if err != nil {
			return "", err
		}

		if equal {
			if err = os.Remove(fileName); err != nil {
				return "", err
			}

			fileName = oldFileName
			break
		}
	}

	return fileName, nil
}

func getFileName(baseURL, URL, baseDir, outputDir string, resp *http.Response) (string, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL %v: %v", baseURL, err)
	}

	parsedURL, err := url.Parse(URL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL %v: %v", URL, err)
	}

	contentType := resp.Header.Get("Content-Type")
	contentDisposition := resp.Header.Get("Content-Disposition")

	var fileName string

	if strings.Contains(contentDisposition, "filename=") {
		results := strings.Split(contentDisposition, "filename=")
		result := strings.Trim(results[len(results)-1], `"`)

		if parsedURL.String() == parsedBaseURL.String() {
			fileName = filepath.Join(baseDir, result)
		} else {
			fileName = filepath.Join(outputDir, result)
		}
	} else {
		if parsedURL.String() == parsedBaseURL.String() {
			fileName = outputDir
		} else if parsedURL.Path == "" || parsedURL.Path == `/` {
			fileName = filepath.Join(outputDir, strings.ReplaceAll(parsedURL.Host, ":", "_"))
		} else {
			fileName = filepath.Join(outputDir, filepath.Base(parsedURL.Path))

			if strings.Contains(fileName, "?") {
				fileName = strings.Split(fileName, "?")[0]
			}
		}

		extension := filepath.Ext(parsedURL.Path)

		if strings.Contains(extension, "?") {
			extension = strings.Split(extension, "?")[0]
		}

		if extension == "" {
			extensions := strings.Split(contentType, "/")

			if len(extensions) == 2 {
				if strings.Contains(extensions[1], ";") {
					extension = strings.Split(extensions[1], ";")[0]
				} else {
					extension = extensions[1]
				}
			} else {
				return "", fmt.Errorf("incorrect Content-Type: %v", contentType)
			}
		}

		if !strings.Contains(fileName, extension) {
			fileName += "." + extension
		}
	}

	return fileName, nil
}

func getUniqueFileName(fileName string) (string, []string) {
	idx := strings.LastIndex(fileName, ".")

	if strings.Count(fileName, ".") == 1 {
		idx = len(fileName) - 1
	}

	fileBody := fileName[:idx]
	fileExt := fileName[idx:]

	if half := len(fileBody) / 2; half > 90 {
		fileBody = fileBody[half:]
		fileName = fileBody + fileExt
	}

	oldFileNames := make([]string, 4)

	for i := 1; ; i++ {
		if i > 2 {
			return "", nil
		}

		_, err := os.Stat(fileName)
		if err == nil {
			oldFileNames[i-1] = fileName

			num := strconv.Itoa(i)

			fileName = fileBody + "(" + num + ")" + fileExt

			continue
		}

		break
	}

	return fileName, oldFileNames
}

func saveFile(fileName string, resp *http.Response) error {
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("file %v already exists", fileName)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file %v", fileName)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying file %v", fileName)
	}
	file.Close()

	return nil
}

func compareFiles(fileName, oldFileName string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, fmt.Errorf("error opening file %v", fileName)
	}

	data, err := os.ReadFile(file.Name())
	if err != nil {
		return false, fmt.Errorf("error reading file %v", fileName)
	}
	file.Close()

	oldFile, err := os.Open(oldFileName)
	if err != nil {
		return false, fmt.Errorf("error opening file %v", oldFileName)
	}

	oldData, err := os.ReadFile(oldFile.Name())
	if err != nil {
		return false, fmt.Errorf("error reading file %v", oldFileName)
	}
	oldFile.Close()

	ignoreLine := `<!-- page generated`

	stringData := string(data)
	stringOldData := string(oldData)

	if strings.Contains(stringOldData, ignoreLine) && strings.Contains(stringData, ignoreLine) {
		stringOldData = strings.Split(stringOldData, ignoreLine)[0] + `\n`
		stringData = strings.Split(stringData, ignoreLine)[0] + `\n`

		oldData = []byte(stringOldData)
		data = []byte(stringData)
	}

	if bytes.Equal(oldData, data) {
		return true, nil
	}

	return false, nil
}

func parseHTML(body io.Reader) ([]string, error) {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			for _, attr := range token.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		}
	}
}
