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

func findCandidates(dependsOn map[string]map[string]bool, workerCandidates [workerCount]string) []string {
	if len(dependsOn) == 0 {
		return nil
	}
	candidates := []string{}
	for candidate, dependencies := range dependsOn {
		if len(dependencies) == 0 {
			// Is a worker already working this one?
			var inProgress bool
			for _, workerCandidate := range workerCandidates {
				if workerCandidate == candidate {
					inProgress = true
				}
			}
			if !inProgress {
				candidates = append(candidates, candidate)
			}
		}
	}
	sort.Strings(candidates)
	return candidates
}

func removeCandidate(dependsOn map[string]map[string]bool, candidateToDelete string) map[string]map[string]bool {
	delete(dependsOn, candidateToDelete)
	for _, dependencies := range dependsOn {
		delete(dependencies, candidateToDelete)
	}
	return dependsOn
}

const workerCount = 5

func main() {
	lines := aoc.FileToLines("input/7.txt")
	dependsOn := buildDependencyMap(lines)
	workerCandidates := [workerCount]string{}
	workerTime := [workerCount]int{}
	ticks := 0
	for {
		inProgress := false
		for i := 0; i < workerCount; i++ {
			// Pick up work if this worker is free
			if workerTime[i] == 0 {
				// Remove whatever candidate the worker had taken
				dependsOn = removeCandidate(dependsOn, workerCandidates[i])
				// Pick up something if possible
				candidates := findCandidates(dependsOn, workerCandidates)
				if len(candidates) == 0 {
					// No work to be done
				} else {
					candidate := candidates[0]
					candidates = candidates[1:]
					// Candiates are just one char so treat it as an int
					time := int(candidate[0]) - 4
					workerCandidates[i] = candidate
					workerTime[i] = time
				}
			}
			// Decrement the time this worker has been working its current task
			if workerTime[i] != 0 {
				inProgress = true
				workerTime[i]--
			}
		}
		if !inProgress {
			fmt.Println(ticks - 1)
			return
		}
		ticks++
	}
}
