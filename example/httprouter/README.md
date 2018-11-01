
最基本的http的服务器的定义

```
//定义路由和处理器
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

//文件服务器
fs := http.FileServer(http.Dir("static/"))

//FileSystem的定义
type FileSystem interface {
	Open(name string) (File, error)
}
//返回值
func FileServer(root FileSystem) Handler {
  	return &fileHandler{root}
}
//fileHandler的定义
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}



http.Handle("/static/", http.StripPrefix("/static/", fs))

//启动服务
http.ListenAndServe(":8089", nil)


```

httprouter的功能

    1 精准匹配 Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params)





http包支持路由插件
```
//http Handler的定义
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

//func
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

//运行http服务
ListenAndServe(addr string, handler Handler)
```

httprouter的原理和实现

