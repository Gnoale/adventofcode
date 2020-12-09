package graph

type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func Bfs(node string, graph map[string][]map[string]int) map[string]int {
	s := make(Stack, 0)
	s = s.Push(node)
	visited := make(map[string]int)
	visited[node] = 0
	for len(s) > 0 {
		var toVisit string
		s, toVisit = s.Pop()
		next := graph[toVisit]
		for _, n := range next {
			for k, _ := range n {
				if _, in := visited[k]; !in {
					s = s.Push(k)
					visited[k] = visited[toVisit] + 1
				}
			}
		}
	}
	return visited
}
