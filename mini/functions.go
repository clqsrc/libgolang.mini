
//常用函数

//package functions//算了 golang 不能是项目的相对目录
package main

import (
//	"container/list"
//	"database/sql"
//	"fmt"
//	"reflect"
	"strconv"
	"strings"
	"time"
	"os"
	"io"
	"io/ioutil"
	"path/filepath"
	"crypto/md5"
    "fmt"
    "encoding/hex"
	"encoding/base64"
	"net/url"
	"net/http"
	"mime/multipart"
	"html/template"
	"crypto/tls"
	"net"
	"unsafe"
//	"sync"

	//_ "github.com/bmizerany/pq"
	//_ "github.com/lib/pq" //驱动的写法一定要这样写,否则会当做无效的导入
)

func Functions_StrToInt(s string) int {
	
	//return (int)(StrToIntDef(s, 0)); //golang 连 int64 和 int 都不能强制转换
	return int(StrToIntDef(s, 0)); //golang 连 int64 和 int 都不能强制转换
}//

func Functions_StrToInt64(s string) int64 {
	
	return StrToIntDef(s, 0);
}//

//有缺省值的字符串向整数转换
func StrToIntDef(s string, def int64) int64{
	
	var r int64 = def;
	
	if (len(s)<1) { return r; }
	
	//var s string = "9223372036854775807"
	i, err := strconv.ParseInt(s, 10, 64);
	if err != nil {
	    //panic(err)
		i = def;
	}
	//fmt.Printf("Hello, %v with type %s!\n", i, reflect.TypeOf(i));
	
	r = i;
	
	return r;
}//

//有缺省值的字符串向 float 转换
func StrToFloatDef(s string, def float64) float64{
	
	var r float64 = def;
	
	if (len(s)<1) { return r; }
	
	//var s string = "9223372036854775807"
	//i, err := strconv.ParseInt(s, 10, 64);
	i, err := strconv.ParseFloat(s, 64);
	if err != nil {
	    //panic(err)
		i = def;
	}
	//fmt.Printf("Hello, %v with type %s!\n", i, reflect.TypeOf(i));
	
	r = i;
	
	return r;
}//

func StrToFloat(s string) float64{
	
	return StrToFloatDef(s, 0.0);
	
}//

//有缺省值的字符串向整数转换
func StrToIntDef_(s string, def int) int{
	
	var r int = def;
	
	if (len(s)<1) { return r; }
	
	//var s string = "9223372036854775807"
	i, err := strconv.ParseInt(s, 10, 64);
	if err != nil {
	    //panic(err)
		i = int64(def);
	}
	//fmt.Printf("Hello, %v with type %s!\n", i, reflect.TypeOf(i));
	
	r = int(i);
	
	return r;
}//

//整数转换为字符串
func IntToStr(v int64) string{
	return  strconv.FormatInt(v, 10);   
	
}//

//小数转换为字符串
func FloatToStr(v float64) string{
	
	////fmt表示格式： "f"/"b"/"e"/"E"/"g"/"G" 包括这六种格式，默认为f
	//prec表示精度，自己控制
	//bitSize表示输入的是float32 还是float64，32/64.
	//fmt.Println(strconv.FormatFloat(vfloat, 'e', 4, 64))
	
	return  strconv.FormatFloat(v, 'f', 4, 64);   
	
}//

//整数转换为字符串
func IntToStr_(v int) string{
	return  strconv.FormatInt(int64(v), 10);   
	
}//

//sql用的 
func last_time_float() string {
	
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05")) # 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	//b1 := time.Now().UTC()//不带时区用这个
	//return time.Now().Format("20060102.150405"); //这个是北京时间
	return time.Now().UTC().Format("20060102.150405"); //这个是北京时间

}//

