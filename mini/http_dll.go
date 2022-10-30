package main;

//因为这里有 export 所以不能将纯 C 代码写在这里，会提示函数 multiple definition of `sayHi'

//import "C" 要紧挨着 /*...*/ 注释块
//要加这个才会有导出 C 语言的函数

//#include <stdlib.h>
/*
#include <stdlib.h>
#include <stdio.h>

#ifndef _http_dll_
#define _http_dll_

void (*func_on_event)(const char * event, const char * key, const char * value);

// void sayHi() {

// printf("Hi");

// }

// void set_on_event(int func)
// {
	
	
// }//

// void on_event(const char * event, const char * key, const char * value)
// {
	
	
// }


#endif


*/
import "C"

import (
	"fmt"
	// "fmt"
//	"fmt"
	//"io/ioutil"
//	"database/sql"
	//"fmt"
//	"reflect"
//	"strconv"
	//"strings"
	// "syscall"
	"unsafe"
	// "net/http"
//	"runtime"
	//"runtime/internal/atomic" //这个是内部包,不能这样引用
	//从 Go\src\syscall\os_windows.go 来看,一个 stdcall 是开了一个线程来调用一个 dll 函数的
	//不对,好象是 tstart_stdcall, newosproc 才开
	
	// "runtime/debug"
	// "runtime"

)

//据说是高手的代码
//https://github.com/liudch/goci/blob/master/oci.go
//2019 更新,这份代码的字符串传递也是有问题的,例如 
/*
func OCILogon(env *OCIHandle, username, password, database string) (svcctx *OCIHandle, err error) {
	var s *OCIHandle = new(OCIHandle) // Service context
	s.t = OCI_HTYPE_SVCCTX
	var e *OCIHandle = new(OCIHandle) // Error handle
	e.t = OCI_HTYPE_ERROR

	r0, _, e1 := procOCILogon.Call( // sword OCILogon (
		env.h, // OCIEnv          *envhp,
		uintptr(unsafe.Pointer(&e.h)),                               // OCIError        *errhp,
		uintptr(unsafe.Pointer(&s.h)),                               // OCISvcCtx       **svchp,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(username))), // CONST OraText   *username,
		uintptr(len(username)),                                      // ub4             uname_len,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(password))), // CONST OraText   *password,
		uintptr(len(password)),                                      // ub4             passwd_len,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(database))), // CONST OraText   *dbname,
		uintptr(len(database)))                                      // ub4             dbname_len );

	if r0 != 0 {
		err = error(e1)
		return
	}

	return s, nil
}
这其中的用户名就是有可能在大压力环境下失效
*/


//其实只要导出一个函数就可以了

var (
	// http_func * syscall.LazyProc = nil; //syscall.NewLazyDLL("C:\\Program Files\\Oracle\\instantclient_11_2\\oci.dll")
	// http_func_sql * syscall.LazyProc = nil;
	// http_func_sql_err1 * syscall.LazyProc = nil;
	// free_buf * syscall.LazyProc = nil;
)

func LoadHttpDlls(){

	//mod_http := syscall.NewLazyDLL("http.dll");
	// mod_http := syscall.NewLazyDLL("http_dll.dll")

	// http_func = mod_http.NewProc("http_func");        // Clear all attribute-value information in a namespace of an application context
	// http_func_sql = mod_http.NewProc("http_func_sql");
	// http_func_sql_err1 = mod_http.NewProc("http_func_sql_err1");
	// free_buf = mod_http.NewProc("free_buf");

}//



