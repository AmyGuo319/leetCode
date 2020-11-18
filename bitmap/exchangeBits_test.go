package bitmap

import "fmt"

//面试题 05.07. 配对交换
//配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位0与位1交换，位2与位3交换，以此类推）。
//
//示例1:
//
//输入：num = 2（或者0b10）
//输出 1 (或者 0b01)
//示例2:
//
//输入：num = 3
//输出：3
//提示:
//
//num的范围在[0, 2^30 - 1]之间，不会发生整数溢出。

func exchangeBits(num int) int {
	//取出奇位数放在
	odd := (num & 0xaaaaaaaa) >> 1
	//取出偶位数：
	even := (num & 0x55555555) << 1

	return odd | even
}

func Example_exchangeBits() {
	fmt.Println(exchangeBits(2))
	fmt.Println(exchangeBits(3))
	fmt.Println(exchangeBits(6))
	//Output:
	//1
	//3
	//9
}