//文件用的 
func Now_id() string {
	
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05")) # 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	//b1 := time.Now().UTC()//不带时区用这个
	//return time.Now().Format("20060102.150405"); //这个是北京时间
	//return time.Now().UTC().Format("20060102_150405_999999999"); //这个是北京时间
	
	//time.Now().Format("20060102_150405_.000.999999999")//ok 目前的 golang 版本在 000 和 999999999 前都要有一个点,知是为何,以后再研究了
	//r :=  time.Now().UTC().Format("20060102_150405.000.999999999"); //这个不是北京时间//目前的 golang 版本对于时间的毫秒 000 前面必须有个点, 999999999 也是

	r := time.Now().UTC().Format("20060102_150405.000.999999999"); //这个不是北京时间
	
	r = str_replace(r, ".", "_");
	
	return r;

}//

//等同 c 语言同名函数
func _time() int64 {
	return time.Now().UTC().Unix();
}//

//C 语言秒数转换为北京时间中文提示
func Now_StringCN() string {  
	//DateTimeToUnix
//    fmt.Println(time.Now().Unix()) //获取当前秒  
//    fmt.Println(time.Now().UnixNano())//获取当前纳秒  
//    fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒  
//    fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒  
//    c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型  
//    fmt.Println(c.String()) //输出当前英文时间戳格式  

    ////c := time.Unix(time_c,0); //将毫秒转换为 time 类型  
    //fmt.Println(c.String()) //输出当前英文时间戳格式 
	
	//return c.UTC().Format("2006年01月02日 15点04分05秒"); //这个是秒 //这个不是北京时间
	c := time.Now().UTC();
	
	hh, _ := time.ParseDuration("8h"); //加上8小时
	c_cn := c.Add(hh);	
	
	//fmt.Println(c.String()) //输出当前英文时间戳格式 
	
	return c_cn.UTC().Format("2006年01月02日 15点04分05秒"); //这个是秒 //这个不是北京时间
	
	//如果用 c.Format 也可得到正确的中文时间，但那是 golang 取了操作系统的时区，严格来说不够严谨。还是自己指定时区比较好。

   
}//

//取当前时间的 24 小时点数//time_zone 时区，北京为 8
func Now_hour24(time_zone int) int {

	//----
	now_hour_24 := 9; //当前的小时数，24小时制
	
	//time.Now().UTC().Format("20060102.150405"); //这个是北京时间//不是，是标准时间
	t_now := time.Now().UTC();
	//t_now.Hour() //源码中说 in the range [0, 23]//
	
	now_hour_24 = t_now.Hour(); //源码中说 in the range [0, 23]//
	now_hour_24 = now_hour_24 + time_zone;
	
	//----
	
	return now_hour_24;
}//

//文件用的//只要年月 
func Now_id_ym() string {
	
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05")) # 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	//b1 := time.Now().UTC()//不带时区用这个
	//return time.Now().Format("20060102.150405"); //这个是北京时间
	//return time.Now().UTC().Format("20060102_150405_999999999"); //这个是北京时间
	
	//time.Now().Format("20060102_150405_.000.999999999")//ok 目前的 golang 版本在 000 和 999999999 前都要有一个点,知是为何,以后再研究了
	//r :=  time.Now().UTC().Format("20060102_150405.000.999999999"); //这个是北京时间//目前的 golang 版本对于时间的毫秒 000 前面必须有个点, 999999999 也是

	//r := time.Now().UTC().Format("20060102_150405.000.999999999"); //这个不是北京时间
	r := time.Now().UTC().Format("200601"); //这个不是北京时间
	
	r = str_replace(r, ".", "_");
	
	return r;

}//

//转换为 unix 的 C 语言秒数
func DateTimeToUnix() int64 {  
	//DateTimeToUnix
//    fmt.Println(time.Now().Unix()) //获取当前秒  
//    fmt.Println(time.Now().UnixNano())//获取当前纳秒  
//    fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒  
//    fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒  
//    c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型  
//    fmt.Println(c.String()) //输出当前英文时间戳格式  

	////return time.Now().UTC().Unix(); //这个是北京时间//按道理永远不要将本地时间转换为 unix 秒，unix 秒就应当是标准英国时间
	
	return time.Now().UTC().Unix(); //这个不是北京时间//按道理永远不要将本地时间转换为 unix 秒，unix 秒就应当是标准英国时间
   
}//