//重要!!! 调用 dll 传递字符串参数因 gc 导致数据失效的示例(实际上会因此内存访问错误崩溃)
func dll_errtest1(sql string) (string, error) {
	
	defer PrintError("dll_errtest1");
	
	/*
		
		//------------------
		//调试 delphi 的 dll 时意外发现 syscall.StringToUTF16Ptr 已经是弃用的了 go 1.7.3 的源码注释中说了用 UTF16PtrFromString 来代替
		sqlp, err := syscall.UTF16PtrFromString(sql);
		if err !=nil { panic("dll 调用参数严重错误!!!");}
		
		a, err2 := syscall.UTF16FromString(sql); //来自 UTF16PtrFromString 的源码
		if err2 !=nil { panic("dll 调用参数严重错误!!! err2");}
		
		sqlp = &a[0];
		
		//------------------
		p1 := uintptr(unsafe.Pointer(sqlp));
		
		runtime.GC(); //不调用这两个应该也是很容易重现的//还是加上容易重现,并且是一大堆一起失效,而没的的话则是久不久失效一个指针
		debug.FreeOSMemory(); 
		
		//放到一个单独的函数时,这里调用 gc 后面会立即崩溃! 不 gc 的时候还能运行一下
		
		r0, _, e1 := http_func_sql.Call(
			//utf82utf16(r.URL.Path),//,
			p1, //uintptr(unsafe.Pointer(sqlp)), // utf82utf16(sql),//, !!! 这个会出错,内存混乱
			//uintptr(unsafe.Pointer(http_func)), // OCIEnv        **envhpp,
			//uintptr(0),                 // ub4           mode,
			//0,                             // CONST dvoid   *ctxp,
	
			);
	
		//if r0 != 0 {
		//	return "", error(e1)
		//}//if
		
		if e1 != nil {
		//	return "", error(e1)
		}//if	
		
		////s := prttostr(r0, 4*1024*1024);	//这个确实是有问题的
		fmt.Println(r0);
		
		//fmt.Println(p1); //强制再引用一会,不让垃圾回收//这个变量没用
		//fmt.Println(a); //强制再引用一会,不让垃圾回收//这个变量有用
		////fmt.Println(sqlp); //强制再引用一会,不让垃圾回收//这个变量有用,这个 *uint16 变量居然也可以让其所在的 []uint16 数组保持引用!!! 也就是说即使是引用了指针,其数组也是会保持的,那么会解引用的原因应该只是那个 unsafe.Pointer 后的变量
		//也就是说 unsafe.Pointer 后的变量就随时可能被释放掉!
		
		
		//fmt.Println(a); //奇怪这里不调用的话, a 的内存在多线程中会发生变化,估计是调用时间拖动得长的话被垃圾回收了,所以一定要在 dll 调用后再使用一下其中的变量,
		//特别是字符串变量一定要注意,不能是只使用它的指针,要整个字节缓冲区一起//过会我写个可以一定重现的 dll 函数调用
			
		//参考 zsyscall_windows.go 的用法,也是一直引用一个内存区的	
			
			
			
		*/
		//return s, nil;
		return "", nil;

		

}//



//测试不保留引用的情况下,看看只传递非 unsafe.Pointer() 的指针,是否能保持内存的存在//感觉 go1.7.3 就是这样传递字符串参数的
//如果成立说明对非 unsafe.Pointer 指针的引用可以保持到传递出去的函数调用返回时
func dll_errtest3(sql string) (string, error) {
	
	defer PrintError("dll_errtest3");
	/*
	
		sqlp, err := syscall.UTF16PtrFromString(sql);
		if err !=nil { panic("dll 调用参数严重错误!!!");}

		//------------------
		//p1 := uintptr(unsafe.Pointer(sqlp));
		
		runtime.GC(); //不调用这两个应该也是很容易重现的//还是加上容易重现,并且是一大堆一起失效,而没的的话则是久不久失效一个指针
		debug.FreeOSMemory(); 
		
		//放到一个单独的函数时,这里调用 gc 后面会立即崩溃! 不 gc 的时候还能运行一下
//		dll_errtest3_sub(sqlp);
		
		//fmt.Println(sqlp); //重要! 可以看到,将非 unsafe.Pointer 指针传递出去后就可以不用再人为的保持内存块引用了
		*/
		return "", nil;
}//

//--------------------------------------------------
//所以综上所述,为了更好的保持 golang 中的指针指向最好是按照 go 本身调用 dll 的做法再封装一层同名函数(加下划线在前面)

//参考 zsyscall_windows.go
//func DnsQuery(name string, qtype uint16, options uint32, extra *byte, qrs **DNSRecord, pr *byte) (status error) {
//	var _p0 *uint16
//	_p0, status = UTF16PtrFromString(name)
//	if status != nil {
//		return
//	}
//	return _DnsQuery(_p0, qtype, options, extra, qrs, pr)
//}


//这里不传递字符串参数,因此不用再来一个包装的 dll 函数了
func dll_free_buf(pbuf uintptr)  {
	
	/*
		//返回值中,第一个就是函数的返回值,第三个是调用中的错误,但第二个一直没看到官方说明
			r0, _, e1 := free_buf.Call(
			//utf82utf16(r.URL.Path),//,
			pbuf, // utf82utf16(sql),//, !!! 这个会出错,内存混乱
			//uintptr(unsafe.Pointer(http_func)), // OCIEnv        **envhpp,
			//uintptr(0),                 // ub4           mode,
			//0,                             // CONST dvoid   *ctxp,
	
			);
	
		//if r0 != 0 {
		//	return "", error(e1)
		//}//if
		
		if e1 != nil {
		//	return "", error(e1)
		}//if	
		
		fmt.Println("dll_free_buf r0:", r0); //这个 dll 函数是没有返回值的,那么这里会得到什么呢
		//参考 golang 自己的 DnsRecordListFree 实现,就是根本不管返回值,虽然最后的 call 调用是有返回值的
		//不过这里我们不能保证自己的 dll 一定是正确的调用方法,所以还是检查返回值的好
		*/
	
}//


