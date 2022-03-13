package sort

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
)

// func generateFile(n int) (fileName string) {
// 	fileName = fmt.Sprintf("mocks/%dn-data.txt", n)

// 	file, err := os.Create(fileName)

// 	if err != nil {
// 		log.Fatalf("failed creating file: %s", err)
// 	}

// 	datawriter := bufio.NewWriter(file)

// 	for i := 0; i < n; i++ {
// 		_, _ = datawriter.WriteString(fmt.Sprintf("%d\n", rand.Intn(65535)))
// 	}

// 	datawriter.Flush()
// 	file.Close()

// 	return fileName
// }

//var data1kk string = generateFile(1_000_000)

func readFile(fileName string) (arr []int) {
	arr = make([]int, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		num, err := strconv.Atoi(line)

		if nil == err {
			arr = append(arr, num)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	return arr
}

func saveFile(arr []int, key string) (fileName string) {
	fileName = fmt.Sprintf("results/sorted-%s-%d-data.txt", key, len(arr))

	os.Remove(fileName)

	os.Mkdir("results", fs.ModeAppend)
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, v := range arr {
		_, _ = datawriter.WriteString(fmt.Sprintf("%d\n", v))
	}

	datawriter.Flush()
	file.Close()

	return fileName
}

var dataArr1kk []int = readFile("mocks/1000000n-data.txt")
