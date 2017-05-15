package maindomain

import (
	"sort"
	"strings"
)

func reverse(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}
	return append(reverse(a[1:]), a[0])
}

// FindMainDomain returns the main domain of a domain.
func FindMainDomain(domain string) string {
	tokens := strings.Split(domain, ".")
	if len(tokens) <= 2 {
		// like alauda.cn
		return domain
	}
	depth := findMainDomainDepth(reverse(tokens), PublicSuffixTree, 1)
	return strings.Join(tokens[len(tokens)-depth:], ".")
}

type node struct {
	name string
	ch   []node
}

func searchNode(name string, ch []node) int {
	idx := sort.Search(len(ch), func(idx int) bool {
		return ch[idx].name >= name
	})
	if idx < len(ch) && ch[idx].name == name {
		return idx
	}
	return -1
}

func findMainDomainDepth(tokens []string, n node, depth int) int {
	if len(tokens) == 0 {
		return -0xFFFF
	}
	if n.ch == nil {
		return depth
	}
	idx := searchNode(tokens[0], n.ch)
	if idx < 0 {
		return depth
	}
	return findMainDomainDepth(tokens[1:], n.ch[idx], depth+1)
}

