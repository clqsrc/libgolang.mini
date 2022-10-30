package main;

//因为 http_dll.go 里有 export 所以不能将纯 C 代码写在那里，会提示函数 multiple definition of `sayHi'
//所以要有一个独立的 c 函数文件，而且还要包装成 go 函数，因为导出的 c 函数只能在本文件中使用

//import "C" 要紧挨着 /*...*/ 注释块
//要加这个才会有导出 C 语言的函数

//#include <stdlib.h>
/*
#include <stdlib.h>
#include <stdio.h>

#ifndef _http_dll_
#define _http_dll_

void (*func_on_event)(const char * event, const char * key, const char * value);

//func_on_event = 0;  //奇怪，这个不行

void sayHi() {

printf("Hi");

}

void set_on_event(long long func)
{
	func_on_event = func;
	
}//

void on_event(const char * event, const char * key, const char * value)
{
	
	func_on_event(event, key, value);
}


#endif


*/
import "C"

import (
	// "fmt"
//	"fmt"
	//"io/ioutil"
//	"database/sql"
	//"fmt"
//	"reflect"
//	"strconv"
	//"strings"
	// "syscall"
	// "unsafe"
	// "net/http"
//	"runtime"
	//"runtime/internal/atomic" //这个是内部包,不能这样引用
	//从 Go\src\syscall\os_windows.go 来看,一个 stdcall 是开了一个线程来调用一个 dll 函数的
	//不对,好象是 tstart_stdcall, newosproc 才开
	
	// "runtime/debug"
	// "runtime"

)

func C_set_on_event(c_func C.longlong) {
	
    // Need a main function to make CGO compile package as C shared library
    
    C.set_on_event(c_func);
    
}//

func C_sayHi() {
	
    // Need a main function to make CGO compile package as C shared library
    
    C.sayHi();
    
}//

func C_on_event(event * C.char, key * C.char, value * C.char) {
	
	C.on_event(event, key, value);
	
}//


func GO_on_event(event string, key string, value string) {
	
	
	//所以调用者要用 free_golang_cstr 来释放这三个参数
	
	cstr_event := C.CString(event);  //这个要用 cfree 释放
	cstr_key   := C.CString(key);  //这个要用 cfree 释放
	cstr_value := C.CString(value);  //这个要用 cfree 释放
	
	//调用函数指针
	C.on_event(cstr_event, cstr_key, cstr_value);
	
	free_golang_cstr(cstr_event);
	free_golang_cstr(cstr_key);
	free_golang_cstr(cstr_value);
	
}//

func GO_set_on_event(func_handle int64) {
	
    var c_func = C.longlong(func_handle);
    
    C.set_on_event(c_func);
    
    //设置成功后可以通知一下调用者 //我们可以看到其实可以通过这个接口让 golang 反过来调用调用者的接口功能
    GO_on_event("event", "key", "value");  //test 2022.10
    
}//
