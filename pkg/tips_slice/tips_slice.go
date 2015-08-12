/*
配列
*/

package tips_slice

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	matrix "github.com/skelterjohn/go.matrix"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//---------------------------------------------------
//プログラムで配列を定義する
//---------------------------------------------------
/*
配列は固定長、スライスは可変長の配列のようなもの、です。
個人的には、特に理由が無い限りスライスを使えばいいんじゃないかと思っています。
初期化は型をTとすると
```
a:=[]T{}
a:=make([]T,len)
```
などで行います。

スライスの取り扱いはGolang Wikiの[SliceTricks](https://github.com/golang/go/wiki/SliceTricks)に詳しいです。

スライスの中身にいろんな型を混在させたい場合はinterface{}型で初期化します。
ただし要素を取り出す際にキャストが必要です。interface型のキャストは、キャスト先の型をTとすると、
```
v,ok:=x.(T)
```と書きます。fmt.Println()などはStringへのキャストを自前で行うので注意。

スライスのネストも可能ですがこれまた同じようにキャストして取り出す必要がありなかなか面倒です。

*/
func slice_Define() {
	fruits := []string{"apple", "orange", "lemon"}
	fmt.Println(fruits) // => "[apple orange lemon]"

	scores := []int{55, 49, 100, 150, 0}
	fmt.Println(scores) // => "[55 49 100 150 0]"

	tmp := []interface{}{"apple", 10, 2.5}
	fmt.Println(tmp) // => "[apple 10 2.5]"

	fruites2 := []interface{}{
		3,
		[]interface{}{"apple", 250},
		[]interface{}{"orange", 400},
		[]interface{}{"lemon", 300},
	}
	fmt.Println(fruites2[0].(int))
	f, _ := fruites2[1].([]interface{})
	fmt.Println(f[1]) // => "250"

	f, _ = fruites2[3].([]interface{})
	fmt.Println(f[0]) // => "lemon"

}

//---------------------------------------------------
//m x n 行列の形で配列の配列を初期化する
//---------------------------------------------------
/*
スライスのスライスを作ってせっせと要素を詰めてもよいですが、
行列演算に特化したMatrixパッケージをつかってもよいでしょう。ここでは
[go.matrix](https://godoc.org/github.com/skelterjohn/go.matrix)
をつかってみます。Zeros, Ones, ParseMatlabなど、Matlab経験者には嬉しい
関数がそろってます。
*/
//import matrix "github.com/skelterjohn/go.matrix"

func slice_SliceOfSlice() {

	// Slice of Slice
	dary := make([][]int, 4)
	for i := range dary {
		dary[i] = make([]int, 3)
	}
	dary[0][1] = 7
	fmt.Println(dary) // => "[[0 7 0] [0 0 0] [0 0 0] [0 0 0]]"

	// Matrix
	dmat := matrix.Zeros(4, 3)
	dmat.Set(0, 1, 7)
	fmt.Println(dmat)
	// =>{0, 7, 0,
	//    0, 0, 0,
	//    0, 0, 0,
	//    0, 0, 0}
}

//---------------------------------------------------
//配列要素をカンマ区切りで出力する
//---------------------------------------------------
// import "strings"

func slice_Join() {
	// stringならJoin一発
	fruits := []string{"apple", "orange", "lemon"}
	fmt.Println(strings.Join(fruits, ",")) // => "apple,orange,lemon"
	fmt.Println(strings.Join(fruits, "#")) // => "apple#orange#lemon"

	//Joinはstringしかできない。。map,reduceみたいなのもないし、しょうがないか。
	numbers := []int{55, 49, 100, 100, 0}
	str := ""
	for _, v := range numbers {
		str += fmt.Sprintf("%d,", v)
	}
	str = strings.TrimRight(str, ",") //右端の","を取り除く
	fmt.Println(str)                  // => "55,49,100,100,0"

	//チートなやり方としては、Sprintfの整形を利用する。。
	str = fmt.Sprintf("%v", numbers)
	str = strings.Trim(str, "[]")
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println(str) // => "55,49,100,100,0"

	//うーんこれもかなり辛い。。
	fruits2 := []interface{}{
		3,
		[]interface{}{"apple", 250},
		[]interface{}{"orange", 400},
	}
	str = ""
	for _, v := range fruits2 {
		w, ok := v.([]interface{})
		if ok == false {
			str += fmt.Sprintf("%v,", v)
		} else {
			for _, x := range w {
				str += fmt.Sprintf("%v,", x)
			}
		}
	}
	str = strings.TrimRight(str, ",") //右端の","を取り除く
	fmt.Println(str)                  // => "3,apple,250,orange,400"

}

