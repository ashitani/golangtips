/*
日付と時刻
*/

package tips_time

import (
	"fmt"
	"time"
)

//---------------------------------------------------
// 現在の時刻を取得する
//---------------------------------------------------
//import "time"

func time_Now() {
	t := time.Now()
	fmt.Println(t)           // => "2015-05-05 07:23:30.757800829 +0900 JST"
	fmt.Println(t.Year())    // => "2015"
	fmt.Println(t.Month())   // => "May"
	fmt.Println(t.Day())     // => "5"
	fmt.Println(t.Hour())    // => "7"
	fmt.Println(t.Minute())  // => "23"
	fmt.Println(t.Second())  // => "30"
	fmt.Println(t.Weekday()) // => "Tuesday"
}

//---------------------------------------------------
// 時刻オブジェクトを作成する
//---------------------------------------------------
/* 秒の後にナノセカンドまで指定するのを忘れがちなので注意。*/
//import "time"

func time_Make() {
	t := time.Date(2001, 5, 20, 23, 59, 59, 0, time.UTC)
	fmt.Println(t) // => "2001-05-20 23:59:59 +0000 UTC"
	t = time.Date(2001, 5, 20, 23, 59, 59, 0, time.Local)
	fmt.Println(t) // => "2001-05-20 23:59:59 +0900 JST"
}

//---------------------------------------------------
// 時刻を任意のフォーマットで扱う
//---------------------------------------------------
/*
time.FormatはrubyのTime::strftimeと違い、
%Hなどの制御文字で行うのではなく例文で書きます。

 "Mon Jan 2 15:04:05 -0700 MST 2006"

を並べ替える仕様のようです。なんでこの時刻？ってのは
[こちら](http://qiita.com/ruiu/items/5936b4c3bd6eb487c182)
にありました。
*/
//import "time"

func time_Format() {
	t := time.Now()
	const layout = "Now, Monday Jan 02 15:04:05 JST 2006"
	fmt.Println(t.Format(layout)) // => "Now, Tuesday May 05 07:53:54 JST 2015"
	const layout2 = "2006-01-02 15:04:05"
	fmt.Println(t.Format(layout2)) // => "2015-05-05 07:53:54"

}

//---------------------------------------------------
// 時刻オブジェクトを文字列に変換する
//---------------------------------------------------
//import "time"

func time_ToString() {
	t := time.Now()
	s := ""
	s = t.String()
	fmt.Println(s) // => "2015-05-05 08:05:15.828452891 +0900 JST"
}

//---------------------------------------------------
// 時刻に任意の時間を加減する
//---------------------------------------------------
/*
時間差を表現する型はtime.Durationで、単位はナノ秒です。
time.Second,time.Hour等を掛け算することで秒単位、時間単位を表現します。

あるいはtime.ParseDuration("100s")などで生成するのも便利です。

減算したい場合は負の値をAdd()すればよいです。

Sub()はtime.Time同士を引き算してtime.Durationを返す関数です。
*/
//import "time"

func time_IncDec() {

	t := time.Date(2001, 5, 20, 23, 59, 59, 0, time.Local)
	t = t.Add(time.Duration(1) * time.Second)
	fmt.Println(t) // => "2001-05-21 00:00:00 +0900 JST"

	t1 := time.Date(2000, 12, 31, 0, 0, 0, 0, time.Local)
	t1 = t1.Add(time.Duration(24) * time.Hour)
	fmt.Println(t1) // => "2001-01-01 00:00:00 +0900 JST"

}

//---------------------------------------------------
// 2つの時刻の差を求める
//---------------------------------------------------
//import "time"

func time_Duration() {
	day1 := time.Date(2000, 12, 31, 0, 0, 0, 0, time.Local)
	day2 := time.Date(2001, 1, 2, 12, 30, 0, 0, time.Local)
	duration := day2.Sub(day1)
	fmt.Println(duration) // => "60h30m0s"

	hours0 := int(duration.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60
	fmt.Printf("%d days + %d hours + %d minutes + %d seconds\n", days, hours, mins, secs)
	// => "2 days + 12 hours + 30 minutes + 0 seconds"
}

//---------------------------------------------------
// 時刻中の曜日を日本語に変換する
//---------------------------------------------------
/*
time.Weekday は普段はintだけどPrintln等でString()メソッドを
呼ばれると曜日の英語表記になるという型です。
*/
//import "time"

func time_JapaneseWeekday() {
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}

	t := time.Now()
	fmt.Println(t.Weekday())        // =>"Tuesday"
	fmt.Println(wdays[t.Weekday()]) // =>"火"
}

