/*
ファイル
*/

//package tips_file
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

//---------------------------------------------------
// ファイルをオープンする
//---------------------------------------------------
/*
読み出しには`os.Open()`書き込みには`os.Create()`を使います。

テキストの読み書きはいろいろ方法があると思いますが、ここでは
読み出しに`bufio.NewScanner()`を、書き込みに`bufio.NewWriter()`を使います。
*/
// import "os"
// import "bufio"

func file_Open() {
	// 書き込み
	fpw, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(fpw)
	fmt.Fprint(w, "Hello, golang!")
	w.Flush()
	fpw.Close()

	// 読み出し
	fp, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fp.Close()
}

//---------------------------------------------------
// テキストファイルをオープンして内容を出力する
//---------------------------------------------------
/*
[ファイルをオープンする](#file_Open)と同じ内容ですが。
*/
// import "os"
// import "bufio"

func file_Read() {
	fp, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fp.Close()
}

//---------------------------------------------------
// 読み込む長さを指定する
//---------------------------------------------------
/*
[]byteバッファを作ってRead()に渡すと、バッファの長さ分だけ
読み込んでくれます。
*/
// import "os"

func file_ReadLength() {
	fp, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	data := make([]byte, 10)
	count, err := fp.Read(data)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes: %s\n", count, data)
}

//---------------------------------------------------
// ファイルの内容を一度に読み込む
//---------------------------------------------------
/*
io/ioutilのioutil.ReadFile()を使います。
*/
// import  "io/ioutil"

func file_ReadAll() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

//---------------------------------------------------
// 1行ずつ読み込みを行う
//---------------------------------------------------
/*
bufioのScannerをつかってみます。この例では、foo.csvから行数と総エントリ数をカウントします。
*/
// import "bufio"
// import "os"
// import "strings"

func file_ReadEachLine() {
	data, err := readAll("foo.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func readAll(filename string) (string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf(filename + " can't be opened")
	}

	ans := ""
	lines := 0
	entries := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		entry := scanner.Text()
		lines += 1
		slice := strings.Split(entry, ",")
		entries += len(slice)
	}

	fp.Close()
	fmt.Printf("Read %d lines, %d entries\n", lines, entries)
	return string(ans), nil
}

//---------------------------------------------------
// テキストファイルの特定の行を読み込む
//---------------------------------------------------
/*
rubyのreadlinesのように、配列に読み込むような標準関数はありません。
改行でSplitすれば代替になりますが、でかいファイルの場合は気をつけないとですね。
*/
// import "ioutil"
// import "strings"

func file_ReadSpecificLine() {
	ans, err := readLines("foo.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(ans[1])
}

func readLines(filename string) ([]string, error) {
	ans := make([]string, 10)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ans, fmt.Errorf(filename + " can't be opened")
	}
	ans = strings.Split(string(data), "\n")

	return ans, err
}

//---------------------------------------------------
// 一時ファイルを作成する
//---------------------------------------------------
/*
ioutil.TempFile(dir,prefix string)を使います。dirは一時ファイルを
作成するフォルダです。""を指定すると
osごとのデフォルトのディレクトリが使われるようです。prefixに指定した文字が
ファイル名に入ります。

下記の例は、test.txtの内容を、TempFileを利用して大文字に変換します。
*/
// import "bufio"
// import "ioutil"
// import "os"

func file_TempFile() {

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	fw, err := ioutil.TempFile(".", "temp")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		entry := scanner.Text()
		fw.Write([]byte(strings.ToUpper(entry)))
	}

	fw.Close()
	f.Close()
	os.Rename(fw.Name(), f.Name())

}

//---------------------------------------------------
// 固定長レコードを読む
//---------------------------------------------------
/*
以下の様な固定長レコードを読みます。
```
■レコード形式
 従業員番号 6桁|氏名 utf5文字|部課コード 4桁|入社年度 4桁

■レコード例
 100001鈴木一郎太12342001

  従業員番号: 100001
  氏名: 鈴木一郎太
  部課コード: 1234
  入社年度: 2001
```

record という構造体を作り、その配列recordsに対して、
レコード追加・表示用のインターフェイス(Insert,List)と、
ソート用のインターフェイス(Len,Less,Swap)を準備。

*/
// import "bufio"
// import "sort"
// import "strconv"
// import "os"

func file_FormattedText() {
	var rs records

	f, err := os.Open("fmtTxt.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		entry := scanner.Text()
		rs.InsertFromString(entry)
	}

	sort.Sort(rs)
	rs.List()
}

type record struct {
	empno  int
	name   string
	deptno int
	year   int
}

type records []record

func (rs *records) InsertFromString(s string) {
	var newRecord record

	r := []rune(s)

	newRecord.empno, _ = strconv.Atoi(string(r[0:6]))
	newRecord.name = string(r[6:11])
	newRecord.deptno, _ = strconv.Atoi(string(r[11:15]))
	newRecord.year, _ = strconv.Atoi(string(r[15:19]))
	*rs = append(*rs, newRecord)
}