//转换为 unix 的 C 语言秒数//按道理永远不要将本地时间转换为 unix 秒，unix 秒就应当是标准英国时间
/*
func DateTimeToUnix_CN() int64 {  
	//DateTimeToUnix
//    fmt.Println(time.Now().Unix()) //获取当前秒  
//    fmt.Println(time.Now().UnixNano())//获取当前纳秒  
//    fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒  
//    fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒  
//    c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型  
//    fmt.Println(c.String()) //输出当前英文时间戳格式  

	return time.Now().UTC().Unix(); //这个不是北京时间//按道理永远不要将本地时间转换为 unix 秒，unix 秒就应当是标准英国时间
   
}//
*/

//C 语言秒数转换为北京时间中文提示
func UnixToDateTime_StringCN(time_c int64) string {  
	//DateTimeToUnix
//    fmt.Println(time.Now().Unix()) //获取当前秒  
//    fmt.Println(time.Now().UnixNano())//获取当前纳秒  
//    fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒  
//    fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒  
//    c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型  
//    fmt.Println(c.String()) //输出当前英文时间戳格式  

    ////c := time.Unix(time_c,0); //将毫秒转换为 time 类型  
    //fmt.Println(c.String()) //输出当前英文时间戳格式 
	
	//return c.UTC().Format("2006年01月02日 15点04分05秒"); //这个是秒 //这个不是北京时间
	c := time.Unix(time_c,0); //将毫秒转换为 time 类型//加上8小时才是北京时间
	
	hh, _ := time.ParseDuration("8h"); //加上8小时
	c_cn := c.Add(hh);	
	
	//fmt.Println(c.String()) //输出当前英文时间戳格式 
	
	return c_cn.UTC().Format("2006年01月02日 15点04分05秒"); //这个是秒 //这个不是北京时间
	
	//如果用 c.Format 也可得到正确的中文时间，但那是 golang 取了操作系统的时区，严格来说不够严谨。还是自己指定时区比较好。

   
}//


//目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//文件是否存在
func FileExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
//	if os.IsNotExist(err) {
//		return false
//	}
	return false
}//

//递归创建目录
func ForceDirectories(path string) (bool) {
	
	//os.MkdirAll(path.Dir(fn))
	//... = os.Create(fn)

   err := os.MkdirAll(path, 0777)
    if err != nil {
        fmt.Printf("%s", err)
		return false;
    } else {
        //fmt.Print("Create Directory OK!")
    }
	
	return true;
}//

//字符串保存为文件
func String2File(s string, fn string) (bool) {
	
	var d1 = []byte(s);
	//err2 := ioutil.WriteFile("./output2.txt", d1, 0666)  //写入文件(字节数组)
	//check(err2)
	
	err2 := ioutil.WriteFile(fn, d1, 0666);  //写入文件(字节数组)
	
	if err2 != nil {
		//panic(e)
		return false;
	}

	return true;
}//

func Bytes2File(data []byte, fn string) (bool) {
	
	//var d1 = []byte(s);
	var d1 = data;
	//err2 := ioutil.WriteFile("./output2.txt", d1, 0666)  //写入文件(字节数组)
	//check(err2)
	
	err2 := ioutil.WriteFile(fn, d1, 0666);  //写入文件(字节数组)
	
	if err2 != nil {
		//panic(e)
		fmt.Println(err2);
		return false;
	}

	return true;
}//


func File2Bytes(fn string, max_size int64) ([]byte) {
	
	defer PrintError("File2Bytes");
	
	fp, err := os.Open(fn) // 获取文件指针
	if err != nil {
		return nil
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		return nil
	}
	
	if max_size<fileInfo.Size() {
		
		fmt.Println("File2Bytes 文件过大保护")
		return nil;
	}
	
	
	buffer := make([]byte, fileInfo.Size());
	
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	if err != nil {
		return nil
	}
	return buffer

}//

func DeleteFile(fn string)(bool) {
    file := fn;//"test.txt"                   //源文件路径
    err := os.Remove(file)               //删除文件test.txt
    if err != nil {
        //如果删除失败则输出 file remove Error!
        fmt.Println("file remove Error!")
        //输出错误详细信息
        fmt.Printf("%s", err);
		return false;
    } else {
        //如果删除成功则输出 file remove OK!
        fmt.Print("file remove OK!")
		return true;
    }

}//

