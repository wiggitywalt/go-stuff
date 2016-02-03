package main

import(
  "fmt"
  // "encoding/csv"
  // "os"
  )

func main() {
	users := []string{"walter","amy","alex","patrick","laserboogie"}
	cats := []string{"maisy","link_meese","midi","pearlie","shucka"}
	twoD := make([][]string, 2)
	for i := 0; i < 2; i++ {
          twoD[i] = make([]string, 2)
        }
	fmt.Println(twoD)
	// for x := 0; x < len(users); x++{
  //   twoD[0][x] = cats[x]
  //   twoD[1][x] = users[x]
	// }

  // csvfile, err := os.Create("/Users/walterkimbrough/Desktop/output.csv")
  //         if err != nil {
  //                 fmt.Println("Error:", err)
  //                 return
  //         }
  //         defer csvfile.Close()
  //         writer := csv.NewWriter(csvfile)
  //                  err = writer.WriteAll(twoD) // flush everything into csvfile
  //                  if err != nil {
  //                          fmt.Println("Error:", err)
  //                          return
  //                  }

}
