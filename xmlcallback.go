package fourinone

import (
	"bytes"
	"encoding/xml"
	"github.com/qtzheng/fourinone/utils"
	"io/ioutil"
)

type xmlCallback struct {
	textFlag                                                           bool
	propsAl                                                            []interface{}
	curProps                                                           utils.MergedConfig
	curKey, c_PROPSROW_DESC, c_KEY_DESC, curPROPSROW_DESC, curKEY_DESC string
}

func (x *XmlCallback) loadXml(filePath string) {
	var (
		t   xml.Token
		err error
	)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		utils.LogInfo(utils.Exit, "xml配置文件加载失败")
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			switch name {
			case "PROPSTABLE":
				x.propsAl = make([]interface{})
			case "PROPSROW":
				for _, attr := range token.Attr {
					if attr.Name.Local == "DESC" {
						x.curPROPSROW_DESC = attr.Value
					}
				}
			default:
				for _, attr := range token.Attr {
					if attr.Name.Local == "DESC" {
						x.curKEY_DESC = attr.Value
					}
				}
				x.curKEY_DESC = token.Name.Local
				x.textFlag = true
			}
		case xml.EndElement:
			name := token.Name.Local
			switch name {
			case "PROPSTABLE":
			case "PROPSROW":
				if x.c_PROPSROW_DESC == null || (x.curPROPSROW_DESC != null && x.curPROPSROW_DESC.equals(x.c_PROPSROW_DESC)) {
					x.propsAl.add(x.curProps)
				}
			}
		}
	}
}