//---------------------------------------------------
//配列の要素数を取得する
//---------------------------------------------------
func slice_Count() {
	fruits := []string{"apple", "orange", "lemon"}
	fmt.Println(len(fruits)) // => "3"

	num := []int{55, 49, 100, 100, 0}
	fmt.Println(len(num)) // => "5"

	fruits2 := []interface{}{
		3,
		[]interface{}{"apple", 250},
		[]interface{}{"orange", 400},
	}
	fmt.Println(len(fruits2)) // => "3"
}

//---------------------------------------------------
//配列に要素を追加する
//---------------------------------------------------
/*
スライスが対象の場合はappendで追加できます。
スライスには容量がありますが、容量を超える場合はメモリ確保も
同時に行ってくれます。

スライス同士を結合する場合は、結合される側のスライスに"..."を付記します。
*/
func slice_Append() {
	num := []int{1, 2, 3, 4, 5}
	num = append(num, 99)
	fmt.Println(num) // => "[1 2 3 4 5 99]"

	//unshiftをしたい場合は、うーん、スライス同士の結合で。
	num = append([]int{99}, num...)
	fmt.Println(num) // => "[99 1 2 3 4 5 99]"
}

//---------------------------------------------------
//配列の先頭または末尾から要素を取りだす
//---------------------------------------------------
/*
スライスの要素を取り除く関数はないので、
新しいスライスを作るしかなさそうです。
途中から取り除きたい場合は、
[こちら](http://stackoverflow.com/questions/25025409/delete-element-in-slice-golang)
によると、
```
a = append(a[:i], a[i+1:]...)
```
のようにすると良いらしいです。

また、スライス引数は参照渡しなので、本来は関数内で
変更した内容が呼び出し元にも影響を与えますが、
append等のcapを変更する可能性のある操作については、
呼び元には影響を与えず、
変更を加えたければ戻り値で返すかポインタ渡しをすることになるようです。
[こちら](http://jxck.hatenablog.com/entry/golang-slice-internals2)
に詳しいです。ここでは戻り値で返してます。

あと、Pythonでよくある、スライスのa[:-1]のような負indexは
非サポートだそうです。残念。
*/
func slice_Pop() {
	num := []int{1, 2, 3, 4, 5}
	num = append(num, 10)
	v, num := pop(num)
	fmt.Println(v) // =>"10"
	v, num = pop(num)
	fmt.Println(v)   // =>"5"
	fmt.Println(num) //="[1 2 3 4]"
}

func pop(slice []int) (int, []int) {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

//---------------------------------------------------
//部分配列を取りだす
//---------------------------------------------------
/*
a[0:2]のように取り出すことが出来ます。
rubyの[1..3]みたいな、個数指定はありません。

破壊的メソッドはありません。pop()と同様に関数の戻り値で対応します。
スライスは配列へのポインタのようなものなので、コピーの類は、makeで
ポインタを作成してcopy()するのがよいです。
*/
func slice_Slice() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[0:2]) // => "[1 2]"
	fmt.Println(a[1:4]) // => "[2 3 4]"

	// 破壊的メソッド
	num, a, _ := slice(a, 0, 2)
	fmt.Println(num) // =>"[1 2]"
	fmt.Println(a)   // =>"[3 4 5]"
	num, a, _ = slice(a, 1, 3)
	fmt.Println(num) // => "[4 5]"
	fmt.Println(a)   // => "[3]"

}

func slice(slice []int, start, end int) ([]int, []int, error) {
	if len(slice) < start || len(slice) < end {
		return nil, nil, fmt.Errorf("Error")
	}
	ans := make([]int, (end - start))
	copy(ans, slice[start:end])
	slice = append(slice[:start], slice[end:]...)
	return ans, slice, nil
}

//---------------------------------------------------
//配列を任意の値で埋める
//---------------------------------------------------
/* fillのような関数はないので、自前でやるしかありません。 */
func slice_Fill() {
	a := []int{1, 2, 3, 4, 5}

	a, _ = fill(a, 255, 2, 4)
	fmt.Println(a) // =>"[1 2 255 255 5]"
	a, _ = fill(a, 0, 1, 3)
	fmt.Println(a) // =>"[1 0 0 255 5]"
}

