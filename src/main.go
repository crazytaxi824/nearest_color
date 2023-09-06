package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"local/src/color"
)

func main() {
	log.SetFlags(log.Llongfile)

	// 获取 args
	hexColors := os.Args[1:]

	if len(hexColors) < 1 {
		// 如果 args 不存在则检查是否有 pipe stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			p := scanner.Text()
			hexColors = append(hexColors, p)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(2)
		}
	}

	for _, h := range hexColors {
		if len(h) != 7 || h[0] != '#' {
			fmt.Println("\"" + h + "\", format error: hex color format should be #xxxxxx")
			continue
		}

		c, err := color.ParseHEX(h[1:])
		if err != nil {
			fmt.Println("\""+h+"\", parse hex color error:", err)
			continue
		}

		c8b, nc := c.FindNearest8bitColor()
		fmt.Println(h, c8b, nc.ToHEX())
	}
}
