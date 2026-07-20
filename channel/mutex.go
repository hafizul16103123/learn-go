// package main
// import(
// 	"fmt"
// 	"sync"
// )
// type Post struct{
// 	views int
// 	mu sync.Mutex
// }
// func(p *Post) inc( wg *sync.WaitGroup){
// 	defer func(){
// 		p.mu.Unlock()
// 		wg.Done()
// 	}()
// 	p.mu.Lock()
// 	p.views+=1
// }

// func main(){
// 	var wg sync.WaitGroup

// 	p:=Post{views:0}

// 	for i:=1;i<=100;i++{
// 		wg.Add(1)
// 		go p.inc(&wg)
// 	}
	
// 	wg.Wait()
// 	fmt.Println(p.views)


	
// }