//----------------------------------------------------------------

//https://www.qycn.com/xzx/article/10494.html

//go build -buildmode=c-shared -o exportgo.dll exportgo.go
//go build -buildmode=c-shared -o libgolang.dll http_dll.go
//linux 下
//go build -buildmode=c-shared -o libgolang.so http_dll.go
//----
//可以指定几个文件名，也可以不指定使用几个文件
//go build -buildmode=c-shared -o libgolang.so http_dll.go check_error.go
//不过这其中一定要包含一个 main() 函数
//----
//要加这个才会有导出 C 语言的函数
//import "C"
//----
//linux 下这样查看是否成功导出了函数
//nm -D libgolang.so

//一定要加 export 在函数名前，才会导出这个函数
//export Sum
func Sum(a int, b int) int {
    return a + b;
}

//export run_golang_json
func run_golang_json(src_json *C.char) (*C.char) {  

	defer PrintError("run_golang_json()");

    print("Hello ", C.GoString(src_json)); //这里不能用fmt包，会报错，调了很久...
    
    
    //参考 https://segmentfault.com/q/1010000040761516/
    //C.CString 的结果是要用 stdlib.h 中的 C.free 来释放的
    // Go string to C string
	// The C string is allocated in the C heap using malloc.
	// It is the caller's responsibility to arrange for it to be
	// freed, such as by calling C.free (be sure to include stdlib.h
	// if C.free is needed).
	//func C.CString(string) *C.char
	
	//C.func_on_event = nil;
	
	//--------------------------------------------------------
	var json = C.GoString(src_json);
	
	var res = GO_run_golang_json(json);
	
	return C.CString(res);
	
	//--------------------------------------------------------
	

    return C.CString("golang aaa");  //这个要用 cfree 释放

}//import "C" 要紧挨着 /*...*/ 注释块


//释放 golang 自己产生的 c 风格字符串

//export free_golang_cstr
func free_golang_cstr(golang_c_str *C.char)  {  

	defer PrintError("free_golang_cstr()");

    print("free_golang_cstr() ok. \r\n"); //这里不能用fmt包，会报错，调了很久...
    
    
    //c.free();
    
    //C.free(golang_c_str);  //这个要用 cfree 释放
    C.free(unsafe.Pointer(golang_c_str))

}//




//----
//golang 要求 dll 中也要有个 main 函数
func main() {
	
    // Need a main function to make CGO compile package as C shared library
    
    //C.sayHi();  //不能直接调用另外的 C 函数
    C_sayHi();
    
    //这里只是功能示例，实际上 C.CString 的返回值都是要手动释放的
    
    var i_client = create_tcp_connect();
    
    //tcp_connect(i_client, "163.com", 80, 0);  //普通连接 ok
    tcp_connect(i_client, C.CString("163.com"), 443, 1); //https 也成功了
    //tcp_connect(i_client, C.CString("163.com"), 4433, 1);
    //tcp_connect(i_client, "163.com", 4433, 1);  //奇怪，端口不存在的话就崩溃了
    
    var client * DLL_TcpClient = _to_tcp_connect(i_client);
    
    fmt.Println("tcp_send_string()");
    
    
    //SendLine(client.Read, client.Write, "GET / HTTP 1.0\r\n\r\n");
    tcp_send_string(i_client, C.CString("GET / HTTP 1.0\r\n\r\n"));
    
    rs2 := tcp_recv_line(i_client);
    fmt.Println("rs2:", C.GoString(rs2));
    free_golang_cstr(rs2);  //调用 C 接口后一定要释放
    
    rs2 = tcp_recv_string(i_client, C.CString("\n"));
    fmt.Println("rs2:", C.GoString(rs2));
    free_golang_cstr(rs2);  //调用 C 接口后一定要释放
    
    rs := _RecvLine(client.Read, client.Write); //只收取一行
    fmt.Println("rs:", rs);
    rs = _RecvLine(client.Read, client.Write); //只收取一行
    fmt.Println("rs:", rs);
    rs = _RecvLine(client.Read, client.Write); //只收取一行
    fmt.Println("rs:", rs);
    rs = _RecvLine(client.Read, client.Write); //只收取一行
    fmt.Println("rs:", rs);
    
    
    
    free_tcp_connect(i_client);
    
}//



