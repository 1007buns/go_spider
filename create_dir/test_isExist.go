/* package main
import (
    "os"
    "fmt"
)
//判断文件或文件夹是否存在
func isExist(path string)(bool){
    _, err := os.Stat(path)
    if err != nil{
        if os.IsExist(err){
            return true
        }
        if os.IsNotExist(err){
            return false
        }
        fmt.Println(err)
        return false
    }
    return true
}
func main(){
    //递归创建文件夹
    err := os.MkdirAll("./test/1/2", os.ModePerm)
    if err != nil{
        fmt.Println(err)
        return
    }
    dirs := []string{"./test/1", "./test/2", "./test/1.txt"}
    for _, v := range dirs{
        if isExist(v){
            fmt.Printf("%s is exist!", v)
        }else{
            fmt.Printf("%s is not exist!", v)
        }
    }
}
*/