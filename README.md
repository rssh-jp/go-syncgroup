# go-syncgroup

## usage
``` main.go
package main

import(
    "github.com/rssh-jp/go-syncgroup"
)

func main(){
    sg := syncgroup.New(4)
    for i:=0; i<10; i++{
        sg.Add()
        go func(){
            defer sg.Done()
            execute()
        }()
    }

    sg.Wait()
}
    
```

