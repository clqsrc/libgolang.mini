package main;

//http 并不能完成全部的工作，所以还是要有 tcp
//之所以用 golang 不用 qt 原生或者操作系统原生 socket ，是因为要连接安全套接字
//其实实现并不难，https://github.com/clqsrc/newbt_go_smtp/blob/master/smtp_client.go 里已经实现了大部分

import "C"

import (
	"time"
    "fmt"
    "bufio"
//    "crypto/tls"
    // "encoding/base64"
//    "errors"
//    "io"
    "net"
    "crypto/tls"
//    "net/smtp" //clq add
//    "net/textproto"
    // "strings"
    "strconv"
    // "time"
    "sync"
    "container/list"
)

//在 socket_functions.go 中已经有了

// var gConn net.Conn;
// var gRead * bufio.Reader;
// var gWrite * bufio.Writer;

//可以放到这样的类里
// type TcpClient struct {
//     Conn net.Conn;
//     Read * bufio.Reader;
//     Write * bufio.Writer;
// }//

type DLL_TcpClient struct {
    Conn net.Conn;
    Read * bufio.Reader;
    Write * bufio.Writer;
    
    use_ssl int;  //是否使用安全套接字
    
}//

//自动管理内存的语言都可以用这种方式，比如 java

//--------------------------------------------------------
var gDLL_Obj map[int]interface{} = make(map[int]interface{});  //存放 dll 生成的所有对象实例，否则它们会被 cg 自动回收
var gDLL_Obj_count = 0;  //个数

var gDLL_Obj_delete list.List;  //被删除的索引列表，还要重用的

var gDLL_Obj_lock *sync.Mutex = new(sync.Mutex) //互斥锁Mutex

//--------------------------------------------------------

//保存一个对象的 引用
func DLL_Add_Object(obj interface{}) int { 

	defer gDLL_Obj_lock.Unlock(); //一定要有

	gDLL_Obj_lock.Lock(); //一定要有
	
	//--------------------------------------------------------
	
	var new_id = true;
	
	var _id int = 0;
	
	//----
	//是使用新 id ，还是用删除列表中回收的
	if (gDLL_Obj_delete.Len() > 0) {
		
		new_id = false;  //有回收的就不会生成新 id
		
		//----
		
		//var old = gDLL_Obj_delete.PushFront();
		//fmt.Println(l.Back().Value)  //输出尾部元素的值,4
		var old = gDLL_Obj_delete.Back().Value;  //输出尾部元素的值
		gDLL_Obj_delete.Remove(gDLL_Obj_delete.Back()); //删除尾部元素
		
		//_old, ok := old.Value.(int) // Alt. non panicking version 
		_old, ok := old.(int) // Alt. non panicking version 
		
		if (ok) { 
		
			_id = _old; 
			
		}//if 2
		
		
	}//if 1
	
	
	//----
	
	//生成新 id
	if (true == new_id) {
		gDLL_Obj_count = gDLL_Obj_count + 1;
		
		_id = gDLL_Obj_count;
	}
	
	//gDLL_Obj[gDLL_Obj_count] = obj;
	gDLL_Obj[_id] = obj;
	
	
	return _id;
	
}//

//删除一个对象的 引用 //参数是这个索引对应的
func DLL_Delete_Object(obj_dll_index int) { 

	defer gDLL_Obj_lock.Unlock(); //一定要有

	gDLL_Obj_lock.Lock(); //一定要有
	
	//--------------------------------------------------------
	
	gDLL_Obj_delete.PushBack(obj_dll_index);
	
	gDLL_Obj[obj_dll_index] = nil;
	
	//gDLL_Obj.re
	
	//golang 中 map 删除元素的方法很奇特
	//delete(m, k)
	delete(gDLL_Obj, obj_dll_index);  //不过所说这样也不会释放内存，至少目前的版本来说，不过至少官方删除代码如此
	
}//



//创建一个 tcp 连接
//export create_tcp_connect
func create_tcp_connect() int {
	
	
	var client = new(DLL_TcpClient);
	
	var i_client int = DLL_Add_Object(client);
	
	return i_client;
	
}//

//export free_tcp_connect
func free_tcp_connect(i_client int) {
	
	fmt.Println("free_tcp_connect() :", i_client);
	
	
	DLL_Delete_Object(i_client);
	
}//

//dll 的索引转换出一个 tcp 对象
func _to_tcp_connect(i_client int) (*DLL_TcpClient) {
	
	defer PrintError("");
	
	
	var obj = gDLL_Obj[i_client];
	
	var client * DLL_TcpClient = nil;
	
	_client, ok := obj.(* DLL_TcpClient); // Alt. non panicking version 
	
	if (ok) { 
	
		client = _client; 
		
	}//if 2
	
	
	return client;
	
}//


