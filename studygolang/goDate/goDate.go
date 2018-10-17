package main

import (
	_ "time"
	"os"
	"fmt"
	_ "strconv"
	_ "flag"
	"strconv"
	"time"
	"reflect"
	"regexp"
	"strings"
)

type RegexCheck struct {}
var p = fmt.Println
func main() {
	  args := os.Args
	 lenArg :=args[1]
	 if args == nil || len(args) !=2 || strings.EqualFold(lenArg,"-") || strings.EqualFold(lenArg,"+") || strings.EqualFold(lenArg,"-0") || strings.EqualFold(lenArg,"+0") {
		 p("Parameter Error !")
		 return
	  }
     	isTrue,_ := regexp.MatchString("^[-\\+]?[\\d]*$", lenArg)  //验证字符串是否为数字(包含正负数)
	  if isTrue {
		  yday :=yesday(lenArg)
		  fmt.Println(yday)
	}else {
		  fmt.Println("Err:验证字符串是否为数字(包含正负数) ")
	  }
	    
}

func IsMail(val string) bool {
	return IsMatch(val, `\w[-._\w]*@\w[-._\w]*\.\w+`)
}

func IsPhone(val string) bool {
	if strings.HasPrefix(val, "+") {
		return IsMatch(val[1:], `\d{13}`)
	} else {
		return IsMatch(val, `\d{11}`)
	}
}


func IsEnglishIdentifier(val string, pattern ...string) bool {
	defpattern := "^[a-zA-Z0-9\\-\\_\\.]+$"
	if len(pattern) > 0 {
		defpattern = pattern[0]
	}

	return IsMatch(val, defpattern)
}

func IsMatch(val, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(val))
	if err != nil {
		return false
	}

	return match
}

//验证字符串是否为数字(包含正负数) //  ^[\-|0-9][0-9]*   , //  ^[\-|0-9][0-9]{1,}$
func (ic *RegexCheck) isNumeric (str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[-\\+]?[\\d]*$", s)
		if false == b {
			return b
		}
	}
	return b
}


func yesday(lenArg string) (string){
	lenArg1, _ := strconv.Atoi(lenArg)
			nTime := time.Now()
			yesTime := nTime.AddDate(0, 0, lenArg1)
			logDay := yesTime.Format("20060102")
 		return logDay
}

//类型检测 要检测的变量  期望变量类型
func typeChck(params interface{}, t string) bool {
	//数据初始化
	//fmt.Println("params_day:",params)
	var (
		return_val bool = false
	)
	v := reflect.ValueOf(params)
	//获取传递参数类型
	v_t := v.Type()

	//类型名称对比
	if v_t.String() == t {
		return_val = true
	}
	return return_val
}

//数字+字母  不限制大小写 6~30位
func (ic *RegexCheck) IsID(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9a-zA-Z]{6,30}$", s)
		if false == b {
			return b
		}
	}
	return b
}

//数字+字母+符号 6~30位
func (ic *RegexCheck) IsPwd(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9a-zA-Z@.]{6,30}$", s)
		if false == b {
			return b
		}
	}
	return b
}

func (ic *RegexCheck) IsInteger(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//纯小数
func (ic *RegexCheck) IsDecimals(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^\\d+\\.[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//手提电话（不带前缀）最高11位
func (ic *RegexCheck) IsCellphone(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^1[0-9]{10}$", s)
		if false == b {
			return b
		}
	}
	return b
}

//家用电话（不带前缀） 最高8位
func (ic *RegexCheck) IsTelephone(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]{8}$", s)
		if false == b {
			return b
		}
	}
	return b
}

//仅小写
func (ic *RegexCheck) IsEngishLowCase(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[a-z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//仅大写
func (ic *RegexCheck) IsEnglishCap(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[A-Z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//大小写混合
func (ic *RegexCheck) IsEnglish(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[A-Za-z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//邮箱 最高30位
func (ic *RegexCheck) IsEmail(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", s)
		if false == b {
			return b
		}
	}
	return b
}