//---------------------------------------------------
// UNIXタイムをTimeオブジェクトに変換する
//---------------------------------------------------
//import "time"

func time_Unix() {
	fmt.Println(time.Unix(1267867237, 0)) // => "2010-03-06 18:20:37 +0900 JST"
	fmt.Println(time.Now().Unix())        // => "1430783059"
}

//---------------------------------------------------
// 現在の日付を求める
//---------------------------------------------------
/* Rubyのようにtime/dateの違いはありません。*/
//import "time"

func time_Date() {
	day := time.Now()
	const layout = "2006-01-02"
	fmt.Println(day.Format(layout)) // => "2015-05-05"
}

//---------------------------------------------------
// 日付オブジェクトを文字列に変換する
//---------------------------------------------------
//import "time"

func time_DateString() {
	day := time.Now()
	const layout = "2006-01-02"
	fmt.Println(day.Format(layout)) // => "2015-05-05"
}

//---------------------------------------------------
// 日付オブジェクトを作成する
///---------------------------------------------------
//import "time"

func time_MakeDate() {
	day := time.Date(2001, 5, 31, 0, 0, 0, 0, time.Local)
	fmt.Println(day) // => "2001-05-31 00:00:00 +0900 JST"
}

//---------------------------------------------------
// 指定の日付が存在するかどうか調べる
//---------------------------------------------------
/*
うーん、別途関数を書くしか無いですね。。

ユリウス日の変換は[こちら](http://play.golang.org/p/ocYFWY7kpo)を参照しました。
*/
//import "time"

func time_Exist() {
	jd, err := isExist(2001, 1, 31)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int(jd))
	}
	// => "2451940"
	jd, err = isExist(2001, 1, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int(jd))
	}
	// => "2001-1-32 is not exist"
}

// 指定の日付が存在するかどうか調べる。
// 存在しない日付を指定してもtime.Date()はよきに計らってくれるので、
// 指定した日付と違うtime.Timeが返ってくれば指定した日付が間違ってると判定。
func isExist(year, month, day int) (float64, error) {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if date.Year() == year && date.Month() == time.Month(month) && date.Day() == day {
		return Julian(date), nil
	} else {
		return 0, fmt.Errorf("%d-%d-%d is not exist", year, month, day)
	}
}

// ユリウス日を求める
func Julian(t time.Time) float64 {
	// Julian date, in seconds, of the "Format" standard time.
	// (See http://www.onlineconversion.com/julian_date.htm)
	const julian = 2453738.4195
	// Easiest way to get the time.Time of the Unix time.
	// (See comments for the UnixDate in package Time.)
	unix := time.Unix(1136239445, 0)
	const oneDay = float64(86400. * time.Second)
	return julian + float64(t.Sub(unix))/oneDay
}

//---------------------------------------------------
// ユリウス日から日付オブジェクトを作成する
//---------------------------------------------------
/*　なんだか誤差が残りますが..。*/
//import "time"

func time_FromJulian() {
	t := time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local)
	jd := Julian(t)
	fmt.Println(jd)             // => "2.451940124997685e+06"
	fmt.Println(FromJulian(jd)) // => "2001-01-31 00:00:00.000014592 +0900 JST"
}

// ユリウス日から時間に
func FromJulian(jd float64) time.Time {
	const julian = 2453738.4195
	unix := time.Unix(1136239445, 0)
	const oneDay = float64(86400. * time.Second)
	return unix.Add(time.Duration((jd - julian) * oneDay))
}

//---------------------------------------------------
// 何日後、何日前の日付を求める
//---------------------------------------------------
/* AddDateが使えます。引数はyear,month,dayの順です。*/
//import "time"

func time_IncDecDay() {
	t := time.Date(2001, 5, 31, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, 1)
	fmt.Println(t) // => "2001-06-01 00:00:00 +0900 JST"

	t = time.Date(2001, 1, 1, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, -1)
	fmt.Println(t) // => "2000-12-31 00:00:00 +0900 JST"
}

//---------------------------------------------------
// 何ヶ月後、何ヶ月前の日付を求める
//---------------------------------------------------
/*
rubyは 1/31 -> 2/28など月末を意識したメソッドのようですね。
AddDate()で単純に１ヶ月足すと、2/31 = 3/3 のように変換されて
しまいます。ので、最終日を意識した加減のメソッドAddMonthを作ります。
*/
//import "time"

