package main

import (
	"fmt"
	_ "fmt"
)

func main(){
	var question1 int
	fmt.Println("今日の天気はどうですか？ 1 = 晴れ 2 = 雨 :")
	_, _ = fmt.Scan(&question1)
	if question1 == 1 {
		fmt.Println("そうですか...良かったです")
	}
	if question1 == 2 {
		fmt.Println("そうですか...明日は晴れるといいですね")
	}

	var question2 int
	fmt.Println("今日の気分はどうですか？ 1 = 結構いい 2 = 良くない :")
	_, _ = fmt.Scan(&question2)
	if question2 == 1 {
		fmt.Println("そうですか。最高の一日になるといいですね。")
	}
	if question2 == 2 {
		fmt.Println("そうですか...リフレッシュできるといいですね。")
	}
}
