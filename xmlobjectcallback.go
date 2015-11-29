package fourinone

type XmlObjectCallback struct {
	textFlag                                                           bool
	objAl                                                              []ObjValue
	curObj                                                             *ObjValue
	curKey, c_PROPSROW_DESC, c_KEY_DESC, curPROPSROW_DESC, curKEY_DESC string
}

func NewXmlObjectCallback() *XmlObjectCallback {
	return &XmlObjectCallback{
		textFlag: false,
		objAl:    make([]ObjValue),
	}
}
func (x *XmlObjectCallback) parse(filePath) {
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
				x.objAl = make([]ObjValue)
			case "PROPSROW":
				for _, attr := range token.Attr {
					if attr.Name.Local == "DESC" {
						x.curPROPSROW_DESC = attr.Value
					}
				}
				x.curObj = new(ObjValue)
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
				break
			case "PROPSROW":
				if x.c_PROPSROW_DESC == "" || (x.curPROPSROW_DESC != "" && x.curPROPSROW_DESC == x.c_PROPSROW_DESC) {
					x.objAl = append(x.curObj)
				}
			default:
				x.textFlag = false
			}
		case xml.CharData:
			content := string([]byte(token))
			if x.textFlag {
				if KEY_DESC == "" || (curKEY_DESC != "" && curKEY_DESC == KEY_DESC) {
					curObj.Add(curKey, content.trim())
				}
			}
		}
	}
}
func (x *XmlObjectCallback) GetObjAl() []ObjValue {
	return x.objAl
}
func (x *XmlObjectCallback) SetPROPSROW_DESC(PROPSROW_DESC string) {
	x.c_PROPSROW_DESC = PROPSROW_DESC
}
func (x *XmlObjectCallback) SetKEY_DESC(KEY_DESC string) {
	x.KEY_DESC = KEY_DESC
}
