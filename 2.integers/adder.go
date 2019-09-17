package integers


//当你有多个相同类型的参数（在我们的例子中是两个整数）
//可以将它缩短为 (x，y int) 而不是 (x int, y int)。
// Add takes two integers and returns the sum of them
func Add(x, y int) int {
    return x + y
}