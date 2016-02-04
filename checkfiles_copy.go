package main

 import (
  "runtime"
  "os"
  "fmt"
  "strings"
  "io/ioutil"
  "log"
  "encoding/csv"
  "bufio"
 )

func process_single_dir(doneCh chan bool, fdir string) {
  fmt.Println("**Inside goroutine ", fdir)
  logdir := "/Users/walterkimbrough/code/pycode/details/" + fdir
  records := [][]string{}
  csvfile_name := "csv_out_2_" + fdir + ".csv"
  csvfile, err := os.Create(csvfile_name)
    if err != nil {
          fmt.Println("Error:", err)
          return
    }
  defer csvfile.Close()

  files, _ := ioutil.ReadDir(logdir)

  for _, f := range files {
    full_path := logdir + "/" + f.Name()
    fmt.Println(full_path)
    input, err := os.Open(full_path)
    scanner := bufio.NewScanner(input)
    defer input.Close()

    // input, err := ioutil.ReadFile(full_path)
    if err != nil {
            log.Fatalln(err)
            continue
    }
    // lines := strings.Split(string(input), "\n")

    for scanner.Scan(){
    // for _, line := range lines {
      line := scanner.Text()
      if strings.Contains(line, "exec enforce_watermark_select @user_id =") && !strings.Contains(line, "BLOCK") {
          my_array := strings.Split(strings.TrimSpace(line), ",")
          user_id_line := strings.Replace(my_array[0], "=", " ", -1)
          user_id := strings.TrimSpace(strings.Replace(user_id_line, "exec enforce_watermark_select @user_id", "", -1))
          file_id_line := strings.Replace(my_array[2], "=", " ", -1)
          file_id := strings.TrimSpace(strings.Replace(file_id_line, "@file_id", "", -1))
          row1 := []string{file_id, user_id}
          records = append(records, row1)
      }
    }
  }
  writer := csv.NewWriter(csvfile)
          err = writer.WriteAll(records) // flush everything into csvfile
          if err != nil {
                  fmt.Println("Error:", err)
                  return
          }
        doneCh<-true
}

func process_dir(doneCh chan bool){
  file_dirs := []string{"i1", "i2", "e1", "e2"}

  for _, fdir := range file_dirs {
    fmt.Println(fdir)
    go process_single_dir(doneCh, fdir)
  }
}

 func main() {
      runtime.GOMAXPROCS(4)

      doneCh := make(chan bool)
      process_dir(doneCh)
  for i:=0;i<4;i++ {
    <-doneCh
  }
 }