func (r records) List() {
	for _, x := range r {
		fmt.Printf("従業員番号:\t%d\n", x.empno)
		fmt.Printf("氏名:\t\t%s\n", x.name)
		fmt.Printf("部課コード:\t%d\n", x.deptno)
		fmt.Printf("入社年度:\t%d\n", x.year)
		fmt.Println(strings.Repeat("-", 20))
	}
}

func (r records) Len() int {
	return len(r)
}

func (r records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r records) Less(i, j int) bool {
	return r[i].year < r[j].year
}

//---------------------------------------------------
// ファイルをコピーする
//---------------------------------------------------
/*
os.Link()が簡単でしょう。
*/
// import "os"

func file_CopyFile() {
	src := "test.txt"
	dest := "test.bak"
	_ = os.Link(src, dest)
}

//---------------------------------------------------
// フィルタ系のコマンドを作成する
//---------------------------------------------------
/*
標準入力を受け取りたいときはos.Stdinを使います。
rubyのように、ファイル名指定されたら勝手に標準入力扱いという器用なことは
できないので、os.Argsでとれるコマンドライン引数リストを走査します。

```
echo hoge | go run ./main.go
go run ./main.go test.txt test.bak
```
のどちらでも機能します。

C言語と違って、os.Argsがどこでも機能する（mainの引数を変更する必要がない）のは
一瞬気持ち悪いですが、スッキリ書けますね。

*/
// import "bufio"
// import "os"

func file_Filter() {

	l := len(os.Args)
	if l == 1 { // 引数なしの場合
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t := scanner.Text()
			if t != "" {
				fmt.Println(t)
			} else {
				break
			}
		}
	} else { // ファイル名渡しの場合
		for _, x := range os.Args[1:l] {
			f, err := os.Open(x)
			if err != nil {
				panic(err)
			}
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}
}

//---------------------------------------------------
// ファイルタイプを取得する
//---------------------------------------------------
/*
[こちら](http://qiita.com/shinofara/items/e5e78e6864a60dc851a6)を参考に。
os.Stat()でos.FileInfoが返ります。

存在有無は、os.Stat()のエラーをos.IsExist()に渡すと確認できます。
ちょっとくどいですね。
*/
// import "os"

func file_FileType() {
	fmt.Println(isDir("/etc/passwd")) // => false
	fmt.Println(isDir("/etc"))        // => true

	fmt.Println(isExist("/etc/passwd"))   // => true
	fmt.Println(isExist("/etc/password")) // => false
}

func isDir(filename string) bool {
	fInfo, _ := os.Stat(filename)
	return fInfo.IsDir()
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	} else {
		return os.IsExist(err)
	}
}

//---------------------------------------------------
// ファイルの詳細情報を取得する
//---------------------------------------------------
/*
むかしはFileStatでi-node番号とかもろもろ取得できたようですが、
少なくともGo1.4ではsyscallにいます。Atimespec->AtimになったのはGo1.5から？
*/
// import "syscall"

func file_Stat() {
	var s syscall.Stat_t
	syscall.Stat("/etc/passwd", &s)
	fmt.Println(s.Dev)
	fmt.Println(s.Ino)
	fmt.Println(s.Mode)
	fmt.Println(s.Nlink)
	fmt.Println(s.Uid)
	fmt.Println(s.Gid)
	fmt.Println(s.Size)
	fmt.Println(s.Blocks)
	// fmt.Println(s.Atim.Unix()) // Go1.4ではエラー
	// fmt.Println(s.Mtim.Unix()) // Go1.4ではエラー
	fmt.Println(s.Atimespec.Unix())
	fmt.Println(s.Mtimespec.Unix())
}

//---------------------------------------------------
// ファイルモードを変更する
//---------------------------------------------------
/*
os.Chmod()を使います。FileMode構造体が返るのですが、String()メソッドが定義されているので、
Printlnすると"-rw-rw-rw-"のような表示が出ます。便利なのか？と思いましたが、
8進数を誤解しやすいからかしら。
*/
// import "os"

func file_ChMod() {
	filename := "test.txt"

	s, _ := os.Stat(filename)
	fmt.Println(s.Mode()) // -> "-rw-------"

	os.Chmod(filename, 0666)

	s, _ = os.Stat(filename)
	fmt.Println(s.Mode()) // -> "-rw-rw-rw-"
}

//---------------------------------------------------
// ファイルの所有者とグループを変更する
//---------------------------------------------------
/*
os.Chown()があります。もちろんchwonできる権限がないとダメですが。

ちなみにOSXでユーザIDを調べるのは、
```
id
dscl . -read /users/$username uid
```
などの方法があります。
*/
//import "os"

func file_ChOwn() {
	err := os.Chown("test.txt", 502, 20) //uid=502,gid=20
	if err != nil {
		fmt.Println("chown not permitted")
	}
}

//---------------------------------------------------
// ファイルの最終アクセス時刻と最終更新日時を変更する
//---------------------------------------------------
/*
下記の例は、test.txtの
最終アクセス日を2001-5-22 23:59:59(JST)、
最終更新日を2001-5-1 00:00:00(JST)に変更するものです。
*/
// import "syscall"
// import "time"

