package web
import (
	"compress/gzip"
	"strings"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/gotools/logs"
)

type RequestParam struct {
	Method string //GET 或者 POST
	Url string //请求的url地址
	UrlParam string  //url 参数
	RequestHeaders map[string] string //request header 参数
	Cookie map[string] string
}

func (self* RequestParam) AddCookie(key string, value string) {
	self.Cookie[key] = value
}

func (self* RequestParam) CookieIsExist(key string) bool {
	if _, ok := self.Cookie[key]; ok {
		return true
	}
	return false
}

func (self* RequestParam) AddRequestHeader(key string, value string) {
	self.RequestHeaders[key] = value
}

type ResponseParam struct {
	ResponseHeaders map[string] []string //response header 参数
	Content string //请求到的HTML内容
	Err error
}

func (self* ResponseParam) UnzipContent()(string) {
	
	if len(self.Content) <= 0 {
		return "";
	}
	
	r, err := gzip.NewReader(strings.NewReader(self.Content))
	if err != nil {//当前内容没有经过gzip进行压缩
		return self.Content;
	}
	defer r.Close()
	
	undatas, err := ioutil.ReadAll(r)
	if err != nil {
		return "";
	} else {
		var result = ""
		if len (undatas) > 0 {
			result = string(undatas)
		}
		return result; 
	}
}

func HttpRequest(reqParam *RequestParam) *ResponseParam {
	resParam := &ResponseParam{Err:nil}
	client := &http.Client{}
	
	var urlBody io.Reader
	if len(reqParam.UrlParam) <= 0 {
		urlBody = nil
	} else {
		urlBody = strings.NewReader(reqParam.UrlParam)
	}

	req, err := http.NewRequest(reqParam.Method, reqParam.Url, urlBody);
	if err != nil {
		logs.Debug("new request error");
		return resParam;
	}
	for k, v := range reqParam.RequestHeaders {
		req.Header.Set(k, v);
	}

	resp, err := client.Do(req);
	if err != nil {
		logs.Debug("client do request error(%s)", err.Error())
		resParam.Err = err;
		return resParam
	}
	
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	
	if err != nil {
		logs.Debug("read all data error")
		resParam.Err = err;
		return resParam;
	}
	resParam.Content = string(body);
	resParam.ResponseHeaders = resp.Header;

	return resParam;
}