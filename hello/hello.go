package main

import "fmt"

/* 
原版
func main()  {
	fmt.Println("Hello, world")
}
 */

/*  便於測試將print字串拆解出來 */
/* func Hello() string {
    return "Hello, world"
} */

/*  傳入字串 */
/* 
func Hello(name string) string {
    return "Hello, " + name
} */


/*  
利用常數重構
常量应该可以提高应用程序的性能，它避免了每次调用 Hello 时创建 "Hello, " 字符串实例。
*/
/* const helloPrefix = "Hello, "

func Hello(name string) string {
    return helloPrefix + name
}  */

// 用if做傳入的字串判斷

/* const helloPrefix = "Hello, "

func Hello(name string) string {
    if name == "" {
        name = "World"
    }
    return helloPrefix + name
} */


//更改傳入兩個參數
/* func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    return helloPrefix + name
} */

/* 檢查语言是否是「西班牙语」 */
/* 
const helloPrefix = "Hello, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == "Spanish" {
        return "Hola, " + name
    }

    return helloPrefix + name
} */

/* 將一些字串拉出來當常數使用 */

/* 
const helloPrefix = "Hello, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == "Spanish" {
        return "Hola, " + name
    }

    return helloPrefix + name
} */

/* 新增法語 */
/* const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjunr"

func Hello(name string,language string)  string{
    if name == ""{
        name = "World"
    }
    if language ==spanish {
        return spanishHelloPrefix + name
    }
    if language == french {
        return frenchHelloPrefix + name
    }
    return helloPrefix + name
} */

/* 改用switch */
/* const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjunr"

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    prefix := helloPrefix

    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    }

    return prefix + name
}
 */



/* 
最後一次重構 
使用了 命名返回值（prefix string）
你只需调用 return 而不是 return prefix 即可返回所设置的值

函数名称以小写字母开头。在 Go 中，公共函数以大写字母开始，私有函数以小写字母开头。
我们不希望我们算法的内部结构暴露给外部，所以我们将这个功能私有化。

*/

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjunr, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishPrefix
    }
    return
}

func main() {
	// fmt.Println(Hello("John"))
	fmt.Println(Hello("ss","French"))
}

