/*
goroutine
*/

// package tips_dir
package main

import (
	"fmt"
	//	"os"
	"time"
)

//---------------------------------------------------
// goroutineを生成する
//---------------------------------------------------
/*
goroutineはスレッドのようなものです。

joinはないので、終了を待つ場合はチャンネルを使います。
```
<-quit
```
では、quitというチャネルに値が書き込まれるまで待ちます。

私の環境では、
```
Create goroutine
Waiting for the goroutine to complete
Start goroutine
End goroutine
Test compleated
```
という順番で表示されました。WaitingとStartの順番が直感と合いませんが、
goroutineの起動にオーバヘッドがあるのでしょうかね。
*/
// import "time"

func goroutine_Create() {
	fmt.Println("Create goroutine")
	quit := make(chan bool)
	go func() {
		fmt.Println("Start goroutine")
		time.Sleep(3 * time.Second)
		fmt.Println("End goroutine")
		quit <- true
	}()
	fmt.Println("Waiting for the goroutine to complete")
	<-quit
	fmt.Println("Test compleated")
}

//---------------------------------------------------
// goroutineに引数を渡す
//---------------------------------------------------
/*
goroutineは外部宣言したfuncでも起動できます。
channelが必要な場合は引数にchannelを渡します。
もちろん、channel以外の引数も渡すことができます。
```
Start goroutine - Apple 10
```
と表示されます。
*/

func goroutine_Argument() {
	fmt.Println("Test start")
	fmt.Println("Create goroutine")

	quit := make(chan bool)
	go appleGoroutine("Apple", 10, quit)
	fmt.Println("Waiting for the goroutine to complete")
	<-quit
	fmt.Println("Test compleated")

}

func appleGoroutine(fruit string, a int, quit chan bool) {
	fmt.Printf("Start goroutine - %s %d\n", fruit, a)
	time.Sleep(3 * time.Second)
	fmt.Println("End goroutine")
	quit <- true
}

//---------------------------------------------------
// goroutineの終了を待つ
//---------------------------------------------------
/*
終了はchannelの受信で行います。[goroutineを生成する](#goroutine_Create)
を参照してください。
*/

//---------------------------------------------------
// goroutineの実行を終了させる
//---------------------------------------------------
/*
スレッドと違って外部から終了させることは出来ません。
チャネルをつかって無限ループを抜けるようにするのが一案です。

"."を入力してenterを押すとスレッドを終了します。
*/
// import "os"
func goroutine_Kill() {
	kill := make(chan bool)
	go killableGoroutine(kill)
	var input string
	for {
		fmt.Scanln(&input)
		if input == "." {
			kill <- true
			break
		}
	}
}
func killableGoroutine(kill chan bool) {
	fmt.Println("Start goroutine")
	for i := 0; i < 50; i++ {
		select {
		case <-kill:
			fmt.Println("Killed")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println("End goroutine")
}

//---------------------------------------------------
// goroutineを停止する
//---------------------------------------------------

//---------------------------------------------------
// 実行中のgoroutine一覧を取得する
//---------------------------------------------------
//---------------------------------------------------
// goroutine間で通信する
//---------------------------------------------------
//---------------------------------------------------
// goroutine間の競合を回避する(Mutex)
//---------------------------------------------------

//---------------------------------------------------
// goroutine
//---------------------------------------------------
//func Tips_goroutine() {
func main() {

	//goroutine_Create() // goroutineを生成する
	//goroutine_Argument() // goroutineに引数を渡す
	// // goroutineの終了を待つ
	goroutine_Kill() // goroutineの実行を終了させる
	// goroutineを停止する
	// 実行中のgoroutine一覧を取得する
	// goroutine間で通信する
	// goroutine間の競合を回避する(Mutex)

}
