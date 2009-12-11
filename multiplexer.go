// 
// example of a compute server using channles and goroutines
// from here: http://golang.org/doc/go_tutorial.html
// 
// updated to shutdown the server cleanly before exiting


package main


import "fmt"

type request struct {
	a, b int;
	replyc chan int;
}

type binOp func(a, b int) int

// execut an arbiary operation aainst a request and post back the result
func run(op binOp, req *request) { 
	reply := op(req.a, req.b);
	req.replyc <- reply;
}

// run for ever processing requests
/*
func server(op binOp, service chan *request) {
	for {
		// consume a request
		req := <-service;
		// execute the request
		go run(op, req); // do not wait for result
	}
}
*/
func server(op binOp, service chan *request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req); // excute request
		case <-quit:
			return;
		}
	}
}

// start the server and return a feed to pass in work
/*
func startServer(op binOp) chan *request {
	req := make(chan *request);
	go server(op, req);
	return req;
}
*/ 
func startServer(op binOp) (service chan *request, quit chan bool) {
	service = make(chan *request);
	quit = make(chan bool);
	go server(op, service, quit);
	return service, quit;
}

// test the server
func main() {
	// defined function inline and boot the server with it
	//adder := startServer(func(a, b int) int {return a + b});
	adder, quit := startServer(func(a, b int) int {return a + b});
	const N = 100;
	var reqs [N]request;
	
	// send off all the requests
	for i := 0; i<N; i++ {
		req := &reqs[i];
		req.a = i;
		req.b = i + N;
		req.replyc = make(chan int);
		adder <- req;	
	}
	
	// consume all of the results
	for i := N-1; i >= 0; i-- {
		if <-reqs[i].replyc != N + 2*i {
			fmt.Println("fail at", i);
		}
	}
	
	// clean shutdown
	quit <- true;
	
	fmt.Println("done");
}