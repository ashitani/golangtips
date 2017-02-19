/*
マップ(ハッシュ)
*/

package tips_map

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//---------------------------------------------------
//プログラム中でマップを定義する
//---------------------------------------------------
/*
Golangではハッシュ（連想配列)のことをmapと呼びます。
*/
func map_Map() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	fmt.Println(m) // => "map[apple:150 banana:300 lemon:300]"
}

//---------------------------------------------------
//キーに関連付けられた値を取得する
//---------------------------------------------------
/*
マップ[キー]でキーに関連した値を取得できます。

未定義のキーを渡すとゼロ(int型の場合は0)が返ります。
```
v, ok:=マップ[キー]
```
で未定義かどうか確認できます。
*/
func map_Get() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

	fmt.Println(m["apple"])  // => "150"
	fmt.Println(m["banana"]) // => "300"
	fmt.Println(m["lemon"])  // => "300"
	fmt.Println(m["papaia"]) // => "0"

	v, ok := m["apple"]
	fmt.Println(v)  // => "150"
	fmt.Println(ok) // => "true"

	v, ok = m["papaia"]
	fmt.Println(v)  // => "0"
	fmt.Println(ok) // => "false"
}

//---------------------------------------------------
//マップに要素を追加する
//---------------------------------------------------
/*
空のマップを作るのは下記のような書き方があります。
*/
func map_Add() {
	m := map[string]int{}
	//m := make(map[string]int)

	m["apple"] = 150
	m["banana"] = 200
	m["lemon"] = 300

	fmt.Println(m["apple"])
}

//---------------------------------------------------
//マップ内にキーが存在するかどうか調べる
//---------------------------------------------------
/*
_, ok = マップ[キー]でokの値を見るのが良さそうです。
*/

func map_HasKey() {

	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

	_, ok := m["apple"]
	fmt.Println(ok) // => "true"
	_, ok = m["orange"]
	fmt.Println(ok) // => "false"

}

//---------------------------------------------------
//マップの要素数を取得する
//---------------------------------------------------
/*
len(マップ)で求められます。
*/
func map_Length() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

	fmt.Println(len(m))
}

//---------------------------------------------------
//キーが存在しない場合のデフォルト値を設定する
//---------------------------------------------------
/*
デフォルトはあくまで0のようですので、新しい型を作るのがよさそうです。
*/
func map_Default() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	dm := dmap{m}

	fmt.Println(dm.Get("apple"))  // => "300"
	fmt.Println(dm.Get("papaia")) // => "100"
}

type dmap struct {
	m map[string]int
}

func (d dmap) Get(key string) int {
	v, ok := d.m[key]
	if ok {
		return v
	} else {
		return 100 // default
	}
}

func (d dmap) Set(key string, value int) {
	d.m[key] = value
}

//---------------------------------------------------
//マップからエントリを削除する
//---------------------------------------------------
/*
delete(マップ,キー)でエントリを削除できます。
rubyのように削除エントリを戻り値に返したりブロックを返したりは
なさそうですので、匿名関数を作って渡す関数を書きます。
*/
func map_Delete() {
	// 消去
	m := map[string]int{"apple": 150, "banana": 300}
	delete(m, "banana")

	// 存在しないときはエラーを表示
	f := func(k string) {
		fmt.Printf("%s not found\n", k)
	}
	delete_if_exist(m, "banana", f) // => "banana not found"
	fmt.Println(m)                  // => "map[apple:150]"

	// 200より小さい値を持つエントリを削除
	m = map[string]int{"apple": 150, "banana": 300, "lemon": 400}
	f_small := func(m map[string]int, k string) bool {
		return (m[k] < 200)
	}
	delete_if(m, f_small)
	fmt.Println(m) // => "map[banana:300 lemon:400]"

}

// 指定したキーが存在しなければf()を実行
func delete_if_exist(m map[string]int, k string, f func(string)) int {
	v, ok := m[k]

	if ok {
		delete(m, k)
		return v
	} else {
		f(k)
		return 0
	}
}

// func_if()がtrueの場合は対象を消す
func delete_if(
	m map[string]int,
	func_if func(map[string]int, string) bool,
) {
	for k := range m {
		ok := func_if(m, k)
		if ok {
			delete(m, k)
		} else {
		}
	}
}

