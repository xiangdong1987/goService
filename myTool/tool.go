package myTool

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//数字转ip
func Long2ip(ip int32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

//打包
func Pack(format string, params ...interface{}) (rs []byte, err error) {
	if len(format) != len(params) {
		err = errors.New("Format is not correct ")
	}
	i := 0

	buf := new(bytes.Buffer)
	byteOrder := binary.BigEndian
	for _, value := range params {
		if string(format[i]) == "N" {
			//fmt.Println(value)
			binary.Write(buf, byteOrder, value)
		}
		i++
	}
	return buf.Bytes(), err
}

//解包
func Unpack(format string, data []byte, params ...interface{}) error {
	if len(format) != len(params) {
		return errors.New("Format is not correct ")
	}
	//fmt.Println(string(data))
	buffer := bytes.NewReader(data)
	var err error
	i := 0
	for _, value := range params {
		if string(format[i]) == "N" {
			err = binary.Read(buffer, binary.BigEndian, value)
			//fmt.Println(value)
		}
		i++
	}
	return err
}

//一次性读取全部文件
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func Jsondecode(jsonStr []byte, returnData interface{}) {
	buffer := new(bytes.Buffer)
	buffer.Write(jsonStr)
	//fmt.Println(string(bodyStr))
	err := json.NewDecoder(buffer).Decode(&returnData)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
