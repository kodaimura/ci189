package chapter1

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
	"unicode/utf8"
)


func Chapter1() {
	c1q1()
	c1q2()
	c1q3()
	c1q4()
	c1q5()
	c1q6()
	c1q7()
	c1q8()
	c1q9()
}


/*
 * Q.1.1
 * 重複のない文字列
*/
func c1q1() {
	fmt.Println(IsUniqueString("abcdef") && !IsUniqueString("abcdea"))
}

func IsUniqueString(str string) bool {
	if str == "" {
		return true
	}
	if strings.Contains(str[1:], str[0:1]) {
		return false
	} else {
		return IsUniqueString(str[1:])
	}
}


/*
 * Q.1.2
 * 順列チェック
*/
func c1q2() {
	fmt.Println(PermutationCheck("abc", "bac") && !PermutationCheck("abc", "abd"))
}

func PermutationCheck(str1, str2 string) bool {
	return sortString(str1) == sortString(str2)
}

func sortString(str string) string {
	sl := strings.Split(str, "")
	sort.Strings(sl)
	return strings.Join(sl, "")
}


/*
 * Q.1.3
 * URLify
文字列の全ての空白文字を"%20"で置き換え

func ReplaceSpace(str, x string) string {
	return strings.Replace(str, " ", x, -1)
}
ReplaceSpace(str, "%20")
*/
func c1q3() {
	fmt.Println(ReplaceSpace("a b c", "%20") == "a%20b%20c")
}

func ReplaceSpace(str, x string) string {
	if str == "" {
		return ""
	}
	if str[0:1] == " " {
		return x + ReplaceSpace(str[1:], x) 
	} else {
		return str[0:1] + ReplaceSpace(str[1:], x) 
	}
}


/*
 * Q.1.4
 * 回文の順列
*/
func c1q4() {
	fmt.Println(IsPalindromePermutation("ab cab") && !IsPalindromePermutation("ab cbb"))
}

func IsPalindromePermutation(str string) bool {
	m := map [rune]int{}
	for _, r := range str {
		if r != 32 && r != 12288 { //半角・全角スペースはカウントしない
			m[r] += 1
		}
	}
	countOdd := 0
	for _, v := range m {
		if v % 2 != 0 {
			countOdd += 1
			if countOdd == 2 {
				return false
			}
		}
	}
	
	return true 
}


/*
 * Q.1.5
 * 一発変換
挿入、削除、置換のどれか1操作以下で文字列が一致するか判定
pale, ple -> true
pale, bale -> true
pale, bake -> false
*/
func c1q5() {
	fmt.Println(
		IsPossibleToMatchInOneShot("pale", "ple") &&
		IsPossibleToMatchInOneShot("pale", "bale") &&
		!IsPossibleToMatchInOneShot("pale", "bake"))
}

func IsPossibleToMatchInOneShot(str1, str2 string) bool {
	rs1, rs2 := []rune(str1), []rune(str2)
	len1, len2 := len(rs1), len(rs2)

	if len1 > len2 {
		rs1, rs2 = rs2, rs1
		len1, len2 = len2, len1
	}

	if len2 - len1 >= 2 {
		return false
	} 

	countDiff := 0
	shift := 0

	for i := 0; i < len1; i++ {
		if rs1[i] != rs2[i + shift] {
			if countDiff += 1; countDiff == 2 {
				return false
			}
			if len1 != len2 {
				shift += 1
			}
		}
	}
	return true
}


/*
 * Q.1.6
 * 文字列圧縮
aabcccccaa -> a2b1c5a3
(ただし、元の方が短い場合は元を返す)
abb -> a1b2 -> abb
*/
func c1q6() {
	fmt.Println(CompressString("aabcccccaaa") == "a2b1c5a3" && CompressString("abb") == "abb")
}

func CompressString(str string) string {
	sl := strings.Split(str, "")
	ret := ""

	pre := sl[0]
	count := 1
	for _, c := range sl[1:] {
		if c == pre {
			count += 1
		} else {
			ret += pre + strconv.Itoa(count)
			pre = c
			count = 1
		}
	}
	ret += pre + strconv.Itoa(count)

	if utf8.RuneCountInString(str) <= utf8.RuneCountInString(ret) {
		return str
	}
	return ret
}


/*
 * Q.1.7
 * 行列の回転
N × N の行列を90°回転
1 2 3      7 4 1
4 5 6  ->  8 5 2
7 8 9      9 6 3
*/

func c1q7() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(Rotation90(matrix))
}

func Rotation90(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix))

	for _, row := range matrix {
		for i, v := range row {
			ret[i] = append([]int{v}, ret[i]...) 
		}
	}
	return ret
}

/*
 * Q.1.8
 * ゼロの行列
M × N の行列で要素が0であれば、その行と列の全ての要素を0にする
1 2 3      1 0 3
4 0 6  ->  0 0 0
7 8 9      7 0 9
*/
func c1q8() {
	sl := [][]int{{1,2,3},{4,0,6},{7,8,9}}
	fmt.Println(SpreadZeroToRowAndCol(sl))
}

func SpreadZeroToRowAndCol(matrix [][]int) [][]int {
	ret := DeepCopy(matrix)

	for y, row := range matrix {
		for x, v := range row {
			if v == 0 {
				setZeroToRowAndCol(ret, x, y)
			} 
		}
	}

	return ret
}


func setZeroToRowAndCol(matrix [][]int, x, y int) {
	for i, _ := range matrix[y] {
		matrix[y][i] = 0
	}

	for i, _ := range matrix {
		matrix[i][x] = 0
	}

}


func DeepCopy(matrix[][]int) [][]int {
	c := make([][]int, len(matrix)) 
	for i, row := range matrix {
		c[i] = make([]int, len(row))
		copy(c[i], row)
	}

	return c
}

/* 
 * Q.1.9
 * 文字列の回転
 2つの文字列s1,s2について、s2がs1を回転させたものかを判定する
 waterbottle, erbottlewat -> true
*/
func c1q9() {
	fmt.Println(IsRotateString("waterbottle", "erbottlewat") && !IsRotateString("waterbottle", "rebottlewat"))
}


func IsRotateString(str1, str2 string) bool {
	len := utf8.RuneCountInString(str2)

	for i := 0; i < len - 1; i++ {
		if str2[i:] + str2[0:i] == str1 {
			return true
		}
	}
	return false
}
