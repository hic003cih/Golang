//我们不再使用 main 包，而是定义了一个名为 integers 的包。
package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T)  {
	sum := Add(2,2)
	expected :=4
	ExampleAdd()


	//打印一个整数而不是字符串
	if sum !=expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd()  {
	sum :=Add(1,5)
	fmt.Println(sum)
	// Output: 6
}