//---------------------------------------------------
//マップの全エントリに対してブロックを実行する
//---------------------------------------------------
/*
```
for key, value := range m {
}
```
でkey,valueを取り出しつつエントリを走査できます。
*/

func map_Block() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

	sum := 0
	fruits := []string{}
	for k, v := range m {
		fruits = append(fruits, k)
		sum += v
	}
	fmt.Println(fruits) // => "[apple banana lemon]"
	fmt.Println(sum)    // => "750"
}

//---------------------------------------------------
//マップを配列に変換する
//---------------------------------------------------
/*
rubyのkeys(), values(), to_a(), indexes() はすべて
range()で実装できます。
*/
func map_ToArray() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	fmt.Println(keys(m))   // => "[apple banana lemon]"
	fmt.Println(values(m)) // => "[150 300 300]"
	fmt.Println(to_a(m))   // => "[[lemon 300] [apple 150] [banana 300]]"

	keys := []string{"apple", "lemon"}
	fmt.Println(indexes(m, keys)) // => "[150 300]"

}

func keys(m map[string]int) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}

func values(m map[string]int) []int {
	vs := []int{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func to_a(m map[string]int) []interface{} {
	a := []interface{}{}
	for k, v := range m {
		a = append(a, []interface{}{k, v})
	}
	return a
}

func indexes(m map[string]int, keys []string) []int {
	vs := []int{}
	for _, k := range keys {
		vs = append(vs, m[k])
	}
	return vs
}

//---------------------------------------------------
//マップを空にする
//---------------------------------------------------
/*
再初期化するのが早そうです。
m=nilでも一見クリアできるのですが、再代入できなくなります。
*/
func map_Clear() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	m = make(map[string]int)
	fmt.Println(m)
}

//---------------------------------------------------
//マップを値で降順、値が等しい場合キーで昇順にソートする
//---------------------------------------------------
/*
[条件式を指定したソート](http://ashitani.jp/golangtips/tips_slice.html#slice_Sort)
と同様のことを行います。
*/
// import "sort"
func map_Sort() {
	m := map[string]int{"ada": 1, "hoge": 4, "basha": 3, "poeni": 3}

	a := List{}
	for k, v := range m {
		e := Entry{k, v}
		a = append(a, e)
	}

	sort.Sort(a)
	fmt.Println(a)
}

type Entry struct {
	name  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name < l[j].name)
	} else {
		return (l[i].value < l[j].value)
	}
}

//---------------------------------------------------
//マップの要素をランダムに抽出する
//---------------------------------------------------
/*
うーん、range()に頼りますか。

乱数は[擬似乱数を生成する](http://ashitani.jp/golangtips/tips_num.html#num_Rand)
にあります。
*/
// import "math/rand"
// import "time"
func map_Random() {
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}

	rand.Seed(time.Now().UnixNano()) //Seed

	fmt.Println(choice(m))
	fmt.Println(choice(m))
	fmt.Println(choice(m))
	fmt.Println(choice(m))

}

func choice(m map[string]int) string {
	l := len(m)
	i := 0

	index := rand.Intn(l)

	ans := ""
	for k, _ := range m {
		if index == i {
			ans = k
			break
		} else {
			i++
		}
	}
	return ans
}

//---------------------------------------------------
//複数のマップをマージする
//---------------------------------------------------
func map_Merge() {
	m1 := map[string]string{"key1": "val1", "key2": "val2"}
	m2 := map[string]string{"key3": "val3"}

	fmt.Println(merge(m1, m2)) // => "map[val1:key1 val2:key2 val3:key3]"
}

func merge(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}

	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

//---------------------------------------------------
// マップ(ハッシュ)
//---------------------------------------------------
func Tips_map() {
	map_Map()     // プログラム中でマップを定義する
	map_Get()     // キーに関連付けられた値を取得する
	map_Add()     // マップに要素を追加する
	map_HasKey()  // マップ内にキーが存在するかどうか調べる
	map_Length()  // マップの要素数を取得する
	map_Default() // キーが存在しない場合のデフォルト値を設定する
	map_Delete()  // マップからエントリを削除する
	map_Block()   // マップの全エントリに対してブロックを実行する
	map_ToArray() // マップを配列に変換する
	map_Clear()   // マップを空にする
	map_Sort()    // マップを値で降順、値が等しい場合キーで昇順にソートする
	map_Random()  // マップの要素をランダムに抽出する
	map_Merge()   // 複数のマップをマージする
}