func time_IncDecMonth() {
	t0 := time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local)
	t1 := t0.AddDate(0, 1, 0)
	t2 := AddMonth(t0, 1)
	fmt.Println(t1) // => "2001-03-03 00:00:00 +0900 JST"
	fmt.Println(t2) // => "2001-02-28 00:00:00 +0900 JST"
	t0 = time.Date(2001, 5, 31, 0, 0, 0, 0, time.Local)
	t1 = t0.AddDate(0, -1, 0)
	t2 = AddMonth(t0, -1)
	fmt.Println(t1) // => "2001-05-01 00:00:00 +0900 JST"
	fmt.Println(t2) // => "2001-04-30 00:00:00 +0900 JST"
}

// ruby風の月加算。これだとうるう秒に対応できてませんが。。
func AddMonth(t time.Time, d_month int) time.Time {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	newMonth := int(month) + d_month
	newLastDay := getLastDay(year, newMonth)
	var newDay int
	if day > newLastDay {
		newDay = newLastDay
	} else {
		newDay = day
	}

	return time.Date(year, time.Month(newMonth), newDay, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

}

// その月の最終日を求める
func getLastDay(year, month int) int {
	t := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, -1)
	return t.Day()
}

//---------------------------------------------------
// うるう年かどうか判定する
//---------------------------------------------------
func time_LeapYear() {
	fmt.Println(isLeapYear(2000)) // => "true"
	fmt.Println(isLeapYear(2001)) // => "false"
}

// うるう年かどうか判定する
func isLeapYear(year int) bool {
	if year%400 == 0 { // 400で割り切れたらうるう年
		return true
	} else if year%100 == 0 { // 100で割り切れたらうるう年じゃない
		return false
	} else if year%4 == 0 { // 4で割り切れたらうるう年

		return true
	} else {
		return false
	}
}

//---------------------------------------------------
// 日付オブジェクトの年月日・曜日を個別に扱う
//---------------------------------------------------
//import "time"

func time_Decompose() {
	t := time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local)
	fmt.Println(t.Year())    // => "2001"
	fmt.Println(t.Month())   // => "January"
	fmt.Println(t.Day())     // => "31"
	fmt.Println(t.Weekday()) // => "Wednesday"
}

//---------------------------------------------------
// 文字列の日付を日付オブジェクトに変換する
//---------------------------------------------------
/*
 Parseもフォーマット整形と同じように

 "Mon Jan 2 15:04:05 -0700 MST 2006"

 をひな形(layout)にして与えます
*/
//import "time"

func time_Parse() {
	str := "Thu May 24 22:56:30 JST 2001"
	layout := "Mon Jan 2 15:04:05 MST 2006"
	t, _ := time.Parse(layout, str)
	fmt.Println(t) // => "2001-05-24 22:56:30 +0900 JST"

	str = "2003/04/18"
	layout = "2006/01/02"
	t, _ = time.Parse(layout, str)
	fmt.Println(t) // => "2003-04-18 00:00:00 +0000 UTC"
}

//---------------------------------------------------
// 日付と時刻
//---------------------------------------------------
func Tips_time() {
	time_Now()             // 現在の時刻を取得する
	time_Make()            // 時刻オブジェクトを作成する
	time_Format()          // 時刻を任意のフォーマットで扱う
	time_ToString()        // 時刻オブジェクトを文字列に変換する
	time_IncDec()          // 時刻に任意の時間を加減する
	time_Duration()        // 2つの時刻の差を求める
	time_JapaneseWeekday() // 時刻中の曜日を日本語に変換する
	time_Unix()            // UNIXタイムをTimeオブジェクトに変換する
	time_Date()            // 現在の日付を求める
	time_DateString()      // 日付オブジェクトを文字列に変換する
	time_MakeDate()        // 日付オブジェクトを作成する
	time_Exist()           // 指定の日付が存在するかどうか調べる
	time_FromJulian()      // ユリウス日から日付オブジェクトを作成する
	time_IncDecDay()       // 何日後、何日前の日付を求める
	time_IncDecMonth()     // 何ヶ月後、何ヶ月前の日付を求める
	time_LeapYear()        // うるう年かどうか判定する
	time_Decompose()       // 日付オブジェクトの年月日・曜日を個別に扱う
	time_Parse()           // 文字列の日付を日付オブジェクトに変換する

}
