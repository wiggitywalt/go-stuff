package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"waltsutil"
)

func process_single_dir(dir_to_process string, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
	logdir := "/Users/walterkimbrough/code/pycode/details/" + dir_to_process
	records := [][]string{}

	time_prefix := strconv.Itoa(t.Hour()) + "_" + strconv.Itoa(t.Minute())
	csvfile_name := "CSV_" + time_prefix + dir_to_process + ".csv"
	csvfile, err := os.Create(csvfile_name)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	files, _ := waltsutil.ReadDir(logdir)

	for _, f := range files {
		full_path := logdir + "/" + f.Name()

		input, err := os.Open(full_path)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		//read line by line here.
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
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
		//explicitly close file. If you use defer while opening a large number of files,
		//eventually you will receive an error message about 'too many files open'.
		input.Close()
	}

	writer := csv.NewWriter(csvfile)
	err = writer.WriteAll(records) // flush everything into csvfile
	if err != nil {
		fmt.Println("Error:", err)
		return
	} //end of func logic
	writer.Flush()
} //end psd func

func main() {
	runtime.GOMAXPROCS(4)
	file_dirs := []string{"i1", "i2", "e1", "e2"}

	var wg sync.WaitGroup
	wg.Add(len(file_dirs))

	start := time.Now()

	for _, dir := range file_dirs {
		go process_single_dir(dir, &wg)
	}

	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println(elapsed)
}
