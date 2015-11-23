package fourinone

import (
	"io/ioutil"
	"strings"
)

type MulBean struct {
	resourceBean ResourceBean
	nativeLangCode string
}

func (m *MulBean) Init(langCode string) {
	langCode = strings.ToUpper()
	switch langCode {
	case "":
		break
	case "ISO-8859-1":
		nativeLangCode = "ISO-8859-1"
	}
	case "GB2312":
		nativeLangCode = "GB2312"
	}
	case "BIG5":
		nativeLangCode = "BIG5"
	}
}
func (m *MulBean)GetString(key string,topStr ...string)string{
	return ""
}
func (m *MulBean)GetSpace()string{
	if m.nativeLangCode!=null&&m.nativeLangCode=="ISO-8859-1"{
			return "&nbsp;"
			}
	return ""
}
func (m *MulBean)GetFileString(relativeUri string)string{
	data,err:=ioutil.ReadFile(relativeUri)
	if utils.CheckError(err){
		return ""
	}else{
		return string(data)
	}
}
func NewMulBean(langCode string) {
	m := &MulBean{}
	m.resourceBean=NewResourceBean(langCode)
	m.Init(langCode)
}
