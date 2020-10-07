package main

import (
	"log"
	"os"
	"bufio"
	"fmt"	
	"strings"
	"io"
	"io/ioutil"
)
var targetpath string = ""
var motionconfpath string = ""
var archivename = "archive"
// ------------------------------------------------
func main() {
	if len(os.Args) == 2 {
		motionconfpath = os.Args[1]	
	}else{
		fmt.Println("##########################################################")
		fmt.Println("Usage: ./movetovideoarchive full_path_to_file_motion.conf")
		fmt.Println("##########################################################")
		fmt.Println("Example: ./movetovideoarchive /etc/motion/motion.conf")
		os.Exit(3)
	}
	getsourcePath()
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		s := strings.Split(f.Name(), ".")
		if(len(s) == 2 && s[1] == "mkv"){
			fn := strings.Split(s[0], "-")
			year := string(fn[1][0:4])
			month := string(fn[1][4:6])
			day := string(fn[1][6:8])
			newpath := targetpath+"/"+archivename+"/"+year+"/"+month+"/"+day+"/"+fn[0]			
			os.MkdirAll(newpath, os.ModePerm)
			err := os.Rename(targetpath+"/"+f.Name(), newpath+"/"+f.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
// ------------------------------------------------
func getsourcePath() bool {	
	if len(motionconfpath) == 0 {
		os.Exit(3)
	}
	file, err := os.Open(motionconfpath)
	if err != nil {
		fmt.Println("File motion.conf not found!")
		os.Exit(3)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	eql := 0
	for {
		line, err := reader.ReadString('\n')		
		if strings.HasPrefix(line, "target_dir") {			
			s := strings.Split(line, " ")
			targetpath = strings.TrimSpace(s[1])
			eql = 0
			break
		}else{
			eql++			
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error!")
			fmt.Println(err)
			os.Exit(3)
		}
	}
	if(eql > 0){
		fmt.Println("##########################################################")
		fmt.Println("In your motion.conf parameter \"target_dir\" is not found!")
		fmt.Println("##########################################################")
		os.Exit(3)
	}
	return true
}
