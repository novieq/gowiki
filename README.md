# Go web services

1. Simple wiki page using go
2. A simple mux used to handle requests in go

Go code execution flow. Excellent resource available in Building web application with golang

Let's take a look at the whole execution flow.

Call http.HandleFunc
Call HandleFunc of DefaultServeMux
Call Handle of DefaultServeMux
Add router rules to map[string]muxEntry of DefaultServeMux
Call http.ListenAndServe(":9090", nil)
Instantiate Server
Call ListenAndServe method of Server
Call net.Listen("tcp", addr) to listen to port
Start a loop and accept requests in the loop body
Instantiate a Conn and start a goroutine for every request: go c.serve()
Read request data: w, err := c.readRequest()
Check whether handler is empty or not, if it's empty then use DefaultServeMux
Call ServeHTTP of handler
Execute code in DefaultServeMux in this case
Choose handler by URL and execute code in that handler function: mux.handler.ServeHTTP(w, r)
How to choose handler: A. Check router rules for this URL B. Call ServeHTTP in that handler if there is one C. Call ServeHTTP of NotFoundHandler otherwise
