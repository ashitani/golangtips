/*
正規表現（パターンマッチ）
*/

package tips_regexp

import (
	"fmt"
	"regexp"
)

//---------------------------------------------------
// 正規表現を使う
//---------------------------------------------------
/*
正規表現はregexpパッケージを使用します。
regexp.Compile()か regexp.MustCompile()で初期化・コンパイルを行います。
MustCompile()にすると、正規表現の解析に失敗したときはpanicを起こします。

なお、後述しますが、正規表現はバッククォーテーション(\`\`)で囲むのがよいです。
*/
//import "regexp"

func regexp_Regexp() {
	r := regexp.MustCompile(`golang`)
	fmt.Println(r) // =>"golang"
}

//---------------------------------------------------
// 文字にマッチさせる
//---------------------------------------------------
/*
文字列にマッチするかどうかを確認するには
(*Regexp) MatchString()　を使います。
使える正規表現は[RE2 Standard](https://github.com/google/re2/wiki/Syntax)とのことです。

正規表現などを記載する文字列リテラルにダブルクオーテーション(\")を使う場合はエスケープを忘れずに。
特殊文字を一切解釈しない「未加工文字列リテラル」はシングルクオーテーション(\')ではなく
バッククォーテーション(\`)を使うので注意。シングルクオーテーションはrune(UTF8文字）を表す表記です。
たとえば、数値マッチさせる場合には
```
	r = regexp.MustCompile('\d')    // OK
	r = regexp.MustCompile("\\d")   // OK
	r = regexp.MustCompile(`\d`)    // NG
```
のような表記を使います。基本はバッククォーテーションということでよいでしょう。
*/
//import "regexp"

func regexp_Match() {
	r := regexp.MustCompile(`abc`)
	fmt.Println(r.MatchString("hello"))     // => false
	fmt.Println(r.MatchString("hello abc")) // => true
}

//---------------------------------------------------
// 繰り返し文字とマッチさせる
//---------------------------------------------------
/*
\*,\+,\?などで繰り返しも対応できます。
*/
// import "regexp"

func regexp_Repeat() {
	check_regexp(`a*c`, "abc")     // => "true"
	check_regexp(`a*c`, "ac")      // => "true"
	check_regexp(`a*c`, "aaaaaac") // => "true"
	check_regexp(`a*c`, "c")       // => "true"
	check_regexp(`a*c`, "abccccc") // => "true"
	check_regexp(`a*c`, "abd")     // => "false"

	check_regexp(`a+c`, "abc")     // => "false"
	check_regexp(`a+c`, "ac")      // => "true"
	check_regexp(`a+c`, "aaaaaac") // => "true"
	check_regexp(`a+c`, "c")       // => "false"
	check_regexp(`a+c`, "abccccc") // => "false"
	check_regexp(`a+c`, "abd")     // => "false"

	check_regexp(`a?c`, "abc")     // => "true"
	check_regexp(`a?c`, "ac")      // => "true"
	check_regexp(`a?c`, "aaaaaac") // => "true"
	check_regexp(`a?c`, "c")       // => "true"
	check_regexp(`a?c`, "abccccc") // => "true"
	check_regexp(`a?c`, "abd")     // => "false"
}

func check_regexp(reg, str string) {
	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
}

//---------------------------------------------------
// 数字だけ・アルファベットだけとマッチさせる
//---------------------------------------------------
// import "regexp"

func regexp_NumAlpha() {
	check_regexp(`[ABZ]`, "A") // => "true"
	check_regexp(`[ABZ]`, "Z") // => "true"
	check_regexp(`[ABZ]`, "Q") // => "false"

	check_regexp(`[0-9]`, "5") // => "true"
	check_regexp(`[0-9]`, "A") // => "false"
	check_regexp(`[A-Z]`, "A") // => "true"
	check_regexp(`[A-Z]`, "5") // => "false"
	check_regexp(`[A-Z]`, "a") // => "false"

	check_regexp(`[^0-9]`, "A") // => "true"
	check_regexp(`[^0-9]`, "5") // => "false"
}

//func check_regexp(reg, str string) {
//	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
//}

//---------------------------------------------------
// 改行コードを含む文字列にマッチさせる
//---------------------------------------------------
/*
改行コードを含む複数ラインに対するマッチを行うには、
正規表現の冒頭に`(?m)`と書くことで対応できます。
*/
// import "regexp"

func regexp_MultiLine() {
	txt := "hello,\nworld"
	check_regexp(`(?m)^(w.*)$`, txt)
}

//func check_regexp(reg, str string) {
//	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
//}

//---------------------------------------------------
// 正規表現を使って文字列を置き換える
//---------------------------------------------------
/*
ReplaceAllString()を使います。
*/
// import "regexp"