func file_ChangeTime() {

	//　変更
	atime := time.Date(2001, 5, 22, 23, 59, 59, 0, time.Local)
	mtime := time.Date(2001, 5, 1, 00, 00, 00, 0, time.Local)
	os.Chtimes("test.txt", atime, mtime)

	//　確認
	var s syscall.Stat_t
	syscall.Stat("test.txt", &s)
	sec, nsec := s.Atimespec.Unix()   // Go1.5以降ではAtimespec -> Atim
	fmt.Println(time.Unix(sec, nsec)) // => "2001-05-22 23:59:59 +0900 JST"
	sec, nsec = s.Mtimespec.Unix()    // Go1.5以降ではAtimespec -> Atim
	fmt.Println(time.Unix(sec, nsec)) // => "2001-05-01 00:00:00 +0900 JST"

}

//---------------------------------------------------
// 相対パスから絶対パスを求める
//---------------------------------------------------
/*
カレントを基準に相対パスを絶対パスに変換するにはpath/filepathのAbs()を使います。

指定したベースパスを基準に相対パスを絶対パスに変換するにはJoin()を使います。

絶対パスを、ベースパスを基準に相対パスに変換するにはRel()を使います。
*/
// import "path/filepath"

func file_AbsPath() {
	apath, _ := filepath.Abs("./test.txt")
	fmt.Println(apath)

	apath = filepath.Join("/etc", "passwd")
	fmt.Println(apath) // => "/etc/passwd"

	rpath, _ := filepath.Rel("/etc", "/etc/passwd")
	fmt.Println(rpath) // -> "passwd"
}

//---------------------------------------------------
// ファイルパスからディレクトリパスを抜き出す
//---------------------------------------------------
/*
path/filepath のDir()を使います。
Windowsのパスはうまく展開されませんけど、OSXで実行したからかしら。
*/
// import "path/filepath"

func file_Dir() {
	d := filepath.Dir("/hoge/piyo")
	fmt.Println(d) // =>"/hoge"

	d = filepath.Dir("/hoge/piyo/")
	fmt.Println(d) // =>"/hoge/piyo"
	d = filepath.Dir("c:\\hoge\\piyo")
	fmt.Println(d) // =>"."  ？？
}

//---------------------------------------------------
// ファイルパスからファイル名を抜き出す
//---------------------------------------------------
/*
path/filepath のBase()でファイル名を分離できます。
拡張子を取り出すのはExt()ですが、拡張子を取り除いたbasenameを
取る方法は簡単にはなさそうですので、正規表現で拡張子を取り除いてからBase()を実行しました。
*/
// import "path/filepath"
// import "regexp"

func file_Basename() {
	b := filepath.Base("/hoge/piyo")
	fmt.Println(b) // => "piyo"

	e := filepath.Ext("/hoge/piyo.c")
	fmt.Println(e) // => ".c"

	rep := regexp.MustCompile(`.c$`)
	e = filepath.Base(rep.ReplaceAllString("/hoge/piyo.c", ""))
	fmt.Println(e) // => "piyo"
}

//---------------------------------------------------
// パス名とファイル名を一度に抜き出す
//---------------------------------------------------
/*
path/filepath のSplit()でディレクトリ名とファイル名を分離できます。
*/
// import "path/filepath"

func file_Split() {
	d, f := filepath.Split("/hoge/piyo")
	fmt.Println(d) // => "/hoge/"
	fmt.Println(f) // => "piyo"
}

//---------------------------------------------------
// 拡張子を調べる
//---------------------------------------------------
/*
path/filepath のExt()で拡張子を取り出せます。
*/
// import "path/filepath"

func file_Ext() {
	e := filepath.Ext("/hoge/piyo.c")
	fmt.Println(e) // => ".c"
}

//---------------------------------------------------
// ファイル
//---------------------------------------------------
func Tips_file() {
	file_Open()             // ファイルをオープンする
	file_Read()             // テキストファイルをオープンして内容を出力する
	file_ReadLength()       // 読み込む長さを指定する
	file_ReadAll()          // ファイルの内容を一度に読み込む
	file_ReadEachLine()     // 1行ずつ読み込みを行う
	file_ReadSpecificLine() // テキストファイルの特定の行を読み込む
	file_TempFile()         // 一時ファイルを作成する
	file_FormattedText()    // 固定長レコードを読む
	file_CopyFile()         // ファイルをコピーする
	file_Filter()           // フィルタ系のコマンドを作成する
	file_FileType()         // ファイルタイプを取得する
	file_Stat()             // ファイルの詳細情報を取得する
	file_ChMod()            // ファイルモードを変更する
	file_ChOwn()            // ファイルの所有者とグループを変更する
	file_ChangeTime()       // ファイルの最終アクセス時刻と最終更新日時を変更する
	file_AbsPath()          // 相対パスから絶対パスを求める
	file_Dir()              // ファイルパスからディレクトリパスを抜き出す
	file_Basename()         // ファイルパスからファイル名を抜き出す
	file_Split()            // パス名とファイル名を一度に抜き出す
	file_Ext()              // 拡張子を調べる
}
