package main

/* [SO] https://stackoverflow.com/a/20438245/4568140 */
import (
    "fmt"
    "os/exec"
    "sync"
    "strings"
)
// debugging 
func loopArr(arr []string) {
    // using for loop
    for index, element := range arr {
        fmt.Println("At index", index, "value is", element)
    }
}

// command line parser , generate exec.Command
// cmd is same command line as running in shell(remove single quote)
func GenCmd(cmdline string) *exec.Cmd {
    fmt.Println("command is ",cmdline)
    // splitting head => g++ parts => rest of the command
    parts := strings.Fields(cmdline)

    // loopArr(parts)
    head := parts[0]
    parts = parts[1:len(parts)]
    
    // exec cmd & collect output
    cmd:= exec.Command(head,parts...)
    fmt.Printf("Generated comdline : %s", cmd)
    return cmd
  }

func exeCmd(cmdline string, wg *sync.WaitGroup) {
    // fmt.Println("Start execCmd() ")
    cmd :=GenCmd(cmdline)
        out, err := cmd.Output()
        if err != nil {
          fmt.Printf("%s", err)
        }
        fmt.Printf("%s", out)
    wg.Done() // signal to waitgroup this goroutine complete
}

func main() {
    x := []string{ 
        `jq (.data.legacyCollection.collectionsPage.stream.edges`+
            `|map({node:(.node|{url,firstPublished,headline:{default:.headline.default},summary})})) as $edges`+
            `|{data:{legacyCollection:{collectionsPage:{stream:{$edges}}}}} nytimes-response.json`,
        "ls -al",
     }
    // x := []string{"echo newline >> foo.o", "echo newline >> f1.o", "echo newline >> f2.o"}
    // x:= make(map[string]string)

    cmdCnt:=len(x)
    wg := new(sync.WaitGroup)
    wg.Add(cmdCnt)

    for _,cmd:= range x{
        go exeCmd(cmd,wg) // empty string output to stdout
    }

    wg.Wait()
}

