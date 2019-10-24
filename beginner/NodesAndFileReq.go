package main

/*
	It's like a multi node cluster environment, there are 3 node cluster( 3 directories in your system), you can add your file randomly in any of these three nodes. 
whenever create a new file or update an existing file should create next version of file and keep both the versions of file. 
	have to list down the file from all three nodes (all the versions of the files)
	find out the total storage consumed by all 3 nodes

conditions : 
	every N minutes the GC will remove the old file or unused files. 

*/
import (
	"fmt"
	"math/rand"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
	"os"
	"regexp"
	"time"
)

var nodes[3] string
var totSize int64

func main(){
	nodes = [3]string{"F:\\Programs\\GoLong\\Node1","F:\\Programs\\GoLong\\Node2","F:\\Programs\\GoLong\\Node3"}

	fmt.Println("1: Get all Files List And Total Size")
	fmt.Println("2: Insert data")

	var input int
	fmt.Scanln(&input)
	if(input == 1) {
		displayAllFileNames()
		fmt.Println("Total Size: ",totSize)
	}else if(input == 2) {
		var content string
		fmt.Println("Please enter content: ")
		fmt.Scanln(&content)

		if(len(content) > 0) {
			ranNo := randInt(0, len(nodes))
			//fmt.Println("Randaom Number: ",ranNo)
			fileVersion := getlatestVersionId()
			write_content(nodes[ranNo], fileVersion+1, content)
			//displayAllFileNames()
			fmt.Println("Total Size: ",totSize)
		} else {
			fmt.Println("Enter content is empty. So data not inserted")
		}
	}
}

func getlatestVersionId() int64{
	var latestVersionId int64 = 0
	files, err := IOReadDir()
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		startIdx := (strings.Index(file,"A")) + 1
		endIdx := (strings.Index(file,".txt"))
		versionStr := file[startIdx:endIdx]
		tempVersion,_ := strconv.ParseInt(versionStr, 0, 64)
		if(latestVersionId < tempVersion){
			latestVersionId = tempVersion
		}
	}
	return latestVersionId
}

func displayAllFileNames() {
	files, err := IOReadDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}

func IOReadDir() ([]string, error) {
	totSize = 0
	var files []string
	for _, root := range nodes {
		fileInfo, err := ioutil.ReadDir(root)
		if err != nil {
			return files, err
		}

		for _, file := range fileInfo {
			regx := regexp.MustCompile("A([0-9]+).txt")
			isFileNameMatch := regx.MatchString(file.Name())
			if isFileNameMatch{
				files = append(files, file.Name())
				totSize += file.Size()
			}
		}
	}
	return files, nil
}

func write_content(drivePath string, version int64,content string){
	var file_dir string = drivePath+"\\A"+strconv.Itoa(int(version))+".txt"
	file, err := os.Create(file_dir)
	check(err)
	file.WriteString(content)
	fmt.Println("Newly created file version is: ",file_dir)
	defer file.Close()

	file1, err := os.Open(file_dir)
	if err != nil {
		log.Fatal(err)
	}
	fi, _ := file1.Stat()
	totSize += fi.Size()
}

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}




