package Textpkg

func UniqueStrList(origList []string) []string {
	m := map[string]bool{}
	var l []string
	for _, elm := range origList {
		m[elm] = true
	}
	for k, v := range m {
		if v {
			l = append(l, k)
		}
	}
	return l
}
