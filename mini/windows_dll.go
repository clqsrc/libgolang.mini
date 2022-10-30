package main;

import (
	//"io/ioutil"
//	"database/sql"
	//"fmt"
//	"reflect"
//	"strconv"
	//"strings"
//	"syscall"
	"unsafe"
//	"runtime"
	//"runtime/internal/atomic" //这个是内部包,不能这样引用
	//从 Go\src\syscall\os_windows.go 来看,一个 stdcall 是开了一个线程来调用一个 dll 函数的
	//不对,好象是 tstart_stdcall, newosproc 才开

)

//windows 下的 dll 调用方法

//https://www.cnblogs.com/pu369/p/10350696.html

// func StrPtr(s string) uintptr {
//     return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)));
// }//

//实际上例子中的 StrPtr 是将 utf8 转换成了 utf16, https://github.com/golang/go/commit/9b6e9f0c8c66355c0f0575d808b32f52c8c6d21c
//中有个更直接的方式
// func utf82utf16(s string) uintptr {
//     return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)));
// }//

//取回可以参考 UTF16ToString 源码中的示例
//		p := (*[0xffff]uint16)(unsafe.Pointer(&data.PathBuffer[0]))
//		s = UTF16ToString(p[data.PrintNameOffset/2 : (data.PrintNameLength-data.PrintNameOffset)/2])


//根据上面地址的源码,传递一个原始缓冲字节可以用以下函数转换指针
//来自 func windowsLoadSystemLib(name []byte) uintptr {
func dll_bytes(buf []byte) uintptr {
	//absName := append(sysDirectory[:sysDirectoryLen], name...)
	//return uintptr(unsafe.Pointer(&absName[0]));
	return uintptr(unsafe.Pointer(&buf[0]));
}//

// windows下的第三种DLL方法调用 //linux 下是编译不了的
func ShowMessage3(title, text string) {
	
    // user32, _ := syscall.LoadDLL("user32.dll");
    // MessageBoxW, _ := user32.FindProc("MessageBoxW");
    // MessageBoxW.Call(0, utf82utf16(text), utf82utf16(title), 0);
	

	
}//

//stdcall1
//https://github.com/golang/go/commit/9b6e9f0c8c66355c0f0575d808b32f52c8c6d21c
//处可以看到很清晰的代码
//对于 stdcall 调用模式的,最后会调用到 stdcall() 这个函数上来
//它的参数都是 unsafe.Pointer 所以是要转换的


//据说是高手的代码
//https://github.com/liudch/goci/blob/master/oci.go

//2019.04.09 
//https://blog.csdn.net/u010154462/article/details/78412833
//取得一个 dll 的结果字符串
//根据DLL返回的指针，逐个取出字节，到0字节时判断为字符串结尾，返回字节数组转成的字符串
func prttostr(vcode uintptr, maxlen int) string {
	
	defer PrintError("prttostr"); //这个确实是有问题的
	
	if vcode == 0 { return""; } //clq add 应该判断一下指针
	
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

