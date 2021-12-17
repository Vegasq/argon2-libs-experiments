package main

import (
	"fmt"
	"github.com/matthewhartstonge/argon2"
	"sync"
	"time"
)

func main(){
	now := time.Now();
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		fmt.Print(".");
		wg.Add(1);
		go cracker(&wg, i);
	}
	wg.Wait();

	fmt.Printf("Done in %f\n", time.Since(now).Seconds());
}

func cracker(wg *sync.WaitGroup, i int){
	hash := []byte("$argon2id$v=19$m=16,t=2,p=1$MTExMTExMTE$SzwiO6Uix4CqutzHAncBwQ");
	pwd := []byte(fmt.Sprintf("password%d", i));
	_, err := argon2.VerifyEncoded(pwd, hash);
	if err != nil {
		fmt.Println(err);
		panic(err)
	}
	fmt.Print(".\n");
	wg.Done();
}
