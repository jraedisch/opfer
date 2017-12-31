package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	names := readNames("players.csv")
	shuffled := shuffleNames(names)
	victims := distributeVictims(shuffled)
	fmt.Println("\n\n\n\nPress Enter when ready!")
	readInput()
	for _, murderer := range names {
		victim := victims[murderer]
		clearScreen()
		fmt.Printf("\n\n\n\nNext victim is for %s's eyes only! (press Enter if you are %s)", murderer, murderer)
		readInput()
		fmt.Printf("Your victim is: %s\nPress Enter if you wrote it down!", victim)
		readInput()
	}
	fmt.Println("All victims have been distributed :)")
}

func distributeVictims(names []string) map[string]string {
	max := len(names)
	victims := map[string]string{}
	for i, name := range names {
		next := int(math.Mod(float64(i+1), float64(max)))
		victims[name] = names[next]
	}
	return victims
}

func readNames(fileName string) []string {
	f, err := os.Open(fileName)
	panicOnError(err)
	lines, err := csv.NewReader(f).ReadAll()
	panicOnError(err)
	names := []string{}
	for _, l := range lines {
		name := strings.TrimSpace(l[0])
		if name != "" {
			names = append(names, name)
		}
	}
	return names
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func shuffleNames(names []string) []string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	l := len(names)
	shuffled := make([]string, l)
	perm := r.Perm(l)
	for i, v := range perm {
		shuffled[i] = names[v]
	}
	return shuffled
}
