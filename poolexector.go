package fourinone

import (
	"sync"
)

var(
	threadPoolExecutor
)
type Runnable interface{}

type poolExector struct{
	
}
func threadPoolExecutor(){
	mu:=sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
}
func scheduledThreadPoolExecutor{
	mu:=sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
}
func execute(d,i Runnable,t int64){
	
}
func close(){
	mu:=sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
}