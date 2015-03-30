package main

import (
    "fmt"
    "log"
    "net/http"
    "os" 
    "syscall"
    "github.com/go-martini/martini"
)

var (
    kernel32 = syscall.MustLoadDLL("kernel32.dll")
    procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func setStdHandle(stdhandle int32, handle syscall.Handle) error {
    r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
    if r0 == 0 {
        if e1 != 0 {
            return error(e1)
        }
        return syscall.EINVAL
    }
    return nil
}

// redirectStderr to the file passed in
func redirectStderr(f *os.File) error {
    return setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))    
}

func main() {
//    f, err := os.OpenFile("D:\\home\\site\\wwwroot\\testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    err = redirectStderr(f);
    
    http.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Error: %v", err)
    })

    if err != nil {
        http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
    } else {
        defer f.Close()
        log.SetOutput(f)
        log.Println("--->   UP @ " + os.Getenv("HTTP_PLATFORM_PORT") +"  <------")
    }

    m := martini.Classic()
    m.Get("/m", func() string {
      return "Hello world!"
    })
    m.Map(log.New(f, "[martini]", log.LstdFlags))    
    http.Handle("/m/", m)

    http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}

