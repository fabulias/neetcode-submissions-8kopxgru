func findTheDifference(s string, t string) string {
    sumS, sumT := 0, 0
    for _, c := range s {
        sumS += int(c)
    }
    for _, c := range t {
        sumT += int(c)
    }
    return string(byte(sumT - sumS))
}