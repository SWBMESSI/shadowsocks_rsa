package main

import (
    "os"
    "fmt"
    "flag"
    ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
)

func main(){

    var msg ss.MsgConf
    var builds bool 

   // fmt.Println(builds)
    flag.StringVar(&msg.Msgpath,"e","onfig.json","a file need to encrypt")
    flag.BoolVar(&builds, "build", false, "build privatekey and pubkey")
   // flag.StringVar(&msg.Breakpath,"b","onfig.json","a file need to break")
    //flag.StringVar(&msg.Msgpath,"e","","a file need to encrypt")
    flag.Parse()
  //  fmt.Println(builds)
    if builds==true {
         fmt.Printf("-------------------------------------------------------\n")
         ss.RSA()
        fmt.Printf("-------------------------------------------------------\n")
        os.Exit(0)
    }/*
    if msg.Breakpath!=""{
        fmt.Printf("-------------------------------------------------------\n")
        ss.BreakMsg(msg.Breakpath)
    }*/
    if msg.Msgpath!=""{
        fmt.Printf("-------------------------------------------------------\n")
        ss.RsaMsg(msg.Msgpath)
    }


}