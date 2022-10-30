

//原始 socket 的读写
//来自我自己的文章 “一步一步 .. 11.完整的发送示例与go语言”
//..-clq/p/8446372.html

package main //clq


import (
    "fmt"
    "bufio"
//    "crypto/tls"
//    "encoding/base64"
//    "errors"
//    "io"
    "net"
//    "net/smtp" //clq add
//    "net/textproto"
    "strings"
    "strconv"
	"time"
	"crypto/tls"
//	"crypto/rand"	
)

//var gConn net.Conn;
//var gRead * bufio.Reader;
//var gWrite * bufio.Writer;

//可以放到这样的类里
type TcpClient struct {
    Conn net.Conn;
    Read * bufio.Reader;
    Write * bufio.Writer;
}//


func Connect(host string, port int) (net.Conn, * bufio.Reader, * bufio.Writer) {
    
    addr := host + ":" + strconv.Itoa(port);
    conn, err := net.Dial("tcp", addr); //用 DialTimeout 可设置超时
    if err != nil {
        return nil, nil, nil
    }
    
    reader := bufio.NewReader(conn);
    writer := bufio.NewWriter(conn);
    
    //writer.WriteString("EHLO\r\n");
    //writer.Flush();
    
    //host, _, _ := net.SplitHostPort(addr)
    //return NewClient(conn, host)    
    return conn, reader, writer;
}//

//带超时的连接，单位是秒
func ConnectTimeOut(host string, port int, timeout_second int) (net.Conn, * bufio.Reader, * bufio.Writer) {
    
    addr := host + ":" + strconv.Itoa(port);
    //conn, err := net.Dial("tcp", addr); //用 DialTimeout 可设置超时
	
	//http://api.w.inmobi.cn/showad/v3
	//conn, err := net.DialTimeout("tcp", "api.w.inmobi.cn:80", 10 * time.Second);
	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout_second) * time.Second); //超时单位为秒，原始函数有点难理解

		
    if err != nil {
        return nil, nil, nil
    }
    
    reader := bufio.NewReader(conn);
    writer := bufio.NewWriter(conn);
    
    //writer.WriteString("EHLO\r\n");
    //writer.Flush();
    
    //host, _, _ := net.SplitHostPort(addr)
    //return NewClient(conn, host)    
    return conn, reader, writer;
}//

//带超时的连接，单位是秒//TCP 的
func Connect_TimeOut_TCP_old(host string, port int, timeout_second int) (net.Conn, * bufio.Reader, * bufio.Writer) {
    
    addr := host + ":" + strconv.Itoa(port);
    //conn, err := net.Dial("tcp", addr); //用 DialTimeout 可设置超时
	
	//http://api.w.inmobi.cn/showad/v3
	//conn, err := net.DialTimeout("tcp", "api.w.inmobi.cn:80", 10 * time.Second);
	//conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout_second) * time.Second); //超时单位为秒，原始函数有点难理解
	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout_second) * time.Second); //超时单位为秒，原始函数有点难理解
	
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
	
	fmt.Println("conn_tcp_p:", conn_tcp_p);
    
    //writer.WriteString("EHLO\r\n");
    //writer.Flush();
    
    //host, _, _ := net.SplitHostPort(addr)
    //return NewClient(conn, host)    
    return conn, reader, writer;
}//

//带超时的连接，单位是秒//TCP 的
func Connect_TimeOut_TCP(host string, port int, timeout_second int) (net.Conn, * bufio.Reader, * bufio.Writer) {
    
    addr := host + ":" + strconv.Itoa(port);

	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout_second) * time.Second); //超时单位为秒，原始函数有点难理解
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

//设置某个连接的超时，注意 golang 的超时值是一个绝对时间，所以得到数据后要再次调用
func SetConnectTimeOut(conn net.Conn, timeout_second int) {
    conn_tcp, ok := conn.(* net.TCPConn)
    if ok {
        fmt.Println("SetConnectTimeOut_tcp");
		SetConnectTimeOut_tcp(conn_tcp, timeout_second);
    }//

    conn_ssl, ok := conn.(* tls.Conn)
    if ok {
        fmt.Println("SetConnectTimeOut_ssl");
		SetConnectTimeOut_ssl(conn_ssl, timeout_second);
    }//

	
}


