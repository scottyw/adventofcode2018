package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/scottyw/adventofcode2018/pkg/aoc"
)

func buildDependencyMap(lines []string) map[string]map[string]bool {
	// Step Z must be finished before step B can begin.
	r := regexp.MustCompile("Step (\\w) must be finished before step (\\w) can begin.")
	dependsOn := map[string]map[string]bool{}
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		before := matches[1]
		after := matches[2]
		if _, ok := dependsOn[before]; !ok {
			// Make sure every step is included in the map
			dependsOn[before] = map[string]bool{}
		}
		if _, ok := dependsOn[after]; !ok {
			dependsOn[after] = map[string]bool{}
		}
		dependsOn[after][before] = true
	}
	return dependsOn
}

func findCandidate(dependsOn map[string]map[string]bool) string {
	if len(dependsOn) == 0 {
		return ""
	}
	candidates := []string{}
	for candidate, dependencies := range dependsOn {
		if len(dependencies) == 0 {
			candidates = append(candidates, candidate)
		}
	}
	if len(candidates) == 0 {
		panic("No candidates found")
	}
	sort.Strings(candidates)
	return candidates[0]
}

func removeCandidate(dependsOn map[string]map[string]bool, candidateToDelete string) map[string]map[string]bool {
	delete(dependsOn, candidateToDelete)
	for _, dependencies := range dependsOn {
		delete(dependencies, candidateToDelete)
	}
	return dependsOn
}

func main() {
	lines := aoc.FileToLines("input/7.txt")
	dependsOn := buildDependencyMap(lines)
	for {
		candidate := findCandidate(dependsOn)
		if candidate == "" {
			return
		}
		fmt.Printf(candidate)
		dependsOn = removeCandidate(dependsOn, candidate)
	}
}
