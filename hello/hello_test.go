package main

import "testing"


/* 程序需要在一个名为 xxx_test.go 的文件中编写
测试函数的命名必须以单词 Test 开始
测试函数只接受一个参数 t *testing.T */

/* func TestHello(t *testing.T)  {
	got :=Hello()
	want := "Hello, world2"

	if got !=want {
		t.Errorf("got %s,want %s",got,want)
	}
} */


/* 重複代碼的寫法 */

/* func TestHello(t *testing.T) {

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("John")
        want := "Hello, John"

        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    })

    t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    })

} */


/* 
	重構,將t.Errorf的功能拉出來,重複使用 
	對Hello function做測試
	main function裡面的資料不管
*/

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {

		 /* t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
		 通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。 */
		
		 t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris","")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("","")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
	})
	
	t.Run("in Spanish",func(t *testing.T){
		got :=Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
	})
}