func fill(slice []int, val, start, end int) ([]int, error) {
	if len(slice) < start || len(slice) < end {
		return nil, fmt.Errorf("Error")
	}
	for i := start; i < end; i++ {
		slice[i] = val
	}
	return slice, nil
}

//---------------------------------------------------
//配列を空にする
//---------------------------------------------------
/*空にしたいときはnilを代入します。*/
func slice_Clear() {
	a := []int{1, 2, 3, 4, 5}
	a = nil
	fmt.Println(a) // => "[]"
}

//---------------------------------------------------
//配列同士を結合する
//---------------------------------------------------
/*
appendでスライス同士を結合するには、
```
append(a,b...)
```
のように記載します。
*/
func slice_Concat() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, []int{10, 20}...)
	fmt.Println(a) // => [1 2 3 4 5 10 20]
}

//---------------------------------------------------
//配列同士の和・積を取る
//---------------------------------------------------
/*
集合の概念はなさそうですので、[golang-set](https://github.com/deckarep/golang-set)
を使います。[]interface{}を受け取るようです。集合の内容は順不同みたいですね。

和集合はUnion()、積集合はIntersect()で演算できます。
*/
//import set "github.com/deckarep/golang-set"

func slice_Union() {
	a := set.NewSetFromSlice([]interface{}{1, 3, 5, 7})
	b := set.NewSetFromSlice([]interface{}{2, 4, 6, 8})
	fmt.Println(a.Union(b)) // => "Set{8, 5, 7, 1, 3, 2, 4, 6}"

	a = set.NewSetFromSlice([]interface{}{1, 2, 3, 4})
	b = set.NewSetFromSlice([]interface{}{3, 4, 5, 6})
	fmt.Println(a.Union(b)) // => "Set{3, 5, 6, 4, 1, 2}"

	a = set.NewSetFromSlice([]interface{}{1, 3, 5, 7})
	b = set.NewSetFromSlice([]interface{}{2, 4, 6, 8})
	fmt.Println(a.Intersect(b)) // => "Set{}"

	a = set.NewSetFromSlice([]interface{}{1, 2, 3, 4})
	b = set.NewSetFromSlice([]interface{}{3, 4, 5, 6})
	fmt.Println(a.Intersect(b)) // => "Set{4,3}"

}

//---------------------------------------------------
//複数の要素を変更する
//---------------------------------------------------
/*
```
a[0:3] = []int{111, 222, 333}
```
のような書き方はできません。
*/
func slice_Replace() {
	a := []int{1, 2, 3, 4, 5}
	a, _ = replace(a, []int{111, 222, 333}, 0, 2)
	fmt.Println(a) //=>"[111 222 333 3 4 5]"

	a, _ = replace(a, []int{444, 555}, 3, 5)
	fmt.Println(a) //=>"[111 222 333 444 555 5]"
}

func replace(slice []int, rep []int, start, end int) ([]int, error) {
	if len(slice) < start || len(slice) < end {
		return nil, fmt.Errorf("Error")
	}
	ans := make([]int, len(slice))
	copy(ans, slice)
	ans = append(ans[:start], rep...)
	ans = append(ans, slice[end:]...)
	return ans, nil
}

//---------------------------------------------------
//配列の配列をフラットな配列にする
//---------------------------------------------------
/*
[]interface{}をたくさん書くと見難いので、ここではany型と名づけます。

T型への（成功するか不明な）キャストは、
```
b,ok:= a.(T)
```
で行えます。okの値次第で処理を変えれば型ごとに違う処理を書けます。
下の例では、intとanyのみで構成されるanyに対して、int変換が成功したら
取り出す、という処理を再帰的に行っています。
*/
type any []interface{}

func slice_Flatten() {
	a := any{1, any{2, any{3, 4}, 5}, any{6, 7}}
	fmt.Println(a)
	result := any{}
	result = flatten(a, result)
	fmt.Println(result)
}

func flatten(in, result any) any {
	for _, x := range in {
		s, ok := x.(int)
		if ok {
			result = append(result, s)
		} else {
			result = flatten(x.(any), result)
		}
	}
	return result
}

//---------------------------------------------------
//配列をソートする
//---------------------------------------------------
/*
sortパッケージの使い方は
[こちら](http://qiita.com/Jxck_/items/fb829b818aac5b5f54f7)
が詳しいです。
*/

//import "sort"
func slice_Sort() {
	a := []int{5, 1, 4, 2, 3}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a) // => [1 2 3 4 5]

	s := []string{"Orange", "Apple", "Lemon"}
	sort.Strings(s)
	fmt.Println(s) // => [Apple Lemon Orange]
}