//设置某个连接的超时，注意 golang 的超时值是一个绝对时间，所以得到数据后要再次调用
func SetConnectTimeOut_tcp(conn * net.TCPConn, timeout_second int) {
	//if err := conn.SetReadDeadline(time.Time{}); err != nil {
	//    return err
	
	//if err := conn.SetReadDeadline(time.Now().Add(time.Minute*3)); err != nil {
	
	//    return err
	//conn.SetReadDeadline(time.Now().Add(time.Minute * timeout_second)); //单位为分钟
	conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout_second))); //单位为秒

}//

//ssl 的要另外弄
func SetConnectTimeOut_ssl(conn * tls.Conn, timeout_second int) {
	//if err := conn.SetReadDeadline(time.Time{}); err != nil {
	//    return err
	
	//if err := conn.SetReadDeadline(time.Now().Add(time.Minute*3)); err != nil {
	
	//    return err
	//conn.SetReadDeadline(time.Now().Add(time.Minute * timeout_second)); //单位为分钟
	conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout_second))); //单位为秒

}//


//设置某个连接的超时，注意 golang 的超时值是一个绝对时间，所以得到数据后要再次调用//这个是取消

//设置最后期限(超时)
//首先，你需要理解Go提供的最初级的网络超时实现：Deadlines（最后期限）。

//func SetConnectTimeOut_non(conn * net.Conn) { //net.Conn 是不行的，要用具体的 tcp 或者 udp 才有
func SetConnectTimeOut_non(conn * net.TCPConn) {
	//if err := conn.SetReadDeadline(time.Time{}); err != nil {
	//    return err
	
	//var conn *net.TCPConn //tcp 才有
	conn.SetReadDeadline(time.Time{}); //golang 的取消语法很特殊
	
	//SetDeadline(time.Time{})就是恢复为阻塞模式。
	
	//net.DialTCP

}//


//收取一行,可再优化 
//func RecvLine(conn *net.Conn) (string) {
//func RecvLine(conn net.Conn, reader * bufio.Reader) (string) {
func _RecvLine(gRead * bufio.Reader, gWrite * bufio.Writer) (string) {
	
	defer PrintError("_RecvLine()"); //不一定有效果，因为后面的 gRead.ReadString 会多次调用 nil 的指针
//	def
//	for ;; {
		
//	    err:=recover();
//		if err!=nil{
	
//			fmt.Print("err:[ + funcName + ]");
//	        fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	
//	    }else {break;}		
		
//	}
    
    //defer conn.Close();
    ////reader := bufio.NewReader(conn);
    //reader := bufio.NewReaderSize(conn,409600)
    
    //line, err := reader.ReadString('\n'); //如何设定超时?
    line, err := gRead.ReadString('\n'); //如何设定超时?
    
    if err != nil { return ""; }
    
    line = strings.Split(line, "\r")[0]; //还要再去掉 "\r"，其实不去掉也可以
    
    return line;
}//

func _RecvLine2(gRead * bufio.Reader, gWrite * bufio.Writer) (string, error) {
	
	defer PrintError("_RecvLine()"); //不一定有效果，因为后面的 gRead.ReadString 会多次调用 nil 的指针
//	def
//	for ;; {
		
//	    err:=recover();
//		if err!=nil{
	
//			fmt.Print("err:[ + funcName + ]");
//	        fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	
//	    }else {break;}		
		
//	}
    
    //defer conn.Close();
    ////reader := bufio.NewReader(conn);
    //reader := bufio.NewReaderSize(conn,409600)
    
    //line, err := reader.ReadString('\n'); //如何设定超时?
    line, err := gRead.ReadString('\n'); //如何设定超时?
    
    if err != nil { return "", err; }
    
    line = strings.Split(line, "\r")[0]; //还要再去掉 "\r"，其实不去掉也可以
    
    return line, nil;
}//


//2022 可指定结束符号(分隔符)的 //使用终止分隔符的
//目前的返回值中应该包括了分隔符
func _RecvString(gRead * bufio.Reader, gWrite * bufio.Writer, sp string) (string) {
	
	defer PrintError("_RecvString()"); //不一定有效果，因为后面的 gRead.ReadString 会多次调用 nil 的指针

    
    //defer conn.Close();
    ////reader := bufio.NewReader(conn);
    //reader := bufio.NewReaderSize(conn,409600)
    
    //line, err := reader.ReadString('\n'); //如何设定超时?
    ////line, err := gRead.ReadString('\n'); //如何设定超时?
    line, err := gRead.ReadString(sp[0]); //如何设定超时?
    
    
    if err != nil { return ""; }
    
    //line = strings.Split(line, "\r")[0]; //还要再去掉 "\r"，其实不去掉也可以
    
    return line;
}//


