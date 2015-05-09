/*
文字列
*/

package tips_string

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	. "github.com/MakeNowJust/heredoc/dot"
	. "golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//---------------------------------------------------
// 文字列を結合する
//---------------------------------------------------
func string_Concat() {
	a := "Hello"
	b := a + " World"
	fmt.Println(b)
}

//---------------------------------------------------
// 繰り返し文字列を生成する
//---------------------------------------------------
//  import "strings"

func string_Repeat() {
	s := "Hey "
	fmt.Println(strings.Repeat(s, 3))
}

//---------------------------------------------------
// 大文字・小文字に揃える
//---------------------------------------------------
//import "strings"

func string_UpperLower() {
	s := "I love GoLang"
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
}

//---------------------------------------------------
// 大文字と小文字の入れ替え
//---------------------------------------------------
//import "strings"

func string_ReplaceUpperLower() {
	//うーん、かなりイマイチ
	s := "i lOVE gOLANG"
	sr := ""
	for _, x := range s {
		xs := string(x)
		u := strings.ToUpper(xs)
		l := strings.ToLower(xs)
		if u == xs {
			sr += l
		} else if l == xs {
			sr += u
		} else {
			sr += xs
		}
	}
	fmt.Println(sr)
}

//---------------------------------------------------
// コマンドの実行結果を文字列に
//---------------------------------------------------
// import "os/exec"

func string_Exec() {
	out, _ := exec.Command("date").Output()
	fmt.Println(string(out))
}

//---------------------------------------------------
// 複数行の文字列を作成する
//---------------------------------------------------
func string_HereDocument() {
	s := `
	This is a test.

	GoLang, programming language developed at Google.
	`
	fmt.Println(s)
}

//---------------------------------------------------
// ヒアドキュメントの終端文字列をインデントする
//---------------------------------------------------
/*
そういうパッケージがありました。
*/
// import . "github.com/MakeNowJust/heredoc/dot"

func string_HereDocumentIndent() {
	s := D(`
	This is a test.

	GoLang, programming language developed at Google.	
	`)
	fmt.Println(s)
}

//---------------------------------------------------
// 複数行のコマンドの実行結果を文字列に設定する
//---------------------------------------------------
/*
 exec.Command()を使います。
(echoはシェル組み込みコマンドのせいか実行されない。。)
*/
// import "os/exec"

func string_ExecMultiLine() {
	s := D(`
	date
	echo "-----------------------------"
	ps
	`)
	outs := ""
	for _, s = range strings.Split(s, "\n") {
		out, _ := exec.Command(s).Output()
		outs += string(out)
	}
	fmt.Println(outs)
}

//---------------------------------------------------
// 部分文字列を取り出す
//---------------------------------------------------
func string_Extract() {
	s := "Apple Banana Orange"
	fmt.Println(s[0:5])       // => Apple
	fmt.Println(s[6:(6 + 6)]) // => Banana
	fmt.Println(s[0:3])       // => App
	fmt.Println(s[6])         // => 66
	fmt.Println(s[13:19])     // => Orange
}

//---------------------------------------------------
// 部分文字列を置き換える
//---------------------------------------------------
/* (うーんいまいち。。) */
//import "strings"

func string_ReplacePart() {
	s := "Apple Banana Orange"
	ss := strings.Split(s, "")
	srep := strings.Split("Vine ", "")
	for i, _ := range srep {
		ss[i] = srep[i]
	}

	fmt.Println(strings.Join(ss, ""))
}

//---------------------------------------------------
// 文字列中の式を評価し値を展開する
//---------------------------------------------------
/* intなら%dでもよいですが、型推定をよさげに行うのなら%vが使えます。*/
func string_Eval() {
	value := 123
	fmt.Printf("value is %v\n", value)
}

