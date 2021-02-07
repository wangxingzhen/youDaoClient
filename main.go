package main

import (
	"flag"
	"fmt"
	"youDaoClient/client"
)




func main()  {
	flag.Parse()

	q := getStr()
	if q == "" {
		return
	}
	res := client.YDFanYi(q)
	//判断是英译汉还是汉译英
	fmt.Printf("\u001B[1;31;40m待翻译的文本：\u001B[0m\n\t\u001B[1;36;40m%v\u001B[0m\n",res.Query)//源语言
	if res.L == client.EnToZh {
		fmt.Printf("\u001B[1;31;40m全球：\u001B[0m\u001B[1;36;40m【%v】\u001B[0m\t",res.Basic.Phonetic)//发音
		fmt.Printf("\u001B[1;31;40m美：\u001B[0m\u001B[1;36;40m【%v】\u001B[0m\t",res.Basic.UsPhonetic)//发音
		fmt.Printf("\u001B[1;31;40m英：\u001B[0m\u001B[1;36;40m【%v】\u001B[0m\n",res.Basic.UkPhonetic)//发音
	}
	if len(res.Translation) == 0 {
		fmt.Println("\u001B[1;31;40m翻译失败\u001B[0m")//翻译失败
	}
	fmt.Println("\u001B[1;31;40m翻译结果：\u001B[0m")//翻译结果
	for _,v := range res.Translation {
		fmt.Printf("\t\u001B[1;36;40m%s\t\u001B[0m",v)//翻译结果
	}
	fmt.Println()
	if len(res.Basic.Explains) != 0 {
		fmt.Println("\u001B[1;31;40m其他结果：\u001B[0m")//其他翻译结果
		for _,v := range res.Basic.Explains {
			fmt.Printf("\u001B[1;36;40m\t%v\u001B[0m\n",v)
		}
	}
	if len(res.Web) != 0 {
		fmt.Println("\u001B[1;31;40m网络翻译：\u001B[0m")//网络翻译
		for _,v := range res.Web {
			fmt.Printf("\u001B[1;36;40m\t%v：", v.Key)
			for _,vv := range v.Value {
				fmt.Printf("%v; ", vv)
			}
			fmt.Println("\u001B[0m")
		}
	}
	return
}

//读取命令行里面的参数
func getStr() string {
	var str string
	allStr := flag.Args()
	for _,v := range allStr{
		if str == "" {
			str = v
		} else {
			str += " " + v
		}
	}
	return str
}
