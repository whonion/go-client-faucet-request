package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	addresses, err := readLines("addresses.txt")
	if err != nil {
		fmt.Println("Error reading addresses:", err)
		return
	}

	proxies, err := readLines("proxy.txt")
	if err != nil {
		fmt.Println("Error reading proxies:", err)
		return
	}

	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	userAgents, err := readLines("useragents.txt")
	if err != nil {
		fmt.Println("Error reading user agents:", err)
		return
	}

	if len(addresses) != len(proxies) {
		fmt.Println("Error: the number of addresses must match the number of proxies.")
		return
	}

	for i, address := range addresses {
		proxy := proxies[i]
		var httpClient *http.Client
		if strings.Contains(proxy, "http://") {
			proxyURL, err := url.Parse(proxy)
			if err != nil {
				fmt.Println("Error parsing proxy URL:", err)
				continue
			}
			httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
		} else if strings.Contains(proxy, "https://") {
			proxyURL, err := url.Parse(proxy)
			if err != nil {
				fmt.Println("Error parsing proxy URL:", err)
				continue
			}
			httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
		} else {
			httpClient = &http.Client{}
		}

		reqBody := fmt.Sprintf(`{"address": "%s", "coins": ["1000000000unibi","1000000000000unusd"]}`, address)
		req, err := http.NewRequest("POST", "https://faucet.itn-1.nibiru.fi/", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		// Set a random user agent in the headers
		req.Header.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		fmt.Printf("Response for address %s through proxy %s: %s\n", address, proxy, string(body))
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