//---------------------------------------------------
// 文字列を1文字ずつ処理する
//---------------------------------------------------
/*
[こちら](http://knightso.hateblo.jp/entry/2014/06/24/090719)によると、
string[n]だとbyte単位、rangeで回すとrune(utf-8文字)単位らしいです。
*/
func string_Each() {
	sum := 0
	for _, c := range "Golang" {
		sum = sum + int(c)
	}
	fmt.Println(sum)
}

//---------------------------------------------------
// 文字列の先頭と末尾の空白文字を削除する
//---------------------------------------------------
/*空白ならTrimSpace, 任意の文字でやりたければTrimでやれます。*/
//import "strings"

func string_Trim() {
	s := "   Hello, Golang!   "
	s = strings.TrimSpace(s)
	fmt.Println(s)
}

//---------------------------------------------------
// 文字列を整数に変換する (to_i)
//---------------------------------------------------
//import "strconv"

func string_ToI() {
	i := 1
	s := "999"
	si, _ := strconv.Atoi(s)
	i = i + si
	fmt.Println(i)
}

//---------------------------------------------------
// 文字列を浮動小数点に変換する (to_f)
//---------------------------------------------------
//import "strconv"

func string_ToF() {
	s := "10"
	sf, _ := strconv.ParseFloat(s, 64) //64 bit float
	fmt.Println(sf)
}

//---------------------------------------------------
// 8進文字列を整数に変換する
//---------------------------------------------------
//import "strconv"

func string_ParseOct() {
	s := "010"
	so, _ := strconv.ParseInt(s, 8, 64) // base 8, 64bit
	fmt.Println(so)
}

//---------------------------------------------------
// 16進文字列を整数に変換する
//---------------------------------------------------
//import "strconv"

func string_ParseHex() {
	s := "ff" // 0xは含んではいけません。
	sh, _ := strconv.ParseInt(s, 16, 64)
	fmt.Println(sh)
}

//---------------------------------------------------
// ASCII文字をコード値に（コード値をASCII文字に）変換する
//---------------------------------------------------
func string_AtoI() {
	s := "ABC"
	fmt.Println(s[0])
	fmt.Println(string(82))
}

//---------------------------------------------------
// 文字列を中央寄せ・左詰・右詰する
//---------------------------------------------------
//import "strings"

func string_Just() {
	//右詰、左詰めはformat文字列が対応しています。±で指定。
	s := "Go"
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)

	//センタリングはなさそうです。
	l := 10
	ls := (l - len(s)) / 2
	cs := strings.Repeat(" ", ls) + s + strings.Repeat(" ", l-(ls+len(s)))
	fmt.Println(cs)
}

//---------------------------------------------------
// "次"の文字列を取得する
//---------------------------------------------------
func string_Succ() {
	fmt.Println(succ("9"))    // => "10"
	fmt.Println(succ("a"))    // => "b"
	fmt.Println(succ("AAA"))  // => "AAB"
	fmt.Println(succ("A99"))  // => "B00"
	fmt.Println(succ("A099")) // => "A100"
}

// 文字列を反転して返す
func reverse(s string) string {
	ans := ""
	for i, _ := range s {
		ans += string(s[len(s)-i-1])
	}
	return string(ans)
}

// "次"の文字列を取得する
func succ(s string) string {
	r := reverse(s)
	ans := ""
	carry := 1
	lastLetter := string(r[0])
	for i, _ := range r {
		lastLetter = string(r[i])
		a := lastLetter
		if carry == 1 {
			if lastLetter == "z" {
				a = "a"
				carry = 1
			} else if lastLetter == "Z" {
				a = "A"
				carry = 1
			} else if lastLetter == "9" {
				a = "0"
				carry = 1
			} else {
				if r[i] == 0 {
					a = "1"
				} else {
					a = string(r[i] + 1)
					carry = 0
				}
			}
		}
		ans += a
	}
	if carry == 1 {
		if lastLetter == "9" {
			ans += "1"
		} else if lastLetter == "z" {
			ans += "a"
		} else if lastLetter == "Z" {
			ans += "A"
		}
	}
	return reverse(ans)
}

