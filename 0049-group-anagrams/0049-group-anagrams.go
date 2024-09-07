func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, v := range strs {
       r := []byte(v)

        sort.Slice(r, func (i, j int) bool {
           return r[i] < r[j]
        })
        n := string(r)
        m[n] = append(m[n], v)
    }

    res := make([][]string, 0)

    for _, v := range m {
        res = append(res, v)
    }
    return res
}