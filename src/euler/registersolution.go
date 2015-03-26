package euler

import "sort"

type SolutionFunc func()(uint64,error)

type solutionType struct {
    Name string
    Func SolutionFunc
}

type solutionsList []solutionType

var solutions solutionsList

func (s solutionsList) Len() int {
    return len(s)
}

func (s solutionsList) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s solutionsList) Less(i, j int) bool {
    return s[i].Name < s[j].Name
}

func RegisterSolution(name string, solution SolutionFunc) {
    solutions = append(solutions, solutionType{name, solution})
    sort.Sort(solutions)
}

func init() {
    solutions = make(solutionsList, 0, 100)
}