// RemoveAll	删除文件夹或者文件，哪怕不存在，文件夹不为空，都可以删除
func DeleteFileOrDir(fn string)(bool) {
    file := fn;//"test.txt"                   //源文件路径
    err := os.RemoveAll(file)               //删除文件test.txt
    if err != nil {
        //如果删除失败则输出 file remove Error!
        fmt.Println("file remove Error!")
        //输出错误详细信息
        fmt.Printf("%s", err);
		return false;
    } else {
        //如果删除成功则输出 file remove OK!
        fmt.Print("file remove OK!")
		return true;
    }

}//

////字符串保存为文件
//func SaveToFile(s string, fn string) (bool) {
//
//	return String2File(s, fn);
//}//

//检查是否有空格这样不合法的字符串
func CheckSafeName(s string) (bool) {
	
    for i:=0; i < len(s); i++{

		c := s[i];
        //fmt.Printf("%c", s[i])
		if (c == ' ')  { return false; }
		if (c == '/')  { return false; }
		if (c == '\\') { return false; }
		if (c == '.')  { return false; }
		
		if (c>128)  { return false; }//判断是否为汉字
		if(!((c>='0' && c<='9') || (c>='a' && c<='z') || (c>='A' && c<='Z') || (c=='_') || (c=='-'))) { return false; }
		
//            if(ord($str_cn[$i])>128)//判断是否为汉字
//            {
//                return false;
//            }
            
//            $c = $str_cn[$i];
 
//            if(!(($c>='0' && $c<='9') || ($c>='a' && $c<='z') || ($c>='A' && $c<='Z') || ($c=='_') || ($c=='-')))
//            {
//                return false;
//            }		

    }
	
	//boolean isNum = str.matches("[0-9]+");
	//isNum := s.matches("[A-Za-z0-9_]+");
	
	//if (false == isNum) { return false; }
	
	return true;
	
	//--------------------------------------------------
	//以上遍历单个字节,因为 golang 是 utf8 的,所以有可能为双字节,所以遍历单字符为
	//s := "Hello 世界！"
	//for i, v := range s { // i 是字符的字节位置，v 是字符的拷贝
	//fmt.Printf("%2v = %c\n", i, v) // 输出单个字符
	//}
	
}//

//ExtractFileName:返回完整文件名中的文件名称 (带扩展名)，如"mytest.doc"
//ExtractFileExt 返回完整文件名中的文件扩展名（带.），如".doc"
func ExtractFileExt(s string) (string) {
	
	r := "";
	
    //for i:=0; i < len(s); i++{
//    for i:=len(s)-1; i >=0; i--{

//		c := s[i]; //这样得到的是 byte
		
//        //fmt.Printf("%c", s[i])
////		if (c == ' ')  { return false; }
//		if (c == '/')  { return r; }
//		if (c == '\\') { return r; }
//		if (c == '.')  { r = c + r; return r; }

		
//		r = c + r;

//    }
	
	//-----------------
	//str := "烫烫烫烫";
	//array := []rune(str);
//	str := "./\\";
//	array := []rune(str);
	
//	ch1 := array[0];
//	ch2 := array[1];
//	ch3 := array[2];
	
	//这样遍历不能从后向前
    //for i, ch := range s {
	//	fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
	
	array := []rune(s);
	
    for i:=len(array)-1; i >=0; i--{
		ch := array[i];
		fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
		
		//if (ch == (string(".")[0])) { return r; }
		//if (ch == ch1) { return r; }
		
		if (ch == rune('/'))  {  return r; }
		if (ch == rune('\\')) {  return r; }
		if (ch == rune('.'))  { r = string(ch) + r; return r; }
		
		r = string(ch) + r;
	}	

	
	return r;

	
}//