//接收到一个缓冲区中去//参考非常稳定了的 D:\gopath\src_http_proxy\tcp_server.go
//func RecvBuf(conn net.Conn, obuf []byte) int{
func RecvBuf(conn net.Conn, obuf * []byte) int{ //这种声明很怪异，不过还算能用
	
	defer PrintError("RecvBuf()");
	
	buf := make([]byte, 128);
	//obuf := make([]byte, 0); //obuf 在函数调用前要这样初始化，应该是一个字节切片
	
	//conn.SetReadDeadline(time.Now().Add(time.Second*10)); //golang 的超时很是奇怪,要每次设置
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("RecvBuf() 读取服务器数据异常:", err.Error());
		return 0;
		//break;
	}
	
	//obuf = append(obuf, buf[:n]...);
	*obuf = append(*obuf, buf[:n]...);
	//body = append(body, buf[:n]...);
	
	
	////fmt.Println("RecvBuf() n:" , n);
	////fmt.Println("RecvBuf() 服务器返回:" + string(buf[0:n]));
	////fmt.Println("RecvBuf() 服务器返回2:" + string(*obuf));
	//fmt.Println(gbk2utf8(string(buf[0:c]))); //应该指的是从 buf 中取 0 到 c 的字节
		
	return n;
}//		

func SendLine(gRead * bufio.Reader, gWrite * bufio.Writer, line string){
	
	defer PrintError("SendLine()");
	
    gWrite.WriteString(line + "\r\n");
    gWrite.Flush();
}//

//2022
func SendString(gRead * bufio.Reader, gWrite * bufio.Writer, line string){
	
	defer PrintError("SendLine()");
	
    //gWrite.WriteString(line + "\r\n");
    gWrite.WriteString(line);
    gWrite.Flush();
}//

//解码一行命令,这里比较简单就是按空格进行分隔就行了
func DecodeCmd(line string, sp string) ([]string){
	
	defer PrintError("DecodeCmd()");

    //String[] tmp = line.split(sp); //用空格分开//“.”和“|”都是转义字符，必须得加"\\";//不一定是空格也有可能是其他的
    //String[] cmds = {"", "", "", "", ""}; //先定义多几个，以面后面使用时产生异常
    
    tmp := strings.Split(line, sp);
    //var cmds = [5]string{"", "", "", "", ""}; //先定义多几个，以面后面使用时产生异常
    var cmds = []string{"", "", "", "", ""}; //先定义多几个，以面后面使用时产生异常
    //i:=0;
    for i:=0;i<len(tmp);i++ {
        if i >= len(cmds) { break;}
        cmds[i] = tmp[i];
    }
    return []string(cmds);
}//

//以后再优化
func GetValue(line string, sp1 string, sp2 string) (string){
	
	//smtp_host := GetValue(to, "@", ""); //这个会崩溃，原因未明
	//smtp_host := GetValue(to, "@", " ");//golang 的有点特殊，以后再改，第二个参数不能是空字符串	

    cmds := DecodeCmd(line, sp1);
	
	if len(sp2) == 0 { return cmds[1]; } //第二个分隔符为空的情况下要特殊处理
	
    cmds = DecodeCmd(cmds[1], sp2);
	
	//--------------------------------------------------
	//第一个分隔符为空的情况下要特殊处理
	if len(sp1) == 0 { 
		cmds = DecodeCmd(line, sp2);
		return cmds[0]; 
	} //第一个分隔符为空的情况下要特殊处理
	
	
	//--------------------------------------------------
    return cmds[0];
}//

//读取多行结果
func RecvMCmd(gRead * bufio.Reader, gWrite * bufio.Writer) (string) {
    i := 0;
    //index := 0;
    //count := 0;
    rs := "";
    //var c rune='\r';
    //var c4 rune = '\0'; //判断第4个字符//golang 似乎不支持这种表示
    
    mline := "";

    for i=0; i<50; i++ {
        rs = _RecvLine(gRead, gWrite); //只收取一行
        
        mline = mline + rs + "\r\n";
        
        //printf("\r\nRecvMCmd:%s\r\n", rs->str);
            
        if len(rs)<4 {break;} //长度要足够
        c4 := rs[4-1]; //第4个字符
        //if ('\x20' == c4) break; //"\xhh" 任意字符 二位十六进制//其实现在的转义符已经扩展得相当复杂，不建议用这个表示空格
        if ' ' == c4 { break;} //第4个字符是空格就表示读取完了//也可以判断 "250[空格]"
    
        
    }//

    return mline;
    
}//


