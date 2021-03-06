/*
ディレクトリ
*/

//package main
package tips_dir

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

//---------------------------------------------------
// ディレクトリの作成
//---------------------------------------------------
/*
os.Mkdir()が使えます。第二引数はモードです。
*/
// import "os"

func dir_MakeDir() {
	err := os.Mkdir("tmp", 0777)
	fmt.Println(err)
}

//---------------------------------------------------
// ディレクトリの削除
//---------------------------------------------------
/*
ファイルでもディレクトリでも
os.Remove()が使えます。ただしディレクトリの場合、中身が空でない場合はエラーになります。
*/
// import "os"

func dir_RemoveDir() {
	err := os.Remove("./tmp")
	fmt.Println(err)
}

//---------------------------------------------------
// 中身が空でないディレクトリを削除する
//---------------------------------------------------
/*
os.RemoveAll()が使えます。
*/
// import "os"

func dir_RemoveDirAll() {
	err := os.RemoveAll("./tmp")
	fmt.Println(err)
}

//---------------------------------------------------
// ディレクトリ名を変更する
//---------------------------------------------------
/*
ファイルでもディレクトリでもos.Rename()が使えます。
*/
// import "os"

func dir_Rename() {
	err := os.Rename("sample", "doc")
	fmt.Println(err)
}

//---------------------------------------------------
// ディレクトリの詳細情報を取得する
//---------------------------------------------------
/*
ファイルでもディレクトリでも同じです。
[ファイルの詳細情報を取得する](http://ashitani.jp/golangtips/tips_file.html#file_Stat)
を参照のこと。
*/

//---------------------------------------------------
// ディレクトリのファイルモードを変更する
//---------------------------------------------------
/*
ファイルでもディレクトリでも同じです。
[ファイルモードを変更する](http://ashitani.jp/golangtips/tips_file.html#file_ChMod)
を参照のこと。
*/

//---------------------------------------------------
// ディレクトリの所有者とグループを変更する
//---------------------------------------------------
/*
ファイルでもディレクトリでも同じです。
[ファイルの所有者とグループを変更する](http://ashitani.jp/golangtips/tips_file.html#file_ChOwn)
を参照のこと。
*/

//---------------------------------------------------
// ディレクトリの最終アクセス時刻と最終更新日時を変更する
//---------------------------------------------------
/*
ファイルでもディレクトリでも同じです。
[ファイルの最終アクセス時刻と最終更新日時を変更する](http://ashitani.jp/golangtips/tips_file.html#file_ChangeTime)
を参照のこと。
*/

//---------------------------------------------------
// カレントディレクトリの取得と変更
//---------------------------------------------------
/*
os.Getwd()とos.Chdir()が使えます。
*/
// import "os"

func dir_Pwd() {
	// 取得
	p, _ := os.Getwd()
	fmt.Println(p)

	// 変更
	os.Chdir("/etc")

	// 確認
	p, _ = os.Getwd()
	fmt.Println(p)
}

//---------------------------------------------------
// ディレクトリ中のファイル一覧を取得する
//---------------------------------------------------
/*
io/ioutilのReadDir()が使えます。
*/
// import "io/ioutil"

func dir_GetFileList() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

//---------------------------------------------------
// ワイルドカードにマッチしたファイル全てに処理を行う
//---------------------------------------------------
/*
path/filepathにGlob()があります。下記の例では全ファイルの名称とファイルサイズを
表示します。ただしフォルダを再帰的に展開したりはしません。

指定したフォルダを再帰的に展開したい場合には、Walk()を使います。WalkFunc()と同じ
インターフェイスを持つ関数を指定すると、訪問先でその関数が実行されます。
*/
// import "path/filepath"
// import "syscall"

func dir_Glob() {

	// 再帰なし
	files, _ := filepath.Glob("/etc/*")
	for _, f := range files {
		printPathAndSize(f)
	}

	fmt.Println("---------")

	// 再帰あり
	filepath.Walk("/etc/", visit)

}

func visit(path string, info os.FileInfo, err error) error {

	printPathAndSize(path)
	return nil
}

func printPathAndSize(path string) {
	// ファイルサイズの取得
	var s syscall.Stat_t
	syscall.Stat(path, &s)

	fmt.Print(path)
	fmt.Print(": ")
	fmt.Print(s.Size)
	fmt.Println(" bytes")

}

//---------------------------------------------------
// ファイル名からディレクトリ部分だけを切り出す
//---------------------------------------------------
/*
os/filepathのDir()を使います。
*/
// import "os/filepath"

func dir_DirName() {
	fmt.Println(filepath.Dir("/usr/bin/ruby")) // => "/usr/bin"
	fmt.Println(filepath.Dir("/etc/passwd"))   // => "/etc"
}

//---------------------------------------------------
// ディレクトリかどうか判定する
//---------------------------------------------------
/*
os.Stat()でos.FileInfoが返ります。FileInfoのIsDir()で確認できます。

シンボリックリンクかどうかなど、詳細情報が知りたい場合は、
FileInfo.Mode()でFileModeが返るのでそちらを使います。
*/
// import "os"

func dir_IsDir() {

	fInfo, _ := os.Stat("/etc")
	fmt.Println(fInfo.IsDir()) // => "true"
}

//---------------------------------------------------
// ディレクトリ内の全ファイルに対して処理を行う
//---------------------------------------------------
/*
[ワイルドカードにマッチしたファイル全てに処理を行う](#dir_Glob)参照のこと。
*/

//---------------------------------------------------
// ディレクトリ内の全ファイル名をフルパスで表示
//---------------------------------------------------

/*
[ワイルドカードにマッチしたファイル全てに処理を行う](#dir_Glob)と同様に、
Walk()で可能です。
*/
// import "path/filepath"

func dir_ShowFullPath() {
	filepath.Walk("/etc/", showFullPath)
}

func showFullPath(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}

//---------------------------------------------------
// ディレクトリ
//---------------------------------------------------
func Tips_file() {

	dir_MakeDir()      // ディレクトリの作成
	dir_RemoveDir()    // ディレクトリの削除
	dir_RemoveDirAll() // 中身が空でないディレクトリを削除する
	dir_Rename()       // ディレクトリ名を変更する
	// ディレクトリの詳細情報を取得する
	// ディレクトリのファイルモードを変更する
	// ディレクトリの所有者とグループを変更する
	// ディレクトリの最終アクセス時刻と最終更新日時を変更する
	dir_Pwd()         // カレントディレクトリの取得と変更
	dir_GetFileList() // ディレクトリ中のファイル一覧を取得する
	dir_Glob()        // ワイルドカードにマッチしたファイル全てに処理を行う
	dir_DirName()     // ファイル名からディレクトリ部分だけを切り出す
	dir_IsDir()       // ディレクトリかどうか判定する
	// ディレクトリ内の全ファイルに対して処理を行う
	dir_ShowFullPath() // ディレクトリ内の全ファイル名をフルパスで表示

}