//2022 奇怪，以前应该实现过了的
func ExtractFileName(s string) (string) {
	
	r := "";
	
	//一定要参考早期的 ExtractFileExt() 函数
	//-----------------

	//	fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
	
	array := []rune(s);
	
    for i:=len(array)-1; i >=0; i--{
		ch := array[i];
		//fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
		
		//if (ch == (string(".")[0])) { return r; }
		//if (ch == ch1) { return r; }
		
		if (ch == rune('/'))  {  return r; }
		if (ch == rune('\\')) {  return r; }
		//if (ch == rune('.'))  { r = string(ch) + r; return r; }
		
		r = string(ch) + r;
	}	

	
	return r;

	
}//

//2022 奇怪，以前应该实现过了的
func ExtractFilePath(s string) (string) {
	
	r := "";
	
	//一定要参考早期的 ExtractFileExt() 函数
	//-----------------

	//	fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
	
	array := []rune(s);
	
	var findPathSP = false;  //是否找到路径分隔符号了
	
    for i:=len(array)-1; i >=0; i--{
		ch := array[i];
		//fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
		

		
		if (false == findPathSP)&&(ch == rune('/'))  {  findPathSP = true; continue; }  //不包括分隔符号本身，所以是 continue
		if (false == findPathSP)&&(ch == rune('\\')) {  findPathSP = true; continue; }
		
		if (false == findPathSP) {  continue; }  //没找到之前也是 continue
		
		r = string(ch) + r;
	}	

	
	return r;

	
}//

//有攻击行为，要清理掉路径参数//参考 mail_AddFJ 上传附件时的攻击行为
func ClearWebParam(s string) (string){
	r := "";
	

	array := []rune(s);
	
    for i:=0; i < len(array); i++{
		ch := array[i];
		//fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
		
		//if (ch == (string(".")[0])) { return r; }
		//if (ch == ch1) { return r; }
		
		if (ch == rune('/'))  {  continue; }
		if (ch == rune('\\')) {  continue; }
		//if (ch == rune('.'))  { r = string(ch) + r; return r; }
		
		r = r + string(ch);
	}	

	
	return r;	
	
}//

func md5_string(s string) (string){
    md5Ctx := md5.New()
	
    //md5Ctx.Write([]byte("test md5 encrypto"))
	md5Ctx.Write([]byte(s))
	
    cipherStr := md5Ctx.Sum(nil)
    fmt.Print(cipherStr)
    fmt.Print("\n")
    fmt.Print(hex.EncodeToString(cipherStr))
	
	return hex.EncodeToString(cipherStr);
}//

//ymd 的日期
func now_ymd() (string){
//	year:=strconv.Itoa(time.Now().Year())
//	month:=time.Now().Month().String();//time.Now().Month().String()
//	day:=strconv.Itoa(time.Now().Day())

	//return c_cn.UTC().Format("2006年01月02日 15点04分05秒"); //这个是秒 //这个不是北京时间
	
	//应该用完整格式，这样才好对 php 接口
	//c_cn.UTC().Format("20060102"); //这个是秒 //这个不是北京时间
	
	r := time.Now().UTC().Format("20060102"); //这个不是北京时间
	
	r = str_replace(r, ".", "_");	

	return r;
}//


//子字符串是否存在//不区分大小写
func FindStr(s, substr string) bool{

	//return strings.Contains("widuu", "wi");
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr));

}//

//模板中的内容不能直接使用原始字符串，如果里面本身有 html 的话会被修改为源码输出
func template_HTML(s string) template.HTML {
	
	return template.HTML(s);
}//

//golang 没有 string 的 SubString,另外这个网友的也是不能处理中文(utf8)的
//https://studygolang.com/articles/3971
//网上一个substring的方法：

//func SubString(str string,begin,length int) (substr string) {
//	// 将字符串的转换成[]rune
//	rs := []rune(str)
//	lth := len(rs)

//	// 简单的越界判断
//	if begin < 0 {
//		begin = 0
//	}
//	if begin >= lth {
//		begin = lth
//	}
//	end := begin + length
//	if end > lth {
//		end = lth
//	}

//	// 返回子串
//	return string(rs[begin:end])
//}

