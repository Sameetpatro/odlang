package investigate

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	green  = "\033[38;2;0;255;0m"   // #00FF00
	dim    = "\033[2;38;2;0;200;0m" // softer #00FF00 for status lines
	reset  = "\033[0m"
	cursor = "\033[?25h"
	hide   = "\033[?25l"
)

const totalQuestions = 7

var reader = bufio.NewReader(os.Stdin)

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func enableGreen() {
	fmt.Print(green)
}

func disableColor() {
	fmt.Print(reset + cursor)
}

func pause(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func typeLine(text string, charDelayMs int) {
	for _, ch := range text {
		fmt.Print(string(ch))
		delay := charDelayMs
		if ch == '\n' {
			delay = charDelayMs / 2
		}
		pause(delay)
	}
	fmt.Println()
}

func typeBlock(lines []string, charDelayMs int) {
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			pause(80)
			continue
		}
		typeLine(line, charDelayMs)
		pause(50)
	}
}

func typePause(text string) {
	typeLine(text, 16)
	pause(180)
}

func logStatus(text string) {
	fmt.Print(dim)
	typeLine("  "+text+"...", 8)
	fmt.Print(green)
	pause(120)
}

func progressBar(label string, width int, steps int) {
	fmt.Print("  ")
	if label != "" {
		fmt.Print(label + " ")
	}
	for i := 1; i <= steps; i++ {
		filled := (width * i) / steps
		bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
		fmt.Printf("\r  [%s] %3d%%", bar, (i*100)/steps)
		pause(22 + rand.Intn(12))
	}
	fmt.Println()
}

func bracketProgress(label string, total int) {
	fmt.Print("  ")
	if label != "" {
		fmt.Print(label + " ")
	}
	width := 26
	for i := 1; i <= total; i++ {
		filled := (width * i) / total
		bar := strings.Repeat("#", filled) + strings.Repeat(".", width-filled)
		fmt.Printf("\r  [%s]", bar)
		pause(28 + rand.Intn(15))
	}
	fmt.Println()
}

func decryptText(final string, rounds int) {
	runes := []rune(final)
	scramble := "█▓▒░@#$%&*?0123456789ABCDEF"
	buf := make([]rune, len(runes))
	for i := range buf {
		buf[i] = rune(scramble[rand.Intn(len(scramble))])
	}
	for r := 0; r < rounds; r++ {
		for i := range buf {
			if r > i*rounds/len(runes) || rand.Float32() < 0.35 {
				buf[i] = runes[i]
			} else {
				buf[i] = rune(scramble[rand.Intn(len(scramble))])
			}
		}
		fmt.Printf("\r  %s", string(buf))
		pause(25)
	}
	fmt.Printf("\r  %s\n", final)
}

func scanSequence(lines []string) {
	for _, line := range lines {
		logStatus(line)
		bracketProgress("", 5 + rand.Intn(3))
	}
}

func questionLabel(n int) string {
	return fmt.Sprintf("  Question %d of %d", n, totalQuestions)
}

func prompt(text string) string {
	fmt.Println()
	typeLine(text, 14)
	fmt.Print("\n  > ")
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func promptYN(n int, text string) bool {
	for {
		label := text + "\n\n  (Y/N)"
		if n > 0 {
			label = questionLabel(n) + "\n\n  " + label
		} else {
			label = "  " + label
		}
		answer := strings.ToLower(prompt(label))
		switch answer {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			typePause("  Invalid input. Y or N only.")
		}
	}
}

func promptChoice(n int, text string, choices []string) int {
	fmt.Println()
	typeLine(questionLabel(n), 14)
	pause(100)
	typeLine("  "+text, 14)
	for i, c := range choices {
		typeLine(fmt.Sprintf("  %d  %s", i+1, c), 10)
		pause(40)
	}
	for {
		fmt.Print("\n  > ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		num := 0
		_, err := fmt.Sscanf(line, "%d", &num)
		if err == nil && num >= 1 && num <= len(choices) {
			return num - 1
		}
		typePause("  Pick a number from the list.")
	}
}

func waitEnter(message string) {
	fmt.Println()
	typeLine(message, 14)
	reader.ReadString('\n')
}

func banner(title string, width int) {
	line := strings.Repeat("█", width)
	fmt.Println()
	typeLine(line, 3)
	typeLine(title, 4)
	typeLine(line, 3)
	fmt.Println()
}
