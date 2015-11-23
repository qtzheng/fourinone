package fourinone

type XmlUtil struct{
	
}
func(x *XmlUtil)GetXmlPropsByFile(filePath string,p ...string)[]interface{}{
	var PROPSROW_DESC,KEY_DESC string
	if len(p)==2{
		PROPSROW_DESC=p[0]
		KEY_DESC=p[1]
	}
	
}