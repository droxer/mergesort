package mergesort

func Sort(values []int) {
    sort(values, 0, len(values)-1)
}

func sort(values []int, l, r int) {
    if l >= r {
        return
    }

    m := l + (r-l)/2
    sort(values, l, m)
    sort(values, m+1, r)
    merge(values, l, m, r)
}

func merge(values []int, l, m, r int) {
    var i, j, k int
    var l1, r1 = m - l + 1, r - m

    lpart := make([]int, l1)
    rpart := make([]int, r1)

    for i = 0; i <= l1-1; i++ {
        lpart[i] = values[l+i]
    }
    for j = 0; j <= r1-1; j++ {
        rpart[j] = values[m+1+j]
    }

    i, j, k = 0, 0, l
    for i < l1 && j < r1 {
        if lpart[i] <= rpart[j] {
            values[k] = lpart[i]
            i++
        } else {
            values[k] = rpart[j]
            j++
        }
        k++
    }

    for i < l1 {
        values[k] = lpart[i]
        k++
        i++
    }

    for j < r1 {
        values[k] = rpart[j]
        k++
        j++
    }
}