//---------------------------------------------------
// 文字列を暗号化する
//---------------------------------------------------
/*とりあえずMD5あたりの例を書いておきます。*/
//import     "crypto/md5"
//import     "io"
//import     "bufio"

func string_Crypt() {
	h := md5.New()
	io.WriteString(h, "hogehoge")

	fmt.Print("input password >")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	h2 := md5.New()
	io.WriteString(h2, scanner.Text())

	if h2.Sum(nil)[0] == h.Sum(nil)[0] {
		fmt.Println("right")
	} else {
		fmt.Println("wrong")
	}

}

//---------------------------------------------------
// 文字列中で指定したパターンにマッチする部分を置換する
//---------------------------------------------------
//import "strings"

func string_Replace() {
	s := "Apple Banana Apple Orange"
	s = strings.Replace(s, "Apple", "Pine", 1) // 最後の引数は回数
	fmt.Println(s)
	s = strings.Replace(s, "Apple", "Pine", -1) // <0で無制限
	fmt.Println(s)
}

//---------------------------------------------------
// 文字列中に含まれている任意文字列の位置を求める
//---------------------------------------------------
//import "strings"

func string_Find() {
	s := "Apple Banana Apple Orange"
	fmt.Println(strings.Index(s, "Apple"))  // => 0
	fmt.Println(strings.Index(s, "Banana")) // => 6
	// 途中から検索する方法は無いのでスライスで渡す
	fmt.Println(strings.Index(s[6:], "Apple") + 6) // => 13
	// 後方検索はLastIndex()
	fmt.Println(strings.LastIndex(s, "Apple"))          // => 13
	fmt.Println(strings.Index(s[:(len(s)-6)], "Apple")) // => 0
}

//---------------------------------------------------
// 文字列の末端の改行を削除する
//---------------------------------------------------
//import "strings"

func string_Chomp() {
	s := "Hello, Golang!\n"
	s = strings.TrimRight(s, "\n")
	fmt.Println(s)
}

//---------------------------------------------------
// カンマ区切りの文字列を扱う
//---------------------------------------------------
//import "strings"

func string_Split() {

	s := "001,ASHITANI Tatsuji,Yokohama"
	slice := strings.Split(s, ",")
	for _, x := range slice {
		fmt.Println(x)
	}
}

//---------------------------------------------------
// 任意のパターンにマッチするものを全て抜き出す
//---------------------------------------------------
// import "regexp"

func string_FindAll() {
	s := "hoge:045-111-2222 boke:045-222-2222"
	re, _ := regexp.Compile("(\\S+):([\\d\\-]+)")
	ans := re.FindAllStringSubmatch(s, -1) // [マッチした全体,１個目のカッコ,２個目のカッコ,..]の配列
	fmt.Println(ans)
}

//---------------------------------------------------
// 漢字コードを変換する
//---------------------------------------------------
/*
下記はEUCの例です。
指定するコーディングは、EUCJP,ISO2022JP,ShiftJISのどれかです。
*/
//import "golang.org/x/text/encoding/japanese"
//import "golang.org/x/text/transform"
//import "strings"
//import "io"
//import "os"
//import "bytes"

func string_Kconv() {
	string_Kconv_fwrite()      // ファイル書き込み
	string_Kconv_fread()       // ファイル読み出し
	string_Kconv_to_buffer()   // バッファに書き込み
	string_Kconv_from_buffer() // バッファから読み出し
}

// ファイル書き込み
func string_Kconv_fwrite() {
	s := "漢字です" // UTF8
	f, _ := os.Create("EUC.txt")
	r := strings.NewReader(s)
	w := transform.NewWriter(f, EUCJP.NewEncoder()) // Encoder->f
	io.Copy(w, r)                                   // r -> w(->Encoder->f)
	f.Close()
}

