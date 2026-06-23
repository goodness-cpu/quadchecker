package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var result strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) {
				result.WriteByte('o')
			} else if i == 1 || i == y {
				result.WriteByte('-')
			} else if j == 1 || j == x {
				result.WriteByte('|')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var result strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				result.WriteByte('/')
			} else if i == 1 && j == x {
				result.WriteByte('\\')
			} else if i == y && j == 1 {
				result.WriteByte('\\')
			} else if i == y && j == x {
				result.WriteByte('/')
			} else if i == 1 || i == y || j == 1 || j == x {
				result.WriteByte('*')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var result strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				result.WriteByte('A')
			} else if i == 1 && j == x {
				result.WriteByte('A')
			} else if i == y && j == 1 {
				result.WriteByte('C')
			} else if i == y && j == x {
				result.WriteByte('C')
			} else if i == 1 || i == y {
				result.WriteByte('B')
			} else if j == 1 || j == x {
				result.WriteByte('B')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var result strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				result.WriteByte('A')
			} else if i == 1 && j == x {
				result.WriteByte('C')
			} else if i == y && j == 1 {
				result.WriteByte('A')
			} else if i == y && j == x {
				result.WriteByte('C')
			} else if i == 1 || i == y {
				result.WriteByte('B')
			} else if j == 1 || j == x {
				result.WriteByte('B')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var result strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				result.WriteByte('A')
			} else if i == 1 && j == x {
				result.WriteByte('C')
			} else if i == y && j == 1 {
				result.WriteByte('C')
			} else if i == y && j == x {
				result.WriteByte('A')
			} else if i == 1 || i == y {
				result.WriteByte('B')
			} else if j == 1 || j == x {
				result.WriteByte('B')
			} else {
				result.WriteByte(' ')
			}
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder

	for scanner.Scan() {
		input.WriteString(scanner.Text())
		input.WriteByte('\n')
	}

	pattern := input.String()

	if pattern == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(strings.TrimSuffix(pattern, "\n"), "\n")
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	y := len(lines)
	x := len(lines[0])

	for _, line := range lines {
		if len(line) != x {
			fmt.Println("Not a quad function")
			return
		}
	}

	if x <= 0 || y <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	var matches []string

	quads := map[string]func(int, int) string{
		"quadA": quadA,
		"quadB": quadB,
		"quadC": quadC,
		"quadD": quadD,
		"quadE": quadE,
	}

	quadNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	for _, name := range quadNames {
		generated := quads[name](x, y)
		if generated == pattern {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, x, y))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
