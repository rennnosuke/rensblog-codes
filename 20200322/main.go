package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	memStats()
}

func memStats() {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// ヒープ上に割り当てられたオブジェクト累積メモリ量
	fmt.Printf("MAlloc : %v\n", humanize.Bytes(m.Mallocs))
	// ヒープ上から開放されたオブジェクト数
	fmt.Printf("Frees : %v\n", m.Frees)

	// ヒープ上に割り当てられたオブジェクトメモリ量
	fmt.Printf("Alloc : %v\n", humanize.Bytes(m.Alloc))
	fmt.Printf("HeapAlloc : %v\n", humanize.Bytes(m.HeapAlloc))
	// ヒープ上に割り当てられたオブジェクトメモリ量。ただし開放されたオブジェクト分も含む
	fmt.Printf("TotalAlloc : %v\n", humanize.Bytes(m.TotalAlloc))

	// OSから割り当てられたプロセスの総メモリ量
	// ヒープ + スタック + その他
	fmt.Printf("Sys : %v\n", humanize.Bytes(m.Sys))

	// ポインタのルックアップ数
	fmt.Printf("Lookups : %v\n", m.Lookups)

	// 到達可能、あるいはGCによって解放されていないヒープオブジェクトメモリ量
	fmt.Printf("HeapAlloc : %v\n", humanize.Bytes(m.HeapAlloc))
	// 未使用ヒープ領域メモリ量
	fmt.Printf("HeapIdle : %v\n", humanize.Bytes(m.HeapIdle))
	// 使用中ヒープ領域メモリ量
	fmt.Printf("HeapInuse : %v\n", humanize.Bytes(m.HeapInuse))
	// OSに返却される物理メモリ量
	fmt.Printf("HeapReleased : %v\n", humanize.Bytes(m.HeapReleased))
	// ヒープに割り当てられたオブジェクト量
	fmt.Printf("HeapObjects : %v\n", m.HeapObjects)

	// 使用中スタック領域メモリ量
	fmt.Printf("StackInuse : %v\n", humanize.Bytes(m.StackInuse))
	// OSから割り当てられたスタック領域メモリ量
	fmt.Printf("StackSys : %v\n", humanize.Bytes(m.StackSys))

	// 割り当てられたmspan構造体バイト数
	fmt.Printf("MSpanInuse : %v\n", humanize.Bytes(m.MSpanInuse))
	// OSから取得したmspan構造体バイト数
	fmt.Printf("MSpanSys : %v\n", humanize.Bytes(m.MSpanSys))

	// 割り当てられたmcache構造体バイト数
	fmt.Printf("MCacheInuse : %v\n", humanize.Bytes(m.MCacheInuse))
	// OSから取得したmcache構造体バイト数
	fmt.Printf("MCacheSys : %v\n", humanize.Bytes(m.MCacheSys))


	//
	fmt.Printf("BuckHashSys : %v\n", humanize.Bytes(m.BuckHashSys))


	go func() { time.Sleep(time.Second * 10) }()
	go func() { time.Sleep(time.Second * 10) }()
	go func() { time.Sleep(time.Second * 10) }()
	// 呼び出し時存在するゴルーチン数
	fmt.Printf("Goroutine : %v\n", runtime.NumGoroutine())

	sr := []runtime.StackRecord{
		{},
		{},
		{},
		{},
	}
	n, ok := runtime.GoroutineProfile(sr)
	fmt.Printf("n-GoroutineProfile : %d\n", n)
	if ok {
		for i, p := range sr {
			fmt.Printf("GoroutineProfile-%d: %v\n", i, p.Stack0)
		}
	}

}
