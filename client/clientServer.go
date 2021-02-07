package client

import (
	"youDaoClient/util"
)

const reqUrl = "http://server.wxz01.top/fanyi/youdao"

type reqServer struct {
	Content string `json:"content"`
}

const ZhToEn = "zh-CHS2en"//汉译英
const EnToZh = "en2zh-CHS"//英译汉

type resServer struct {
	Code int `json:"code"`
	Data Data `json:"data"`
	Status string `json:"status"`
	Taketime float64 `json:"taketime"`
}
type Basic struct {
	Phonetic string `json:"phonetic"`
	UkPhonetic string `json:"uk-phonetic"`
	UsPhonetic string `json:"us-phonetic"`
	Explains []string `json:"explains"`
}
type Web struct {
	Key string `json:"key"`
	Value []string `json:"value"`
}
type List struct {
	ErrorCode string `json:"errorCode"`
	Query string `json:"query"`
	Translation []string `json:"translation"`
	Basic Basic `json:"basic"`
	Web []Web `json:"web"`
	L string `json:"l"`
	ReturnPhrase []string `json:"returnPhrase"`
	IsWord bool `json:"isWord"`
}
type Data struct {
	List List `json:"list"`
	Msg string `json:"msg"`
}

func YDFanYi(q string) *List {
	req := &reqServer{
		Content: q,
	}
	return req.reqSer()
}

func (r *reqServer)reqSer() *List {
	res :=  &resServer{}
	_,err := util.PostJson(reqUrl,r,nil, res)
	if err != nil {
		panic(err.Error())
	}
	if res.Code != 0 {
		panic("Service unavailable")
	}
	return &res.Data.List
}