//---------------------------------------------------
//条件式を指定したソート
//---------------------------------------------------
/*
Len(),Less(),Swap()という３つのメソッドを
実装したクラスを作れば、独自条件のソートが可能です。
*/
func slice_CaseSort() {
	a := People{"Hitoshi,045", "Sizuo,046", "Yoshi,0138"}
	sort.Sort(a)
	fmt.Println(a) // => "[Yoshi,0138 Hitoshi,045 Sizuo,046]"
}

type Person string

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	xi := strings.Split(string(p[i]), ",")
	xj := strings.Split(string(p[j]), ",")
	return xi[1] < xj[1]
}

//---------------------------------------------------
//配列の配列を任意の要素でソートする
//---------------------------------------------------
/*
ほとんど[こちら](http://qiita.com/Jxck_/items/fb829b818aac5b5f54f7)の丸写しですが。

基底クラスを作り、そのクラスを要素に持つstructを作ると、親クラスのメソッドが引き継がれるのですね。
ここはちょっとわかりにくい。。
*/
func slice_SortAnyColumn() {
	ar := NSslice{
		NS{2, "b"},
		NS{3, "a"},
		NS{1, "c"},
	}

	sort.Sort(ByN{ar})
	fmt.Println(ar) // => "[{1 c} {2 b} {3 a}]"

	sort.Sort(ByS{ar})
	fmt.Println(ar) // => "[{3 a} {2 b} {1 c}]"
}

// 基本クラス
type NS struct {
	Num int
	Str string
}

type NSslice []NS

func (n NSslice) Len() int {
	return len(n)
}

func (n NSslice) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// 数字昇順でソートするクラス
type ByN struct {
	NSslice
}

func (n ByN) Less(i, j int) bool {
	return n.NSslice[i].Num < n.NSslice[j].Num
}

// 文字でソートするクラス
type ByS struct {
	NSslice
}

func (n ByS) Less(i, j int) bool {
	return n.NSslice[i].Str < n.NSslice[j].Str
}

//---------------------------------------------------
//配列を逆順にする
//---------------------------------------------------
/*
sort.Reverse()が使えます。ソートせずに逆順にするのはなさそう。
*/
func slice_Reverse() {
	a := []int{5, 1, 4, 2, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a) // => "[5,4,3,2,1]"
}

//---------------------------------------------------
//指定した位置の要素を取り除く
//---------------------------------------------------
func slice_Delete() {
	a := []int{5, 1, 4, 2, 3}
	d := 0

	d, a, _ = delete(a, 0)
	fmt.Println(d) // => "5"
	fmt.Println(a) // => "[1 4 2 3]"

	d, a, _ = delete(a, 1)
	fmt.Println(d) // => "4"
	fmt.Println(a) // => "[1 2 3]"

}

func delete(slice []int, i int) (int, []int, error) {
	ret := slice[i]
	if len(slice) < i || len(slice) < i {
		return 0, nil, fmt.Errorf("Error")
	}
	ans := make([]int, len(slice))
	copy(ans, slice)

	ans = append(slice[:i], slice[(i+1):]...)

	return ret, ans, nil
}

//---------------------------------------------------
//一致する要素を全て取り除く
//---------------------------------------------------
/*
[こちら](http://ymotongpoo.hatenablog.com/entry/2015/05/20/143004)でご指摘頂きました。ありがとうございます。

*/
func slice_DeleteAll() {
	a := []string{"apple", "orange", "lemon", "apple", "vine"}

	str, a, err := delete_strings(a, "apple")
	fmt.Println(str) // => "apple"
	fmt.Println(a)   // => "[orange lemon vine]"
	fmt.Println(err) // => "<nil>"

	str, a, err = delete_strings(a, "apple")
	fmt.Println(str) // => ""
	fmt.Println(a)   // => "[orange lemon vine]"
	fmt.Println(err) // => "Couldn't find"
}

func delete_strings(slice []string, s string) (string, []string, error) {
	ret := make([]string, len(slice))
	i := 0
	for _, x := range slice {
		if s != x {
			ret[i] = x
			i++
		}
	}
	if len(ret[:i]) == len(slice) {
		return "", slice, fmt.Errorf("Couldn't find")
	}
	return s, ret[:i], nil
}

//---------------------------------------------------
//配列から重複した要素を取り除く
//---------------------------------------------------
/*
これも[golang-set](https://github.com/deckarep/golang-set)
を使います。
*/
//import set "github.com/deckarep/golang-set"

