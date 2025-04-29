package main
import "fmt"

func search(nums []int, target int) int {
    var n, i int
    n = len(nums)

    for i = 0; i < n; i++ {
        if nums[i] == target {
            return i
        }
    }
    return -1
}

func main() {
    var n, target int
    fmt.Scan(&n)

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
    }
    fmt.Scan(&target)

    hasil := search(nums, target)
    fmt.Println("Hasil:", hasil)
}
