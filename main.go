/*
逆引きGolang

[逆引きRuby](http://www.namaraii.com/rubytips)の内容をGolang化したものです。
[こちら](http://ashitani.jp/golangtips)で公開しています。

*/

package main

import (
	. "./pkg/tips_map"
	. "./pkg/tips_num"
	. "./pkg/tips_regexp"
	. "./pkg/tips_slice"
	. "./pkg/tips_string"
	. "./pkg/tips_time"
)

func main() {
	Tips_string()
	Tips_num()
	Tips_time()
	Tips_slice()
	Tips_map()
	Tips_regexp()
}
