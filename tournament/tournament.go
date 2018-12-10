package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

var validResults = map[string]bool{
	"win":  true,
	"draw": true,
	"loss": true,
}

// Tally tallies up results of a football competition.
func Tally(r io.Reader, w io.Writer) error {
	counts := map[string]map[string]int{
		"MP": map[string]int{}, // matches
		"W":  map[string]int{}, // wins
		"L":  map[string]int{}, // losses
		"D":  map[string]int{}, // draws
		"P":  map[string]int{}, // points
	}
	scanner := bufio.NewScanner(r)

	// scan each line to tally results
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\n" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Split(line, ";")
		if len(fields) < 3 {
			return errors.New("Invalid input")
		}
		t1, t2, result := fields[0], fields[1], fields[2]
		if !validResults[result] {
			return errors.New("Invalid input")
		}

		counts["MP"][t1] = counts["MP"][t1] + 1
		counts["MP"][t2] = counts["MP"][t2] + 1
		if result == "draw" {
			counts["D"][t1] = counts["D"][t1] + 1
			counts["D"][t2] = counts["D"][t2] + 1
			counts["P"][t1] = counts["P"][t1] + 1
			counts["P"][t2] = counts["P"][t2] + 1
			continue
		}

		var winner string
		var loser string

		if result == "win" {
			winner, loser = t1, t2
		} else {
			winner, loser = t2, t1
		}
		counts["W"][winner] = counts["W"][winner] + 1
		counts["P"][winner] = counts["P"][winner] + 3
		counts["L"][loser] = counts["L"][loser] + 1
		counts["P"][loser] = counts["P"][loser] + 0
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// sort teams by total points
	type Team struct {
		Name   string
		Points int
	}
	var Teams []Team

	for team, points := range counts["P"] {
		Teams = append(Teams, Team{team, points})
	}
	sort.Slice(Teams, func(i, j int) bool {
		if Teams[i].Points == Teams[j].Points {
			return Teams[i].Name < Teams[j].Name
		}
		return Teams[i].Points > Teams[j].Points
	})

	// write results
	line := "%-30s | %2v | %2v | %2v | %2v | %2v\n"
	fmt.Fprintf(w, line, "Team", "MP", "W", "D", "L", "P")
	for _, team := range Teams {
		fmt.Fprintf(w, line, team.Name,
			counts["MP"][team.Name],
			counts["W"][team.Name],
			counts["D"][team.Name],
			counts["L"][team.Name],
			counts["P"][team.Name],
		)
	}
	return nil
}
