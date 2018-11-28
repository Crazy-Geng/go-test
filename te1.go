package main

import (
	"fmt"
	//"github.com/360EntSecGroup-Skylar/excelize"
	//"reflect"
	//"crypto/md5"
	//"encoding/base64"
	//"strings"
	//"encoding/json"
	"os"
)

type test struct {
	A string
	B bool
}

func main() {
	/*
		tt := "aa,bb,cc,dd"
		res := make(map[string]interface{})
		ret := make([]map[string]interface{}, 0)
		fmt.Println(ret)
		for _, v := range strings.Split(tt, ",") {
			fmt.Println(&ret)
			res["name"] = v
			fmt.Println(&ret)
			//ret = append(ret, map[string]interface{}{"name": v})
			ret = append(ret, res)
			fmt.Println(&res)
			fmt.Println(&ret)
		}
		fmt.Println(ret)
		bb := "aaaa"
		fmt.Println(&bb)
	*/
	//aa := "HHHHe Hffff"
	//fmt.Errorf(format, ...)
	err := os.MkdirAll("/tmp/fktest111", 0755)
	if err != nil {
		fmt.Errorf("create mkdir is error %v", err)
	}

	fmt.Println("ok~")

}
