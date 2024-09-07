func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, v := range strs {
       r := []rune(v)

        sort.Slice(r, func (i, j int) bool {
           return r[i] < r[j]
        })
        n := string(r)
        m[n] = append(m[n], v)
    }

    var res [][]string

    for _, v := range m {
        res = append(res, v)
    }
    return res
}