package main

import (
	"os"
	"io/ioutil"
	"log"
	"bufio"
	"io"
	"fmt"
)

/**
	0 done 创建文件，创建文件夹
	1  读取文件的集中方式
	1.1 写文件
	2 done 检测文件或者文件夹是否存在

 */

const (
	TMP_DIR = "tmp"
	DATA_DIR = "data/"
)

func main()  {
	makeTempDir()
	//simpleWriteFile()
	//allWriteData()
	//bufWriteRead()
	justWriteRead()

}


func makeTempDir(){
	isExist,_ := checkPath(TMP_DIR)
	if !isExist{
		currentPath,_ :=os.Getwd()
		createDir(currentPath+"/"+TMP_DIR, 0755)
	}
}


//通过buff去读写
func bufWriteRead()  {
	currentPath,_ :=os.Getwd()
	fileName := currentPath+"/"+TMP_DIR+"/bufWriteRead.log"

	isExist,_ := checkPath(fileName)
	var err error
	var fp *os.File
	if !isExist{
		fp,err = createFileFunc1(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		checkError(err)
	}else{
		fp,err = os.OpenFile(fileName,os.O_RDWR, 0755 )
		checkError(err)
	}
	defer fp.Close()

	fp2,_ := os.Open(currentPath+"/"+DATA_DIR+"/pid3000.txt")
	bufRead := bufio.NewReader(fp2)

	bufWriter := bufio.NewWriter(fp)

	for {
		line,err := bufRead.ReadString('\n')
		if err != nil && err == io.EOF{
			break
		}
		bufWriter.WriteString(line)
		fmt.Println(line)
	}
}

//去读写
func justWriteRead()  {
	currentPath,_ :=os.Getwd()
	fileName := currentPath+"/"+TMP_DIR+"/justWriteRead.log"

	isExist,_ := checkPath(fileName)
	var err error
	var fp *os.File
	if !isExist{
		fp,err = createFileFunc1(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		checkError(err)
	}else{
		fp,err = os.OpenFile(fileName,os.O_RDWR, 0755 )
		checkError(err)
	}
	defer fp.Close()

	fp2,_ := os.Open(currentPath+"/"+DATA_DIR+"/pid3000.txt")
	byteData := make([]byte,1024)

	for {
		n,err := fp2.Read(byteData)
		if err != nil && err == io.EOF{
			break
		}
		if n == 0{
			break
		}
		fp.Write(byteData[0:n])
	}
}

//全量写和读
func allWriteData()  {
	currentPath,_ :=os.Getwd()
	fileName := currentPath+"/"+TMP_DIR+"/allWriteData.log"
	bytesData := []byte("adfsdfdsfds1 dfds 131 41 431 5 1  da f dsf ds fa ds");
	ioutil.WriteFile(fileName, bytesData, 0666)
	fp,_ := os.Open(fileName)
	bytesData2,_ := ioutil.ReadAll(fp)
	log.Println(string(bytesData2))
	bytesData3,_ := ioutil.ReadFile(fileName)
	log.Println(string(bytesData3))
}

//简单读写
func simpleWriteFile(){
	//创建文件
	currentPath,_ :=os.Getwd()
	fileName := currentPath+"/"+TMP_DIR+"/simpleWriteFile.log"
	isExist,_ := checkPath(TMP_DIR)

	var err error
	var fp *os.File
	if !isExist{
		fp,err = createFileFunc1(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		checkError(err)
	}else{
		fp,err = os.OpenFile(fileName,os.O_RDWR, 0755 )
		checkError(err)
	}
	defer fp.Close()

	fp.Write([]byte("afdsd"))
	fp.WriteString("cdef\n")
}

func checkError(err error){
	if err != nil{
		panic(err)
	}

}


func testFile(){
	//创建目录
	//创建文件

}




func createFileFunc1(name string, flag int, perm os.FileMode)(*os.File, error){
	return os.OpenFile(name, flag, perm)
}

func createFileFunc2(file string)(*os.File, error){
	return os.Create(file)
}

func createFileFun3(ptr uintptr, name string)*os.File{
	return os.NewFile(ptr,name)
}


func createDir(path string, mode os.FileMode)error{
	return os.Mkdir(path, mode)
}

func checkPath(file string)(bool,error){
	_,err := os.Stat(file)

	if err == nil{
		return true,nil
	}
	if os.IsExist(err){
		return false,nil
	}
	return false,err
}