//本人应用的时候发现，多次截取字符串时出现截取失败问题，后来仔细读了一些他的代码，发现有rune，修改了一些rune去掉之后就真正可以截取字符串了，代码如下：
//clq 注意不能处理中文(utf8),实际上目前只能转为 bytes
func SubString_byte(str string, begin, length int) (substr string) {
	lth := len(str)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(str[begin:end])
}//

//截取字符串
//clq 专门针对 golang utf8 的
func SubString(_str string, begin, length int) (string) {
    str := []rune(_str); //utf8 字符的切分体
    lth := len(str);


	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(str[begin:end]);
}//

//太常用了，写一个吧
func strlen(s string) int {
	
	return len(s);
}//


//得到站点的顶级域名//即最后一个 . 前的
func Functions_getRootDomain(str string)string {
	
	s := []rune(str);
    length := len(s);
	
    r := "";
    str_len := length;//strlen($s);
    count := 0;

    for i := str_len-1; i >= 0; i-- {

       if(s[i] == rune('.')) {
          
          count++;
          
          if (count >= 2) { break; }
          r = string(s[i]) + r;
       }else{
          r = string(s[i]) + r;
       }

    }
    
    //echo $array_cn[5];
    return r;
}//



//检查一个 host 是否是域名//这个函数来自 php ，其实有问题，对于 ipv6 是错误的
func Functions_hostIsDomain(str string) bool {
	
	s := []rune(str);
    length := len(s);
	
    //含有 '.' 并含有字符
    //r := false;
    str_len := length;//strlen($s);
    count1 := 0;
    count2 := 0;

    for i := 0; i < str_len; i++ {
       if(s[i] == rune('.')) {
          count1++;
       }

       if((s[i] >= rune('A'))&&(s[i] <= rune('z'))) {//含有字符
          count2++;
       }


    }
    
    if ((count1 > 0)&&(count2 > 0)) {
        return true;
        
    }else{
        return false; 
    }

}//

//截取字符串//网上函数，未测试
func Substr_pos(str string, start int, end int) string {
    rs := []rune(str)
    length := len(rs)

    if start < 0 || start > length {
        return ""
    }

    if end < 0 || end > length {
        return ""
    }
    return string(rs[start:end])
}

//以后再优化,还是弄个同 delphi 的函数名吧
func get_value1(line string, sp1 string, sp2 string) (string){
	
	return GetValue(line, sp1, sp2);
	
}//

func Trim(s string) (string){
	
	return strings.TrimSpace(s); //好象也可以
	//return strings.Trim(s, " ");
	
}//

//字符串是否相等//不区分大小写
func str_eq_i(s1, s2 string) bool{

	return strings.EqualFold(s1, s2);

}//

//字符串替换
func str_replace(s, old, new string,) string {
	s = strings.Replace(s, old, new, -1);
	
	return s;
	
}//

func LowerCase(s string) string {
	
	s = strings.ToLower(s);
	
	return s;	
	
}//

//目录是否存在
func DirectoryExists(fn string) bool {
	//_, err := os.Stat(path)
	_, err := os.Stat(fn)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		fmt.Println("DirectoryExists:", err)
		return false
	}
	return false	
	
}//

//复制文件
func CopyFile(src,dst string) (w int64, err error){
	srcFile,err := os.Open(src);
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close();
	dstFile,err := os.Create(dst); //2022 一定是要记得关闭的
	if err != nil{
		fmt.Println(err);
	}
	defer dstFile.Close(); //2022 一定是要记得关闭的
	return io.Copy(dstFile,srcFile);

}//

//base64 解码
func Base64ToStr(s string) string {
	//PrintError("Base64ToStr");
	
	decodeBytes, err := base64.StdEncoding.DecodeString(s)
    if err != nil {
        //log.Fatalln(err)
		fmt.Println("Base64ToStr:", err);
    }
    //fmt.Println(string(decodeBytes))

    //fmt.Println()	
	return string(decodeBytes);
}//

//base64 编码码
func Base64Encode(s string) string {
	//PrintError("Base64ToStr");
	
	res := base64.StdEncoding.EncodeToString([]byte(s));

	return res;
}//

func Length(s string) int {
	
	return len(s);

}//

