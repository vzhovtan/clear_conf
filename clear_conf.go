package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fname := flag.String("fName", "", "Config file to sanitize")
	flag.Parse()

	file, err := os.Open(*fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inp := scanner.Text()
		sanitize(inp)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func sanitize(inp string) {
	if checkToRemove(inp) {
		if strings.HasSuffix(inp, ";") {
			inpTrim := strings.Replace(inp, ";", "", 1)
			out := buildNewLine(inpTrim)
			fmt.Printf("%s;\n", out)
		} else {
			fmt.Println(buildNewLine(inp))
		}
	}
}

func checkToRemove(inp string) bool {
	re, _ := regexp.Compile(`^\s*description.*;`)
	if strings.Contains(inp, "**") {
		return false
	} else if strings.Contains(inp, "*/") {
		return false
	} else if re.MatchString(inp) {
		return false
	}
	return true
}

func buildNewLine(inp string) string {
	out := []string{}
	inpSlice := strings.Split(inp, " ")
	for _, v := range inpSlice {
		if validIP4add(v) {
			out = append(out, newIPv4(v))
			continue
		} else if validIP6add(v) {
			out = append(out, newIPv6(v))
			continue
		}
		out = append(out, v)
	}
	return fmt.Sprint(strings.Join(out, " "))
}

func validIP4add(ipAddress string) bool {
	re, _ := regexp.Compile(`^((25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(\.|/|$)){4}`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}

func validIP6add(ip6Address string) bool {
	re, _ := regexp.Compile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(\/|$)`)
	if re.MatchString(ip6Address) {
		return true
	}
	return false
}

func newIPv4(addr string) string {
	if strings.Contains(addr, "/") {
		mask := strings.Split(addr, "/")[1]
		oldv4 := strings.Split(addr, "/")[0]
		newv4 := buildNewV4(oldv4)
		return fmt.Sprint(newv4, "/", mask)
	}
	return buildNewV4(addr)
}

func buildNewV4(inp string) string {
	newv4slice := []string{}
	oldv4 := strings.Split(inp, ".")
	if oldv4[0] != "0" {
		for idx, item := range oldv4 {
			if idx == 0 {
				newv4slice = append(newv4slice, fmt.Sprint(item))
			} else {
				chars := strings.Split(item, "")
				randomv4 := randomizeStrSlice(chars)
				v4elem := wrapOver(strings.Join(randomv4, ""))
				newv4slice = append(newv4slice, v4elem)
			}
		}
		return strings.Join(newv4slice, ".")
	}
	return inp
}

func newIPv6(addr string) string {
	if strings.Contains(addr, "/") {
		mask := strings.Split(addr, "/")[1]
		oldv6 := strings.Split(addr, "/")[0]
		newv6 := buildNewV6(oldv6)
		return fmt.Sprint(newv6, "/", mask)
	}
	return buildNewV6(addr)
}

func buildNewV6(inp string) string {
	newv6slice := []string{}
	oldv6 := strings.Split(inp, ":")
	for idx, item := range oldv6 {
		if idx == 0 {
			newv6slice = append(newv6slice, string(item))
		}
		if idx > 0 {
			chars := strings.Split(item, "")
			randomv6 := randomizeStrSlice(chars)
			newv6slice = append(newv6slice, strings.Join(randomv6, ""))
		}
	}
	return strings.Join(newv6slice, ":")
}

func randomizeStrSlice(slice []string) []string {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

func wrapOver(inp string) string {
	num, err := strconv.Atoi(inp)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%d", byte(num))
}