func regexp_Replace() {
	str := "Copyright 2015 by ASHITANI Tatsuji."
	rep := regexp.MustCompile(`[A-Za-z]*right`)
	str = rep.ReplaceAllString(str, "Copyleft")

	fmt.Println(str) // => "Copyleft 2015 by ASHITANI Tatsuji."
}

//---------------------------------------------------
// n番めのマッチを見つける
//---------------------------------------------------
/*
見つけたグループを指定しつつ置換をしたい場合は`$1`などを使います。
*/
// import "regexp"

func regexp_Numbering() {
	str := "123456"
	rep := regexp.MustCompile(`1(.)3(.)5(.)`)
	str = rep.ReplaceAllString(str, "1($1)3($2)5($3)")

	fmt.Println(str) // => "1(2)3(4)5(6)"
}

//---------------------------------------------------
// パターンで区切られたレコードを読む
//---------------------------------------------------
/*
Split(s string, n int)を使います。nは回数指定で、負を指定すれば全て行います。
*/
// import "regexp"

func regexp_Split() {
	str := "001,ASHITANI Tatsuji, Yokohama"
	rep := regexp.MustCompile(`\s*,\s*`)
	result := rep.Split(str, -1)

	fmt.Println(result[0]) // = >"001"
	fmt.Println(result[1]) // = >"ASHITANI Tatsuji"
	fmt.Println(result[2]) // = >"Yokohama"
}

//---------------------------------------------------
// マッチした文字列を全て抜き出して配列へ格納する
//---------------------------------------------------
/*
FindAllStringSubmatch()を使います。

グルーピングを指定すると、[[マッチ全体, 1つ目のグループ、２つめのグループ、、], ...]
のようなスライスが返ります。
*/
// import "regexp"

func regexp_FindAll() {
	s := "hoge:0045-111-2222 boke:0045-222-2222"
	r := regexp.MustCompile(`[\d\-]+`)
	fmt.Println(r.FindAllStringSubmatch(s, -1)) // => "[[0045-111-2222] [0045-222-2222]]"

	r2 := regexp.MustCompile(`(\S+):([\d\-]+)`)

	result := r2.FindAllStringSubmatch(s, -1)
	fmt.Println(result[0]) // => "[hoge:0045-111-2222 hoge 0045-111-2222]"
	fmt.Println(result[1]) // => "[boke:0045-222-2222 boke 0045-222-2222]"
}

//---------------------------------------------------
// 正規表現にコメントを付ける
//---------------------------------------------------
/*
rubyには正規表現を見やすくするxオプションがありますが、golangにはなさそうなので、
自前で実装します。といっても空白とコメントを削除するだけですが。

*/
// import "regexp"

func regexp_Comment() {
	r := `(\S+):       # 名前
    	 ([\d\-]+)    # 電話番号
    `
	str := "hoge:0045-111-2222 boke:0045-222-2222"
	result := find_x(r, str)
	fmt.Println(result[0]) // => "[hoge:0045-111-2222 hoge 0045-111-2222]"
	fmt.Println(result[1]) // => "[boke:0045-222-2222 boke 0045-222-2222]"
}

func find_x(r, str string) [][]string {
	com := regexp.MustCompile(`(?m)(\s+)|(\#.*$)`)
	r = com.ReplaceAllString(r, "")
	reg := regexp.MustCompile(r)
	return reg.FindAllStringSubmatch(str, -1)
}

//---------------------------------------------------
// 正規表現内でString型変数を使う
//---------------------------------------------------
/*
Rubyと違って正規表現リテラルがあるわけではないので、
stringを編集してcompile()です。
*/
// import "regexp"

func regexp_String() {
	regEx := `^a`
	reg := regexp.MustCompile(regEx)
	fmt.Println(reg.FindAllStringSubmatch("abc", -1)) // =>[[a]]
}

//---------------------------------------------------
// 正規表現（パターンマッチ）
//---------------------------------------------------
func Tips_regexp() {
	regexp_Regexp()    // 正規表現を使う
	regexp_Match()     // 文字にマッチさせる
	regexp_Repeat()    // 繰り返し文字とマッチさせる
	regexp_NumAlpha()  // 数字だけ・アルファベットだけとマッチさせる
	regexp_MultiLine() // 改行コードを含む文字列にマッチさせる
	regexp_Replace()   // 正規表現を使って文字列を置き換える
	regexp_Numbering() // n番めのマッチを見つける
	regexp_Split()     // パターンで区切られたレコードを読む
	regexp_FindAll()   // マッチした文字列を全て抜き出して配列へ格納する
	regexp_Comment()   // 正規表現にコメントを付ける
	regexp_String()    // 正規表現内でString型変数を使う

}
