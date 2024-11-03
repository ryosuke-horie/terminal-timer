package main

import (
	"fmt"
	"os"
	"strconv" // 文字列からint,その逆の変換を行うパッケージ
	"time"    // 時間の測定と表示
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("使用法： go run timer.go 分 秒")
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("分が指定されていません")
		return
	}

	sec, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("秒が指定されていません")
		return
	}

	// 合計の秒数に変換
	totalseconds := min*60 + sec

	// タイマーを作成
	ticker := time.NewTicker(1 * time.Second)
	// 最後に必ず停止させる
	defer ticker.Stop()

	// チャネル: Goの平行処理を実現するため
	// パイプのようなもので、一方のgoroutineがチャネルにデータを送信し、他方のgoroutineがデータを受信する
	// chan boolなのでブール値を送受信するチャネル
	done := make(chan bool)

	// goroutineの開始
	go func() {
		for {
			select {
				case <-done:
					// goroutineを終了
					return

				// ticker.Cはタイマーのチャネル 
				case t := <-ticker.C:
					// 残り時間 ： 指定時間 - 経過時間
					remaining := totalseconds - int(t.Sub(t).Seconds())
					if remaining <= 0 {
						// doneチャネルにtrueを送信
						done <- true
						return
					}

					// 後で変数名を変える
					mins := totalseconds / 60
					secs := totalseconds % 60
					fmt.Printf("\r残り時間: %02d:%02d", mins, secs)

					totalseconds--
					if totalseconds < 0 {
						// doneチャネルにtrueを送信
						done <- true
						return
					}
				}
			}
		}()

	<-done
	fmt.Printf("\n指定時間%02d:%02dが経過しました", min, sec)
}
