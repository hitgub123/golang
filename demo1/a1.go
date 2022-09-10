package main

import (
	"fmt"
	// "os"
	"os/exec"
	// "time"
)

func main() {
	// pname:="D:\\program\\WinMerge\\WinMergeU.exe"
	pname:="d: & cd d:/workspace/go/demo1 & go run a1.go"
	fmt.Println(pname)
	cmd:= exec.Command(pname)
    err:=cmd.Run()
	fmt.Println("error:", err)

    // _, e, _ := os.Pipe()
	// attr := &os.ProcAttr{Env: os.Environ(), Files: []*os.File{nil, e, nil}}
	// p,err:=os.StartProcess(pname,[]string{pname},attr)
	// if err!= nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v",p)
}
