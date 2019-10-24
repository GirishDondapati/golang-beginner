package main

import ("os"
	"fmt"
	"math/rand"
	"sync"
	"strconv"
	"path/filepath"
)

var a[3] string
var file_ves = make(map[string]int)
var wg sync.WaitGroup

func main(){
	setDriveDetails()
	for j := 0; j <= 2; j++ {
		dIndex := rand.Intn(3)
		driveLetter := a[dIndex]
		version := file_ves[driveLetter]
		fmt.Println(driveLetter)
		fmt.Println(version)
		wg.Add(1)
		go write_content(driveLetter,version)
		file_ves[driveLetter] = version+1
	}
	wg.Wait()
	getFilesNames()
}

// Get all files names
func getFilesNames(){
	for _, value := range a {
		var files []string
		filePath := value+":\\Temp_Files"
		err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
		})
		if err != nil {
			panic(err)
		}
		for _, file := range files {
		fmt.Println(file)
		}
	}
	
}

//It will take available tree drives namesin a array
func setDriveDetails(){
	i := 0
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ"{
	driveName := string(drive)
	_, err := os.Open(driveName+":\\")
	if err == nil && driveName != "C" && i < 3{
		a[i] = driveName
		//fmt.Println(string(drive))
		i++
		}
	}
	//fmt.Println(a)
}


func write_content(drive string, version int){
	fmt.Println(version)
	var file_dir string = drive+":\\Temp_Files\\A"+strconv.Itoa(version)+".txt"
	//var file_dir string = drive+":\\Temp_Files\\A.txt"
	sample_text := "this is testing line"
	f, err := os.Create(file_dir)
	check(err)
	n3, err := f.WriteString(sample_text)
	defer f.Close()
	fmt.Println(n3)
	wg.Done()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}