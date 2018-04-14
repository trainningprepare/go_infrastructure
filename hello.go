package main

import "fmt"
import (
  "os"
  "os/exec"
)


func init()  {

}


func main(){

  cmd:=exec.Command("D:\\go\\src\\main.exe")
  err:=cmd.Start()
  if err!=nil {
    fmt.Println(err.Error())
  }

  fmt.Fprint(os.Stdout,"my name is ","wenweiping go \n");
  fmt.Println("this is branch: go_infrastructure_20180413")
  fmt.Println("886")
}


// this git clinet lenovo  branch master 