//连接一个地址，并可选择是否使用 ssl/tls
func GO_tcp_connect(i_client int, host string, port int, use_ssl int) int {
	
	defer PrintError("GO_tcp_connect");

	var r int = 0;	
	
	var client * DLL_TcpClient = _to_tcp_connect(i_client);
	
	if (nil == client) { return 0; }
	
	client.use_ssl = use_ssl;   //要根据这个标识来判断调用不同的超时设置函数，因为 golang 的超时比较奇特，每次读取/写入后都要设置
	
	if (0 == use_ssl) {
		
		client.Conn, client.Read, client.Write = ConnectTimeOut(host, port, 20);
		
	}else {
		
		client.Conn, client.Read, client.Write = Connect_TimeOut_TCP_SSL(host, port, 20);
		
	}
	
	r = 1; //成功
	return r;
}//


//export tcp_connect
func tcp_connect(i_client int, host *C.char, port int, use_ssl int) int {
	
	return GO_tcp_connect(i_client, C.GoString(host), port, use_ssl);


}//


//发送，有超时
func tcp_send_line(i_client int, line string){
	
	var client * DLL_TcpClient = _to_tcp_connect(i_client);
	
	if (nil == client) { return; }
    
    if (0 == client.use_ssl) {
		
		SetConnectTimeOut(client.Conn, 20);
		
	}else {
		
		//SetConnectTimeOut_ssl(client.Conn, 20);
		SetConnectTimeOut_ssl(client.Conn.(*tls.Conn), 20);
		
	}
    
    //SendLine(client.Read, client.Write, "GET / HTTP 1.0\r\n\r\n");
    SendLine(client.Read, client.Write, line);
}//

func GO_tcp_send_string(i_client int, line string){
	
	defer PrintError("GO_tcp_send_string");
	
	var client * DLL_TcpClient = _to_tcp_connect(i_client);
	
	if (nil == client) { return; }
    
    if (0 == client.use_ssl) {
		
		SetConnectTimeOut(client.Conn, 20);
		
	}else {
		
		//SetConnectTimeOut_ssl(client.Conn, 20);
		SetConnectTimeOut_ssl(client.Conn.(*tls.Conn), 20);
		
	}
    
    //SendLine(client.Read, client.Write, "GET / HTTP 1.0\r\n\r\n");
    //SendLine(client.Read, client.Write, line);
    SendString(client.Read, client.Write, line);
}//

//export tcp_send_string
func tcp_send_string(i_client int, line *C.char){
	
	GO_tcp_send_string(i_client, C.GoString(line));
}//

func GO_tcp_recv_Line(i_client int) string {
	
	var client * DLL_TcpClient = _to_tcp_connect(i_client);
	
	if (nil == client) { return ""; }
    
    if (0 == client.use_ssl) {
		
		SetConnectTimeOut(client.Conn, 20);
		
	}else {
		
		//SetConnectTimeOut_ssl(client.Conn, 20);
		SetConnectTimeOut_ssl(client.Conn.(*tls.Conn), 20);
		
	}
	
	return _RecvLine(client.Read, client.Write); //只收取一行
	
}//

//export tcp_recv_line
func tcp_recv_line(i_client int) (*C.char) {
	
	defer PrintError("tcp_recv_line");
	
	var client * DLL_TcpClient = _to_tcp_connect(i_client);
	
	if (nil == client) { return C.CString(""); }
    
    if (0 == client.use_ssl) {
		
		SetConnectTimeOut(client.Conn, 20);
		
	}else {
		
		//SetConnectTimeOut_ssl(client.Conn, 20);
		SetConnectTimeOut_ssl(client.Conn.(*tls.Conn), 20);
		
	}
	
	s := _RecvLine(client.Read, client.Write); //只收取一行
	
	return C.CString(s);
	
}//

//使用终止分隔符的

//export tcp_recv_string
func tcp_recv_string(i_client int, sp *C.char) (*C.char) {
	
	defer PrintError("tcp_recv_string");
	
	var client * DLL_TcpClient = _to_tcp_connect(i_client);
	
	if (nil == client) { return C.CString(""); }
    
    if (0 == client.use_ssl) {
		
		SetConnectTimeOut(client.Conn, 20);
		
	}else {
		
		//SetConnectTimeOut_ssl(client.Conn, 20);
		SetConnectTimeOut_ssl(client.Conn.(*tls.Conn), 20);
		
	}
	
	//s := _RecvLine(client.Read, client.Write); //只收取一行
	s := _RecvString(client.Read, client.Write, C.GoString(sp)); //只收取一行
	
	return C.CString(s);
	
}//


//--------------------------------------------------------
//

//export http_get
func http_get(url *C.char) (*C.char) {
	
	defer PrintError("http_get");
	
	var buf = HttpGet_TimeOut(C.GoString(url), 20);
	
	var s = string(buf);
	
	return C.CString(s);
}//

