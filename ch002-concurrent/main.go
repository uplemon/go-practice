package main

import "go-practice/ch002-concurrent/concurrent"

func main() {
	concurrent.GoroutineDemo()
	concurrent.SyncDemo()
	concurrent.ContextDemo()
}