//golang 自己内置的文件流保存为附件
//参考 func mail_AddFJ(w http.ResponseWriter, r *http.Request) {	
//按道理，参数中的 file 可以是 os.file 也可以是 multipart.file//参考 Go\src\mime\multipart/formdata.go
func StreamToFile(formFile multipart.File, tofile string) {	


	defer PrintError("StreamToFile");

	//--------------------------------------------------

	////formFile, header, err := r.FormFile("file");

	
	////defer formFile.Close();

//	//var buf = make([]byte, 0); //[]byte(""); //奇怪,不是动态的,要直接指定长度
//	var buf = make([]byte, 10);
//	//var buf = make([]byte, header.Header.); //[]byte(""); //奇怪,不是动态的,要直接指定长度
//	formFile.Read(buf);

//	//Bytes2File(file, fn + file_name + "tmp.txt");
//	Bytes2File(buf, "d:\\tmp.txt");
//	Bytes2File(buf, fn + file_name + "tmp.txt");
//	fmt.Println(fn + file_name + "tmp.txt");
	
	formFile.Seek(0, 0); //前面有读取动作,所以要跳到开始位置
	//---------------
	// 创建保存文件
	//destFile, err := os.Create("./upload/" + header.Filename)
//	destFile, err := os.Create("d:\\tmp2.txt")
	destFile, err := os.Create(tofile)
	if err != nil {
		//log.Printf("Create failed: %s\n", err)
		fmt.Println("Create failed: %s\n", err)
		//return
	}
	defer destFile.Close();
	
	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		//log.Printf("Write file failed: %s\n", err)
		fmt.Println("Write file failed: %s\n", err);
		//return
	}
	
	
}//	


//获取程序运行路径//exe extractfilepath
//func getCurrentDirectory() string {
func GetAppPath_EXE() string {
	
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]));
	
	if err != nil {
		//beego.Debug(err)
		
		return "";
	}
	
	//return strings.Replace(dir, "\\", "/", -1);
	return dir;
	
}//


//用来实现 key,value 的级联,树形
//function HTTPEncode_kv(const AStr: String): String;
func HTTPEncode_kv(s string) string {
//    var urlStr string = "http://baidu.com/index.php/?abc=1_羽毛"
//    l, err := url.ParseQuery(urlStr)
//    fm.Println(l, err)
//    l2, err2 := url.ParseRequestURI(urlStr)
//    fm.Println(l2, err2)

//    l3, err3 := url.Parse(urlStr)
//    fm.Println(l3, err3)
//    fm.Println(l3.Path)
//    fm.Println(l3.RawQuery)
//    fm.Println(l3.Query())
//    fm.Println(l3.Query().Encode())

//    fm.Println(l3.RequestURI())
//    fm.Printf("Hello World! version : %s", rt.Version());
	
	return url.QueryEscape(s);
	
}//