//export http_get_to_file
func http_get_to_file(url *C.char, fn *C.char) {
	
	defer PrintError("http_get_to_file");
	
	var buf = HttpGet_TimeOut(C.GoString(url), 20);
	
	//return string(buf);
	
	Bytes2File(buf, C.GoString(fn));
	
}//


//--------------------------------------------------------

//从别的文件复制过来的函数


//--------------------------------------------------------


//ssl/tls 有带超时的连接
//带超时的连接，单位是秒//TCP 的
func Connect_TimeOut_TCP_SSL(host string, port int, timeout_second int) (net.Conn, * bufio.Reader, * bufio.Writer) {
    
    addr := host + ":" + strconv.Itoa(port);
    
    //--------------------------------------------------------
    //不校验的默认 ssl 系统证书
    
    tlsConfig := &tls.Config{
        //Certificates: []tls.Certificate{certificate},
        InsecureSkipVerify: true, //clq add  //不校验证书
    }	
    
    //--------------------------------------------------------
    
    //奇怪，tls 怎么处理超时 //要用 return DialWithDialer(new(net.Dialer), network, addr, config)
    //conn, err := tls.Dial("tcp", addr, tlsConfig);
    
    //--------
    //dialer *net.Dialer
    var dialer = new(net.Dialer);
    
    dialer.Timeout = time.Second * time.Duration(timeout_second); //2;  //这个的单位应该是 Nanosecond 纳秒 ，因为它的类型是 time.Duration ，而 time.Duration 的 1 个单位就是 1 纳秒
    
    conn, err := tls.DialWithDialer(dialer, "tcp", addr, tlsConfig);
    
    //--------
    
    SetConnectTimeOut_ssl(conn, 20);
    //SetConnectTimeOut_ssl(conn, timeout_second);
    

	//conn, err := tls.DialTimeout("tcp", addr, time.Duration(timeout_second) * time.Second); //超时单位为秒，原始函数有点难理解
    if err != nil {
        return nil, nil, nil
    }
	
	var conn_net net.Conn;
	var conn_tcp net.TCPConn;
	var conn_tcp_p * net.TCPConn;
	//reader := bufio.NewReader(conn_tcp); //这个语法很奇怪//err 但对于 net.Conn 可以
	reader := bufio.NewReader(&conn_tcp); //这个语法很奇怪//ok
	reader = bufio.NewReader(conn_net); //ok
	
	conn_net = &conn_tcp; //可以这样转换//ok 所以 net.Conn 是指针 net.TCPConn 是结构体（暂时理解）
	
	//*conn_tcp = *conn_net;//conn_net;
	//conn_tcp_p = (* net.TCPConn)(conn_net);//err golang 不支持这种强制转换写法
	conn_tcp_p = conn_net.(*net.TCPConn); //可以这样，相当于 C 语言的 conn_tcp_p = (*net.TCPConn)conn_net; //在 golang 中这叫断言,没有的话会提示 need type assertion
	
	//按道理应该是可以的，因为小写所以不能在其他包使用//cc := &net.TCPConn{c};
	//newTCPConn(fd);
	//net.DialTCP()
	
	//rect3 := &net.TCPConn{conn};
//	c = net.TCPConn(conn);

//	tcpconn := &net.TCPConn{conn};

    
    reader = bufio.NewReader(conn);
    writer := bufio.NewWriter(conn);
    
    //writer.WriteString("EHLO\r\n");
    //writer.Flush();
	
	fmt.Println("conn_tcp_p:", conn_tcp_p);	
    
    //host, _, _ := net.SplitHostPort(addr)
    //return NewClient(conn, host)    
    return conn, reader, writer;
}//


// func Connect(host string, port int) (net.Conn, * bufio.Reader, * bufio.Writer) {
    
//     addr := host + ":" + strconv.Itoa(port);
//     conn, err := net.Dial("tcp", addr);
//     if err != nil {
//         return nil, nil, nil
//     }
    
//     reader := bufio.NewReader(conn);
//     writer := bufio.NewWriter(conn);
    
//     //writer.WriteString("EHLO\r\n");
//     //writer.Flush();
    
//     //host, _, _ := net.SplitHostPort(addr)
//     //return NewClient(conn, host)    
//     return conn, reader, writer;
// }//

//收取一行,可再优化 
//func RecvLine(conn *net.Conn) (string) {
//func RecvLine(conn net.Conn, reader * bufio.Reader) (string) {
// func _RecvLine() (string) {
    
//     //defer conn.Close();
//     ////reader := bufio.NewReader(conn);
//     //reader := bufio.NewReaderSize(conn,409600)
    
//     //line, err := reader.ReadString('\n'); //如何设定超时?
//     line, err := gRead.ReadString('\n'); //如何设定超时?
    
