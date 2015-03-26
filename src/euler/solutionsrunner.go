package euler

type SolutionRunner func(string,SolutionFunc)

func RunSolutions(runner SolutionRunner) {
    for _, solution := range solutions {
        runner(solution.Name, solution.Func)
    }
}