package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result //total visits per domain
	domains []string
	total   int
	lines   int
}

func main() {

	p := parser{
		sum: make(map[string]result),
	}

	//scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		//parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input format: %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input format: %q (line #%d)\n", fields[1], p.lines)
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		//Keep track of total and per domain visits
		p.total += visits

		p.sum[domain] = result{domain: domain, visits: visits + p.sum[domain].visits}
	}

	//Print the visits per domain
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parser := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parser.visits)
	}

	//Print the toatal visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	//Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
