package kata

import "strings"

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}

	lines := strings.Split(s, "\n")
	var out []string

	for _, line := range lines {
		scaled := scaleLine(line, k)
		for i := 0; i < n; i++ {
			out = append(out, scaled)
		}
	}

	return strings.Join(out, "\n")
}

func scaleLine(line string, k int) string {
	var b strings.Builder
	for _, r := range line {
		b.WriteString(strings.Repeat(string(r), k))
	}
	return b.String()
}