func slice_Uniq() {
	a := []interface{}{30, 20, 50, 30, 10, 10, 40, 50}
	as := set.NewSetFromSlice(a)
	fmt.Println(as) // => Set{30,20,50,10,40}

	s := []interface{}{"/tmp", "/home/", "/etc", "/tmp"}
	ss := set.NewSetFromSlice(s)
	fmt.Println(ss) // =>Set{/tmp, /home/, /etc}
}

//---------------------------------------------------
//配列から指定条件を満たす要素を取り除く
//---------------------------------------------------
/*
条件をいろいろ変えて処理をしたい場合は、無名関数を作って
関数渡しするのがよいでしょう。
*/
func slice_CaseDelete() {
	a := []int{30, 100, 50, 80, 79, 40, 95}

	f0 := func(x int) bool { return x < 80 }
	fmt.Println(reject_map(f0, a)) // => "[100 80 95]"

	f1 := func(x int) bool { return x < 90 }
	fmt.Println(reject_map(f1, a)) // => "[100 95]"
}

// sから、f(x)==true なxを取り除く (f(x)==falseなxの配列を返す）
func reject_map(f func(s int) bool, s []int) []int {
	ans := make([]int, 0)
	for _, x := range s {
		if f(x) == false {
			ans = append(ans, x)
		}
	}
	return ans
}

//---------------------------------------------------
//配列から指定条件を満たす要素を抽出する
//---------------------------------------------------
/*
条件をいろいろ変えて処理をしたい場合は、無名関数を作って
関数渡しするのがよいでしょう。
*/
func slice_CaseSelect() {
	a := []int{1, 2, 3, 4}
	f0 := func(x int) bool { return (x%2 == 0) }
	fmt.Println(select_map(f0, a)) // => [2 4]
}

// sから、f(x)==true なxを返す
func select_map(f func(s int) bool, s []int) []int {
	ans := make([]int, 0)
	for _, x := range s {
		if f(x) == true {
			ans = append(ans, x)
		}
	}
	return ans
}

//---------------------------------------------------
//配列中の要素を探す
//---------------------------------------------------
/*
```
type any []interface{}
```
でany型を設定してあるとします。

うーん、、やはり複雑な配列の処理はrubyみたいにはいかないですね。。
*/
func slice_Search() {
	a := any{"apple", 10, "orange", any{"lemon", "vine"}}

	i, err := index(a, any{"apple"})
	fmt.Println(i)   // => "0"
	fmt.Println(err) // => "<nil>"

	i, err = index(a, any{10})
	fmt.Println(i)   // => "1"
	fmt.Println(err) // => "<nil>"

	i, err = index(a, any{"fruit"})
	fmt.Println(i)   // => "-1"
	fmt.Println(err) // => "Couldn't find"

}

// aにqueryが見つかればindexを返します。
// とりあえずstringとintにしか対応してません。。
func index(a, query any) (int, error) {
	v, ok := query[0].(string)
	vi := -1
	if !ok {
		vi, ok = query[0].(int)
		if !ok {
			return -1, fmt.Errorf("Only string/int query is supported.")
		}
	}

	for i, x := range a {
		xs, ok := x.(string)
		if ok {
			if xs == v {
				return i, nil
			}
		} else {
			xi, ok := x.(int)
			if ok {
				if xi == vi {
					return i, nil
				}
			}
		}
	}
	return -1, fmt.Errorf("Couldn't find")
}

//---------------------------------------------------
//配列の配列を検索する
//---------------------------------------------------
/*
うーんまあ愚直に書きましょうか。。
*/

func slice_Assoc() {
	a := []interface{}{
		[]interface{}{"apple", 100},
		[]interface{}{"vine", 500},
		[]interface{}{"orange", 300},
	}

	fmt.Println(assoc(a, "apple"))  // => "[apple 100]"
	fmt.Println(assoc(a, "orange")) // => "[orange 300]"
	fmt.Println(assoc(a, "peer"))   // => "[]"
}

func assoc(a []interface{}, s string) []interface{} {
	ans := make([]interface{}, 0)
	for _, x := range a {
		xs, _ := x.([]interface{})
		v := xs[0].(string)
		if v == s {
			ans = append(ans, xs...)
		}
	}
	return ans
}

