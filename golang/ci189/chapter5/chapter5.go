package chapter5

import (
	"fmt"
	"strconv"
)

func Chapter5() {
	c5q1()
	c5q2()
}

//指定位置のbitが1かどうか
func GetBit(num, i int) bool {
	return ((num & (1 << i)) != 0)
}

//指定位置のbitを1にする
func SetBit(num, i int) int {
	return num | (1 << i)
}

//指定位置のbitを0にする
func ClearBit(num, i int) int {
	mask := ^(1 << i)
	return num & mask
}

/* Q.5.1
 * 挿入
Nのjビット目からiビット目にMを挿入する。
jとiの幅はMのビット数と一致するとする。
N = 10000000000, M = 10011
i = 2, j = 6

=> 10001001100
 */
func c5q1() {
	fmt.Println(InsertBit(1024, 19, 2, 6) == 1100)
}

func InsertBit(n, m, i, j int) int {
	for x := i; x <= j; x++ {
		n = ClearBit(n, x)
	}
	return n | m << i
}


/* Q.5.2
 * 実数の2進数表記
0 ~ 1までの実数を2進数表記
32bit以内で表せない場合は"ERROR"
*/
func c5q2() {
	fmt.Println(FloatToBinaryNumber(0.2) == "ERROR" && FloatToBinaryNumber(0.25) == "0.010")
}

func FloatToBinaryNumber(n float64) string {
	if n >= 1 || n <= 0{
		return "ERROR"
	}
	n = n * 2
	ret := "0." + strconv.Itoa(int(n))
	for i := 0; i < 30; i++ {
		if n >= 1 {
			n = (n - 1) * 2
		} else if n > 0 {
			n = n * 2
		} else if n == 0 {
			return ret
		}
		ret += strconv.Itoa(int(n))
	}
	return "ERROR"
}


/* Q.5.3
 * ベストな反転位置
整数の1ビットだけ0->1に反転できる時、
この操作をした場合に一番1の並びが長くなる長さを求める
*/
/*
func c5q3() {
	fmt.Println(MaxLenSerialThenSetBit(1775) == 8)
}

func CountBitSerialOne() {}

func MaxLenSerialThenSetBit(n int) int {}
*/