package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"z3ntl3/token-checker/builder"

	"golang.org/x/sync/errgroup"
)

/*
Programmed by Z3NTL3 - pix4.dev

	Fastest Token Checker!!!
*/
type information struct {
	req    *http.Request
	client *http.Client
	token  string
	file   *os.File
}

const (
	API string = "https://discord.com/api/v9/users/@me"
)

var goods int
var good_tokens string

func (client information) CheckToken() error {
	t := client.token
	c := client.client

	t = strings.TrimSpace(t)
	req := client.req

	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.91 Mobile Safari/537.36")
	req.Header.Add("Authorization", t)

	content := make([]byte, 5000000)
	resp, _ := c.Do(req)
	length, _ := resp.Body.Read(content)

	if !strings.Contains(string(content[0:length]), `Unauthorized`) {
		fmt.Println("\033[1m\033[0;97m[INFO] \033[32mGood Token: \033[1m\033[0;97m", t, "\033[0m")
		goods += 1
		good_tokens += t

		client.file.WriteString(t + "\n")
	} else {
		fmt.Println("\033[1m\033[0;97m[INFO] \033[31mBad Token: \033[1m\033[0;97m", t, "\033[0m")
	}
	return nil
}

func main() {
	logo := builder.LogoBuild()
	args := os.Args
	max_worker_count := runtime.NumCPU()
	free_cores := 3

	workers := new(errgroup.Group)
	workers.SetLimit(10000 * (max_worker_count - free_cores))

	fmt.Printf(logo)
	if len(args) != 2 {
		fmt.Println(builder.Usage(), "\n")
		return
	}
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(string(content)) == 0 {
		fmt.Println("\n\033[1m\033[0;97m[INFO] \033[31mAdd tokens into your token file\033[0m")
		return
	}
	clear := strings.Trim(string(content), "\n")
	tokens := strings.Split(clear, "\n")

	fmt.Println()

	f, err := os.OpenFile("goods.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for _, v := range tokens {
		v = strings.TrimSpace(v)
		client := http.Client{}
		request, err := http.NewRequest(http.MethodGet, API, nil)

		if err != nil {
			fmt.Println(err)
		} else {
			info := information{
				req:    request,
				client: &client,
				token:  v,
				file:   f,
			}

			workers.Go(func() error {
				return info.CheckToken()
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := workers.Wait(); err != nil {
		fmt.Println(err)
	}
}
