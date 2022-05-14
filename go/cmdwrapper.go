package main

/* [SO] https://stackoverflow.com/a/20438245/4568140 */
import (
    "fmt"
    "log"
    "io"
    "os/exec"
    "os"
    // "sync"
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

func exeCmd(cmdline string, output string) {
    fmt.Println("Start execCmd() ")
    cmd :=GenCmd(cmdline)
    // check if assigned output file
    if output != "" { 
        fmt.Println("Start main() ")

        f, err := os.Create(output)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
        cmd.Stdout = f // set stdout to short-response.json
        err = cmd.Run()
        if err != nil {
            log.Fatal(err)
        }
        // Reopen file, copy to stdout confirm cmdline output is there
        f.Close()
        f, _ = os.Open(output)
        io.Copy(os.Stdout, f) 
    } else {
        fmt.Println("Start main() ")
        out, err := cmd.Output()
        if err != nil {
          fmt.Printf("%s", err)
        }
        fmt.Printf("%s", out)
    }
}
func main() {
    fmt.Println("Start main()")
    x := []string{ "jq (.data.legacyCollection.collectionsPage.stream.edges|map({node:(.node|{url,firstPublished,headline:{default:.headline.default},summary})})) as $edges|{data:{legacyCollection:{collectionsPage:{stream:{$edges}}}}} nytimes-response.json", "ls -al" }
    // x := []string{"echo newline >> foo.o", "echo newline >> f1.o", "echo newline >> f2.o"}
    // exeCmd(x[0],"jia-bash.json")
    exeCmd(x[0],"jq.txt")
    exeCmd(x[1],"ls-al.txt")
}