//---------------------------------------------------
//配列の各要素にブロックを実行し配列を作成する
//---------------------------------------------------
/*
ブロックのような記法がないので愚直にやるしかないです。
*/
func slice_Block() {
	a := []int{10, 20, 30, 40, 50}

	b := make([]int, len(a))
	copy(b, a)

	for i, x := range a {
		b[i] = x * 10
	}
	fmt.Println(b) // => "[100 200 300 400 500]"
	fmt.Println(a) // => "[10 20 30 40 50]"
}

//---------------------------------------------------
//配列の各要素に対して繰り返しブロックを実行する
//---------------------------------------------------
/* こちらも同様。愚直にやるしかないです。*/
func slice_Block2() {
	a := []string{"Taro", "Hanako", "Ichiro"}
	for _, x := range a {
		fmt.Println("Hello,", x)
	}
	// => "Hello, Taro"
	// => "Hello, Hanako"
	// => "Hello, Ichiro"
}

//---------------------------------------------------
//配列の要素の和を求める
//---------------------------------------------------
func slice_Sum() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, x := range a {
		sum += x
	}
	fmt.Println(sum) // => "55"
}

//---------------------------------------------------
//配列の要素をランダムに抽出する
//---------------------------------------------------
// import "time"
// import "math/rand"
func slice_Choice() {
	a := []int{1, 2, 3}
	fmt.Println(choice(a))
	fmt.Println(choice(a))
}

func choice(s []int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}

//---------------------------------------------------
//複数の配列を同時に動かす
//---------------------------------------------------

func slice_ThreeItems() {
	fruits := []string{"mango", "apple", "orange"}
	prices := []int{200, 120, 100}
	counts := []int{5, 10, 12}

	for i, _ := range fruits {
		fmt.Println(fruits[i], prices[i], counts[i], prices[i]*counts[i])
	}
	//=>	"mango 200 5 1000"
	//=>	"apple 120 10 1200"
	//=>	"orange 100 12 1200"
}

//---------------------------------------------------
//二次元，三次元の座標の配列の成分ごとの最大，最小を求める
//---------------------------------------------------
/*
matrix.goのColCopy(), RowCopy()で指定行・列の配列を取り出せます。
*/
// import matrix "github.com/skelterjohn/go.matrix"

func slice_MatMax() {
	a, _ := matrix.ParseMatlab("[1 5;8 4;2 9;4 3]")
	x := a.ColCopy(0)   // => [1 8 2 4]
	y := a.ColCopy(1)   // => [5 4 9 3]
	fmt.Println(max(x)) // => "8"
	fmt.Println(max(y)) // => "9"
}

func max(a []float64) float64 {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

//---------------------------------------------------
// 配列
//---------------------------------------------------
func Tips_slice() {

	slice_Define()        // プログラムで配列を定義する
	slice_SliceOfSlice()  // m x n 行列の形で配列の配列を初期化する
	slice_Join()          // 配列要素をカンマ区切りで出力する
	slice_Count()         //配列の要素数を取得する
	slice_Append()        //配列に要素を追加する
	slice_Pop()           //配列の先頭または末尾から要素を取りだす
	slice_Slice()         //部分配列を取りだす
	slice_Fill()          //配列を任意の値で埋める
	slice_Clear()         //配列を空にする
	slice_Concat()        //配列同士を結合する
	slice_Union()         //配列同士の和・積を取る
	slice_Replace()       //複数の要素を変更する
	slice_Flatten()       //配列の配列をフラットな配列にする
	slice_Sort()          //配列をソートする
	slice_CaseSort()      //条件式を指定したソート
	slice_SortAnyColumn() //配列の配列を任意の要素でソートする
	slice_Reverse()       //配列を逆順にする
	slice_Delete()        //指定した位置の要素を取り除く
	slice_DeleteAll()     //一致する要素を全て取り除く
	slice_Uniq()          //配列から重複した要素を取り除く
	slice_CaseDelete()    //配列から指定条件を満たす要素を取り除く
	slice_CaseSelect()    //配列から指定条件を満たす要素を抽出する
	slice_Search()        //配列中の要素を探す
	slice_Assoc()         //配列の配列を検索する
	slice_Block()         //配列の各要素にブロックを実行し配列を作成する
	slice_Block2()        //配列の各要素に対して繰り返しブロックを実行する
	slice_Sum()           //配列の要素の和を求める
	slice_Choice()        //配列の要素をランダムに抽出する
	slice_ThreeItems()    //複数の配列を同時に動かす
	slice_MatMax()        //二次元，三次元の座標の配列の成分ごとの最大，最小を求める

}
