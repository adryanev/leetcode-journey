func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, v := range strs {
        n := sortWord(v)
        m[n] = append(m[n], v)
    }

    res := make([][]string, 0)

    for _, v := range m {
        res = append(res, v)
    }
    return res
}

func sortWord(str string) string {
     r := []rune(str)

        sort.Slice(r, func (i, j int) bool {
           return r[i] < r[j]
        })
       return string(r)

        
}