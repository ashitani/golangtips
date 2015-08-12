/*
goroutine
*/

package tips_dir

import (
	"fmt"
	"github.com/tlorens/go-ibgetkey"
	"io/ioutil"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"sync"
	"time"
)

//---------------------------------------------------
// goroutineを生成する
//---------------------------------------------------
/*
goroutineはスレッドのようなものです。

joinはないので、終了を待つ場合はchannelを使います。
```
<-quit
```
とすると、quitというchannelに値が書き込まれるまで待ちます。

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
//import "time"

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
終了待ちはchannelの受信で行います。[goroutineを生成する](#goroutine_Create)
を参照してください。
*/

//---------------------------------------------------
// goroutineの実行を終了させる
//---------------------------------------------------
/*
スレッドと違ってgoroutineは外部から終了させることは出来ません。
チャネルをつかって無限ループを抜けるようにするのが一案です。
無限ループを抜ける際に、breakにラベルを与えてどの階層まで抜けるかを指示できるようです。

下記のプログラムはGorutine内で"."を100msec置きに50回表示して終了しますが、
"."を押すとGorutineを途中で終了します。終了の意思を伝えるためにkillというチャネルで
通信を行っています。

なお、一文字入力を受け付けるのに
[go-ibgetkey](https://github.com/tlorens/go-ibgetkey)という
ライブラリを使いました。
*/
// import "github.com/tlorens/go-ibgetkey"
// import "time"

func goroutine_Kill() {
	kill := make(chan bool)
	finished := make(chan bool)
	go killableGoroutine(kill, finished)
	targetkey := "."
	t := int(targetkey[0])
loop:
	for {
		input := keyboard.ReadKey()
		select {
		case <-finished:
			break loop
		default:
			if input == t {
				kill <- true
				break loop
			}
		}
	}
}

func killableGoroutine(kill, finished chan bool) {
	fmt.Println("Started goroutine. Push \".\" to kill me.")
	for i := 0; i < 50; i++ {
		select {
		case <-kill:
			fmt.Println()
			fmt.Println("Killed")
			finished <- true
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println()
	fmt.Println("Finished..push any key to abort.")
	finished <- true
	return
}

//---------------------------------------------------
// goroutineを停止する
//---------------------------------------------------
/*
[goroutineの実行を終了させる](#goroutine_)と同様にchannelで通信を行います。

下記のプログラムはGorutine内で"."を100msec置きに50回表示して終了しますが、
"."を入力するとGorutineを途中で停止・再開します。
*/
// import "github.com/tlorens/go-ibgetkey"
// import "time"

func goroutine_Stop() {
	com := make(chan string)
	finished := make(chan bool)
	go stoppableGoroutine(com, finished)
	targetkey := "."
	t := int(targetkey[0])
	running := true
loop:
	for {
		input := keyboard.ReadKey()
		select {
		case <-finished:
			break loop
		default:
			if input == t {
				if running == true {
					com <- "stop"
					running = false
				} else {
					com <- "start"
					running = true
				}
			}
		}
	}
}

func stoppableGoroutine(command chan string, finished chan bool) {
	fmt.Println("Started goroutine. Push \".\" to stop/start me.")
	running := true
	i := 0
	for i < 50 {
		select {
		case com := <-command:
			if com == "stop" {
				running = false
			} else {
				running = true
			}
		default:
		}
		if running == true {
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
	fmt.Println()
	fmt.Println("Finished..push any key to abort.")
	finished <- true
	return
}

//---------------------------------------------------
// 実行中のgoroutine一覧を取得する
//---------------------------------------------------
/*
Ruby版ではリストを出してそれぞれをkillするというプログラムでしたが、
そもそもkillを簡単に実装する方法がないので、リスト表示だけにとどめます。
goroutineをリスト化する関数自体やmain関数自身もgoroutineなので
多数のgoroutineが表示されます。
*/
//import "os"
//import "runtime/pprof"
//import "time"

func goroutine_ListGoroutines() {
	go goroutine1()
	go goroutine2()

	time.Sleep(1 * time.Second)                     // goroutineの起動のオーバヘッド待ち
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 2) // 2はデバッグレベル。goroutineだけリストする、の意味。
}

func goroutine1() {
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine1 finished")
}

func goroutine2() {
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine2 finished")
}

//---------------------------------------------------
// goroutine間で通信する
//---------------------------------------------------
/*
Ruby版ではQueueを使っているところを、channelに変更します。
makeでchannelの深さを指定しますが、深さ以上に値を突っ込もうとすると
その時点でロックします。[こちら](http://rosylilly.hatenablog.com/entry/2013/09/26/124801)をご参考。

下記プログラムはキー入力を受付け、平方根を返します。-1を入力すると終了します。
*/
// import "math"

func goroutine_Com() {
	queue := make(chan int, 3) // 3はキューの深さ
	go sqrtGoroutine(queue)

	line := 0
loop:
	for {
		fmt.Scanln(&line)
		if line == -1 {
			break loop
		} else {
			queue <- line
		}
	}
}

func sqrtGoroutine(queue chan int) {
	for {
		n := <-queue
		if int(n) >= 0 {
			val := math.Sqrt(float64(n))
			fmt.Printf("Square(%d) = %f\n", int(n), val)
		} else {
			fmt.Println("?")
		}
	}
}

//---------------------------------------------------
// goroutine間の競合を回避する(Mutex)
//---------------------------------------------------
/*
syncパッケージにMutexがあります。複数のgoroutine終了を待つにはWaitGroupを使います。
[こちら](http://mattn.kaoriya.net/software/lang/go/20140625223125.htm)がよくまとまっています。

ファイルcount.txtを作って0を書き込み、gotoutine内ではファイルの数値を読んでインクリメントして書き戻します。
終了後は、goroutineが10個走るのでcount.txtの値は10になっています。
*/
//import "sync"
//import "io/ioutil"

func goroutine_Mutex() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	write(0)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go countupGoroutine(wg, m)
	}
	wg.Wait()
}

func countupGoroutine(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	counter := read() + 1
	write(counter)
	wg.Done()
}

func write(i int) {
	s := strconv.Itoa(i)
	ioutil.WriteFile("count.txt", []byte(s), os.ModePerm)
}

func read() int {
	t, _ := ioutil.ReadFile("count.txt")
	i, _ := strconv.Atoi(string(t))
	return i
}

//---------------------------------------------------
// goroutine
//---------------------------------------------------
func Tips_goroutine() {
	goroutine_Create()   // goroutineを生成する
	goroutine_Argument() // goroutineに引数を渡す
	// goroutineの終了を待つ
	goroutine_Kill()           // goroutineの実行を終了させる
	goroutine_Stop()           // goroutineを停止する
	goroutine_ListGoroutines() // 実行中のgoroutine一覧を取得する
	goroutine_Com()            // goroutine間で通信する
	goroutine_Mutex()          // goroutine間の競合を回避する(Mutex)

}
