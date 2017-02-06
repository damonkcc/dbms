package main

import (
    "fmt"
   lc "epicgames.com/dbms/ldap"
)


func main(){
   
   fmt.Printf("Hello World\n")
   lc.Client_Authenticate("username","password")
}
