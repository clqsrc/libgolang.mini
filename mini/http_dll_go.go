package main;

//因为 http_dll.go 分出来的 golang 语言特性部分，这里不涉及任何的 C 安全操作，全部调用封闭好的函数



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


//输入，输出都是 json ，相关字符串已经在 C 接口函数中转换过了
func GO_run_golang_json(src_json string) (string) {  

	defer PrintError("run_golang_json()");

    //print("Hello ", C.GoString(src_json)); //这里不能用fmt包，会报错，调了很久...
    
    var r = "{}";
    
    var line = DecodeJson([]byte(src_json)).(map[string]interface{}) //(map[string]interface{}) 类型的强制转换语法很奇特

	var func_name = line["func_name"];  //函数名
	var param1 = ToString(line["param1"]);  //参数1
	// var param2 = line["param2"];  //参数1
	// var param3 = line["param3"];  //参数1
	// var param4 = line["param4"];  //参数1
	// var param5 = line["param5"];  //参数1

    
    
    if ("set_on_event" == func_name) {
    	
		var i_param1 = StrToIntDef(param1, 0);
		//var func_handle = C.longlong(i_param1); 
    	
		GO_set_on_event(i_param1);
		
    }
    

	return r;
}//


