package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")

	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("input: ", input)
	}
}

/*
inputReader 是一个指向 bufio.Reader 的指针。inputReader := bufio.NewReader(os.Stdin) 这行代码，将会创建一个读取器，并将其与标准输入绑定。

bufio.NewReader() 构造函数的签名为：func NewReader(rd io.Reader) *Reader

该函数的实参可以是满足 io.Reader 接口的任意对象（任意包含有适当的 Read() 方法的对象，请参考章节11.8），函数返回一个新的带缓冲的 io.Reader 对象，它将从指定读取器（例如 os.Stdin）读取内容。

返回的读取器对象提供一个方法 ReadString(delim byte)，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。

ReadString 返回读取到的字符串，如果碰到错误则返回 nil。如果它一直读到文件结束，则返回读取到的字符串和 io.EOF。如果读取过程中没有碰到 delim 字符，将返回错误 err != nil。

在上面的例子中，我们会读取键盘输入，直到回车键（\n）被按下。


*/