//--------------------------------------------------
//有超时的 http 请求,单位秒
func HttpGet_TimeOut(url string, second time.Duration) ([]byte) {

	defer PrintError("HttpGet_TimeOut");
	
	//var r []byte = nil;

	var c *http.Client = &http.Client{
	
	    Transport: &http.Transport{
			
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //InsecureSkipVerify参数值只能在客户端上设置有效//clq add 让客户端跳过对证书的校验
			
	        Dial: func(netw, addr string) (net.Conn, error) {
	            ////c, err := net.DialTimeout(netw, addr, time.Second*3)
	            c, err := net.DialTimeout(netw, addr, time.Second * second);
	            if err != nil {
	                fmt.Println("HttpGet_TimeOut() dail timeout", err);
	                return nil, err;
	            }
				
				fmt.Println("HttpGet_TimeOut 连接成功 ..."); //test 仍然有卡死的情况
				
				//clq add 似乎可以在这里设置整个通话过程中的时间，超时
				SetConnectTimeOut(c, 10);//test add
				//c.SetDeadline( //其实这样也可以
				
	            return c, nil;
	
	        },
	        MaxIdleConnsPerHost:   10,
	        ////ResponseHeaderTimeout: time.Second * 2,
	        ResponseHeaderTimeout: time.Second * second, //这个应该指的是读取头信息时的超时
			//IdleConnTimeout: time.Second * second, //test 据说 可以控制连接池中一个连接可以idle多长时间
			IdleConnTimeout: time.Second * 60, //test add 据说 可以控制连接池中一个连接可以idle多长时间
	    },
	}
	
	//c.Get(url);
	//resp, err := http.Get(url);
	fmt.Println("HttpGet_TimeOut c.Get(url) ..."); //test 仍然有卡死的情况
	resp, err := c.Get(url);
	fmt.Println("HttpGet_TimeOut c.Get(url)end  ..."); //test 仍然有卡死的情况
	
    if err != nil {
        fmt.Println("error:", err);
        return nil;
    }	
	
	defer resp.Body.Close(); //一定要有
	fmt.Println("HttpGet_TimeOut ioutil.ReadAll ..."); //test 仍然有卡死的情况
	//resp.Header
	fmt.Println(resp.Header.Get("Content-Length")); //test
	fmt.Println(resp.Header.Get("Date")); //test
	fmt.Println(resp.Header); //test
	
	fmt.Println("resp.ContentLength: ", resp.ContentLength); //test //https://www.bitstamp.net/api/ticker/ 没有这个头，也许是这个原因导致后面的 ioutil.ReadAll 无效 
	
	body, err := ioutil.ReadAll(resp.Body); //就是这里卡的
	fmt.Println("HttpGet_TimeOut ioutil.ReadAll end ..."); //test 仍然有卡死的情况
	//fmt.Println(string(body));	
    if err != nil {
        fmt.Println("error:", err);
        return nil;
    }	
			
	//return r;	
	return body;	

}//

//get请求可以直接http.Get方法，非常简单。
func HttpGet(url string) string {
	
	defer PrintError("HttpGet");
	
//    resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
    resp, err := http.Get(url);
	
//    if err != nil {
//        // handle error
//    }
	CheckErr(err);

    defer resp.Body.Close(); //一定要记得关闭
	
    body, err := ioutil.ReadAll(resp.Body);
	
//    if err != nil {
//        // handle error
//    }
	CheckErr(err);
 
    //fmt.Println(string(body));
	
	return string(body);
}//


//2019.04.09 
//https://blog.csdn.net/u010154462/article/details/78412833
//取得一个 dll 的结果字符串
//根据DLL返回的指针，逐个取出字节，到0字节时判断为字符串结尾，返回字节数组转成的字符串
func prttostr2(vcode uintptr, maxlen int) string {
	var vbyte []byte;
	//for i:=0; i<10; i++ { 
	for i:=0; i<maxlen; i++ { 
		sbyte:=*((*byte)(unsafe.Pointer(vcode)));
		
		if sbyte==0{ 
			break;
		} 
		vbyte=append(vbyte,sbyte);
		vcode += 1;
	} 
	return string(vbyte); 
}//


//利用 golang 的断言转换字符串，这个是一定成功的//2020 不一定，还得用 Sprintf
func ToString(s interface{}) string {
	if nil == s { return ""; }
	
	r, ok := s.(string);
	
	//if (false == ok) { return ""; }
	if (false == ok) { r = ""; }
	
	//注意，这时候的 line["CheckMail_s"] 可能是 template.HTML，所以要特别转换
	
	//s = fmt.Sprintf("%v", line["CheckMail_s"]);
	r = fmt.Sprintf("%v", s); //一定要判断是否为 nil ，否则会转换为字符串 "<nil>"
	
	return r;
}//


//从文件中读取字符串
func LoadFromFile_String(fn string) string {
	
	defer PrintError("LoadFromFile_String");
	
    //logger.Infof("get file content as lines: %v", filePath) 
	fmt.Println("LoadFromFile_String() 读取文件:", fn);
	
	var result = "";
	
    b, err := ioutil.ReadFile(fn)  
    if err != nil {  
        //logger.Errorf("read file: %v error: %v", filePath, err)  
		fmt.Println("读取文件失败:", fn, err);
        return result;  
    }  
    s := string(b)  

	
    return s;
	
}//