// ファイル読み出し
func string_Kconv_fread() {
	f, _ := os.Open("EUC.txt")
	b := new(bytes.Buffer)
	r := transform.NewReader(f, EUCJP.NewDecoder()) // f -> Decoder
	io.Copy(b, r)                                   // (f->Decoder)->b
	fmt.Println(b.String())
	f.Close()
}

// バッファに書き込み
func string_Kconv_to_buffer() {
	s := "漢字です" // UTF8
	b := new(bytes.Buffer)
	r := strings.NewReader(s)
	w := transform.NewWriter(b, EUCJP.NewEncoder()) // Encoder->f
	io.Copy(w, r)                                   // r -> w(->Encoder->f)

	st := b.String()
	for i := 0; i < len(st); i++ {
		fmt.Println(st[i])
	}
	fmt.Println(b.String())

}

// バッファから読み出し
func string_Kconv_from_buffer() {
	str_bytes := []byte{180, 193, 187, 250, 164, 199, 164, 185}
	s := bytes.NewBuffer(str_bytes).String() // "漢字です" in EUC

	sr := strings.NewReader(s)
	b := new(bytes.Buffer)
	r := transform.NewReader(sr, EUCJP.NewDecoder()) // sr -> Decoder
	io.Copy(b, r)                                    // (sr->Decoder)->b
	fmt.Println(b.String())
}

//---------------------------------------------------
// マルチバイト文字の数を数える
//---------------------------------------------------
/*
[こちら](http://qiita.com/reiki4040/items/b82bf5056ee747dcf713)が詳しいです。
len()だとbyteカウント、[]runeに変換するとutf-8カウント。
*/
func string_Count() {
	s := "日本語"
	fmt.Println(len(s))         // => 9
	fmt.Println(len([]rune(s))) // => 3
}

//---------------------------------------------------
// マルチバイト文字列の最後の1文字を削除する
//---------------------------------------------------
func string_ChopRune() {
	s := "日本語"
	sc := []rune(s)
	fmt.Println(string(sc[:(len(sc) - 1)])) // => "日本"
}

//---------------------------------------------------
// 文字列
//---------------------------------------------------
func Tips_string() {
	string_Concat()             // 文字列を結合する
	string_Repeat()             // 繰り返し文字列を生成する
	string_UpperLower()         // 大文字・小文字に揃える
	string_Exec()               // コマンドの実行結果を文字列に
	string_HereDocument()       // 複数行の文字列を作成する
	string_HereDocumentIndent() // ヒアドキュメントの終端文字列をインデントする
	string_ExecMultiLine()      // 複数行のコマンドの実行結果を文字列に設定する
	string_Extract()            // 部分文字列を取り出す
	string_ReplacePart()        // 部分文字列を置き換える
	string_Eval()               // 文字列中の式を評価し値を展開する
	string_Each()               // 文字列を1文字ずつ処理する
	string_Trim()               // 文字列の先頭と末尾の空白文字を削除する
	string_ToI()                // 文字列を整数に変換する (to_i)
	string_ToF()                // 文字列を浮動小数点に変換する (to_f)
	string_ParseOct()           // 8進文字列を整数に変換する
	string_ParseHex()           // 16進文字列を整数に変換する
	string_AtoI()               // ASCII文字をコード値に（コード値をASCII文字に）変換する
	string_Just()               // 文字列を中央寄せ・左詰・右詰する
	string_Succ()               // "次"の文字列を取得する
	string_Crypt()              // 文字列を暗号化する
	string_Replace()            // 文字列中で指定したパターンにマッチする部分を置換する
	string_Find()               // 文字列中に含まれている任意文字列の位置を求める
	string_Chomp()              // 文字列の末端の改行を削除する
	string_Split()              // カンマ区切りの文字列を扱う
	string_FindAll()            // 任意のパターンにマッチするものを全て抜き出す
	string_Kconv()              // 漢字コードを変換する
	string_Count()              // マルチバイト文字の数を数える
	string_ChopRune()           // マルチバイト文字列の最後の1文字を削除する

}
