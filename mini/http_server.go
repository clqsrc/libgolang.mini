// web 服务器
//from http://www.cnblogs.com/yjf512/archive/2012/09/03/2668384.html
//from https://github.com/jianfengye/webdemo
package main



import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"strings"
	"reflect"
	"runtime"
	"time"
	"mime"
	//"runtime/debug"

	
)

func main1() {
	
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"));return;
	//fmt.Println(time.Now().Format("2006-01-02_15:04:05_ .999999999 -0700 MST"));return;
	
	////fmt.Println(time.Now().Format("20060102_150405_.000.999999999"));return;


	LoadAppConfig(); //读取本程序的配置

	//Create_PrivateKey_dkim_email(); return; //生成密钥,运行一次就行了
	
	if (UConfig_IsMainServer() == 1) { fmt.Println("这是主服务器程序[UConfig_IsMainServer]"); 
								}else{ fmt.Println("这是单服务器版本[one server]"); }
								
							

								
	//主服务器包括个单服务器的功能,可以响应对方的备份请求,也可以去备份他人的数据,单服务器只有操作自己的数据
	//--------------------------------------------------							
		
	//var firstName, lastName string;
	//fmt.Scanln(&firstName, &lastName); //似乎没用,内存占用就是这么高
	    
	// 限制为CPU的数量减一//据说这个对性能影响很大
    runtime.GOMAXPROCS( runtime.NumCPU() - 1 );
	
	fmt.Println("Hello World!")
	fmt.Println("123中文")
	
	http.Handle("/css/", http.FileServer(http.Dir("template")));
    http.Handle("/js/", http.FileServer(http.Dir("template")));

	//http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html")))); //只有这样才能正确的访问 http://127.0.0.1:8888/html/mail_vip.html
	//第二个参数其实是根目录,这里指的是程序在在目录即为网站的根目录
	mime.AddExtensionType(".apk", "application/vnd.android");
	mime.AddExtensionType(".js", "text/javascript");
	mime.AddExtensionType(".css", "text/css");
	http.Handle("/html/", http.FileServer(http.Dir(""))); //只有这样才能正确的访问 http://127.0.0.1:8888/html/mail_vip.html
	//这样是不能正确的访问 http://127.0.0.1:8888/html/mail_vip.html 的,大概后者指的是根目录的位置
	//http.Handle("/html/", http.FileServer(http.Dir("html"))); 

    //http.Handle("/mail/files/", http.FileServer(http.Dir("mail/files/")))
     
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/login/",loginHandler)
    http.HandleFunc("/ajax/",ajaxHandler)
    http.HandleFunc("/",NotFoundHandler)
	//http.HandleFunc("/fast/",fastHandler) //clq 原有的用反射,并不好,这个是新示例
	http.HandleFunc("/fast_db/",fastDbHandler) //clq 原有的用反射,并不好,这个是新示例//数据库备份接口
	http.HandleFunc("/sql/",fastDbHandler) //clq 原有的用反射,并不好,这个是新示例//数据库备份接口
	http.HandleFunc("/sql.php",fastDbHandler) //clq 模拟一个文件,这样就可以了

//    err := http.ListenAndServe(":8888", nil)
//    go http.ListenAndServe(":8888", nil); //放线程中,要不会一直阻塞,开启不了别的服务

//	f := func(){
//			if (err != nil) {
//				println("err 监听 8888 端口错误:", err);
//			}
//		}

	http_port := gConfig["http_port"];
	
	LoadHttpDlls();
	//ExecuteHttp(nil, nil);
	////加了守护进程就不要再弹阻塞消息了//ShowMessage3("", "监听端口:" + http_port);
	
	

	go func(){
		//err := http.ListenAndServe(":8888", nil)
		fmt.Println("http_port:" + http_port);
		err := http.ListenAndServe(":" + http_port, nil);
		if (err != nil) {
			println("err 监听端口[" + http_port + "]错误:", err);
		}
	}()
	
	
	//--------------------------------------------------
	go func(){
		//server_log(); //服务器运行日报
	}();

	//--------------------------------------------------
	
	//go DoDbBak(); //自动备份
	
	fmt.Println("exit?")
	
	var i int = 0;
	for{
		time.Sleep(1 * time.Second); //奇特的语法
		//fmt.Scanln(); //这个在 linux 下会出问题
		
		i++;
		//fmt.Println(i);
	}
	
	fmt.Println("exit!")
	
}//



func adminHandler(w http.ResponseWriter, r *http.Request) {
    // 获取cookie
    cookie, err := r.Cookie("admin_name")
    if err != nil || cookie.Value == ""{
        http.Redirect(w, r, "/login/index", http.StatusFound)
    }
    
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }
    
    admin := &adminController{}
    controller := reflect.ValueOf(admin)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    userValue := reflect.ValueOf(cookie.Value)
    method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    ajax := &ajaxController{}
    controller := reflect.ValueOf(ajax)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("loginHandler")
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    login := &loginController{}
    controller := reflect.ValueOf(login)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("NotFoundHandler");
	fmt.Println(r.RequestURI);
	
    if r.URL.Path == "/" {
        //clq 2019//http.Redirect(w, r, "/sql/index", http.StatusFound); //这个是跳转的示例用法
    }
	
	//--------------------------------------------------
	//Ajax跨域问题的两种解决方法之一,据说 html5 后的才支持
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Content-Type", "text/html; charset=utf-8") //要改这个的话,前面的 w 都不能有输出
	w.Write([]byte("404 page (NotFoundHandler)"));
	w.Write([]byte("服务器访问成功.<br>此为无数据处理的默认页面.如外网机器无法查看到本页面,请确认本服务器防火墙监听端口 " + gConfig["http_port"] + " 是否已打开. "));
	return;
	//--------------------------------------------------
    
    //t, err := template.ParseFiles("template/html/404.html")
    t, err := template.ParseFiles("404.html") //奇怪,没有这个文件的情况下服务器会关闭这个连接,连空白都不返回
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}//