//     if err != nil { return ""; }
    
//     line = strings.Split(line, "\r")[0]; //还要再去掉 "\r"，其实不去掉也可以
    
//     return line;
// }//

// func SendLine(line string){
//     gWrite.WriteString(line + "\r\n");
//     gWrite.Flush();
// }//

//解码一行命令,这里比较简单就是按空格进行分隔就行了
// func DecodeCmd(line string, sp string) ([]string){

//     //String[] tmp = line.split(sp); //用空格分开//“.”和“|”都是转义字符，必须得加"\\";//不一定是空格也有可能是其他的
//     //String[] cmds = {"", "", "", "", ""}; //先定义多几个，以面后面使用时产生异常
    
//     tmp := strings.Split(line, sp);
//     //var cmds = [5]string{"", "", "", "", ""}; //先定义多几个，以面后面使用时产生异常
//     var cmds = []string{"", "", "", "", ""}; //先定义多几个，以面后面使用时产生异常
//     //i:=0;
//     for i:=0;i<len(tmp);i++ {
//         if i >= len(cmds) { break;}
//         cmds[i] = tmp[i];
//     }
//     return []string(cmds);
// }//

//读取多行结果
//func RecvMCmd() (string)

//简单的测试一下 smtp
/*
func test_smtp() {
    
    //连接
    //gConn, gRead, gWrite = Connect("newbt.net", 25);
    //gConn, gRead, gWrite = Connect("newbt.net", 25);
    gConn, gRead, gWrite = Connect("smtp.163.com", 25);
    
    //收取一行
    line := _RecvLine();
    fmt.Println("recv:" + line);
    
    //解码一下,这样后面的 EHLO 才能有正确的第二个参数
    cmds := DecodeCmd(line, " ");
    domain := cmds[1]; //要从对方的应答中取出域名//空格分开的各个命令参数中的第二个
    
    //发送一个命令
    //SendLine("EHLO"); //163 这样是不行的,一定要有 domain
    SendLine("EHLO" + " " + domain); //domain 要求其实来自 HELO 命令//HELO <SP> <domain> <CRLF>    
    
    //收取多行
    //line = _RecvLine();
    line = RecvMCmd();
    fmt.Println("recv:" + line);
    
    //--------------------------------------------------
    //用 base64 登录 
    SendLine("AUTH LOGIN");    
    
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line);
    
    //s :="test1@newbt.net"; //要换成你的用户名,注意 163 邮箱的话不要带后面的 @域名 部分 
    s :="clq_test"; //要换成你的用户名,注意 163 邮箱的话不要带后面的 @域名 部分 
    s = base64.StdEncoding.EncodeToString([]byte(s));
    //s = base64_encode(s);
    SendLine(s);    
    
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line);
    

    s = "123456"; //要换成您的密码 
    //s = base64_encode(s);
    s = base64.StdEncoding.EncodeToString([]byte(s));
    SendLine(s);    
    
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line);
    
    //--------------------------------------------------    
    //邮件内容 
    //from := "test1@newbt.net";
    from := "clq_test@163.com";
    to := "clq@newbt.net";
    
    SendLine("MAIL FROM: <" + from +">"); //注意"<" 符号和前面的空格。空格在协议中有和没有都可能，最好还是有 
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line);        
    
    SendLine("RCPT TO: <" + to+ ">");
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line);        
    
    SendLine("DATA");
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line)        
    
    // = "From: \"test1@newbt.net\" <test1@newbt.net>\r\nTo: \"clq@newbt.net\" <clq@newbt.net>\r\nSubject: test golang\r\nDate: Sun, 21 Jan 2018 11:48:15 GMT\r\n\r\nHello World.\r\n";//邮件内容，正式的应该用一个函数生成 
    s = MakeMail(from,to,"test golang","Hello World.");
    SendLine(s);    
    
    
    s = "\r\n.\r\n"; //邮件结束符 
    SendLine(s);
    
    //收取一行
    line = _RecvLine();
    fmt.Println("recv:" + line)        
    
}//



//这只是个简单的内容，真实的邮件内容复杂得多
func MakeMail(from,to,subject,text string)(string) {
    //s := "From: \"test1@newbt.net\" <test1@newbt.net>\r\nTo: \"clq@newbt.net\" <clq@newbt.net>\r\nSubject: test golang\r\nDate: Sun, 21 Jan 2018 11:48:15 GMT\r\n\r\nHello World.\r\n";//邮件内容，正式的应该用一个函数生成 
    s := "From: \"" + from + "\"\r\nTo: \"" + to + "\" " + to + "\r\nSubject: " + subject + 
        "\r\nDate: Sun, 21 Jan 2018 11:48:15 GMT\r\n\r\n" + //内容前是两个回车换行
        text + "\r\n";
    
    return s;    

}//

*/



