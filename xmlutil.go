package fourinone

type XmlUtil struct{
	
}
func(x *XmlUtil)GetXmlPropsByFile(filePath string,p ...string)[]interface{}{
	var PROPSROW_DESC,KEY_DESC string
	if len(p)==2{
		PROPSROW_DESC=p[0]
		KEY_DESC=p[1]
	}
	//handler:=
	return nil
}
func (x *XmlUtil) GetXmlPropsByTable() []interface{}{
	return nil
}
func (x *XmlUtil) GetXmlPropsByObject() []interface{}{
	return nil
}
func (x *XmlUtil) GetXmlFileByTable() []interface{}{
	return nil
}
func (x *XmlUtil) GetXmlObjectByFile() []interface{}{
	return nil
}
func (x *XmlUtil) getDefaultConfig()string{
	config:=`<?xml version="1.0" encoding="UTF-8"?>
<PROPSTABLE DESC="TABLENAME">
  <PROPSROW DESC="PARK">
    <SERVICE>ParkService</SERVICE>
    <SERVERS>localhost:1888,localhost:1889</SERVERS>
    <SAFEMEMORYPER>0.95</SAFEMEMORYPER>
    <HEARTBEAT>3000</HEARTBEAT>
    <MAXDELAY>30000</MAXDELAY>
    <EXPIRATION>24</EXPIRATION>
    <CLEARPERIOD>0</CLEARPERIOD>
    <ALWAYSTRYLEADER>false</ALWAYSTRYLEADER>
    <STARTWEBAPP>true</STARTWEBAPP>
  </PROPSROW>
  <PROPSROW DESC="COOLHASH">
    <DATAROOT>data</DATAROOT>
    <KEYLENTH DESC="B">256</KEYLENTH>
    <VALUELENGTH DESC="M">2</VALUELENGTH>
    <REGIONLENGTH DESC="M">2</REGIONLENGTH>
    <LOADLENGTH DESC="M">64</LOADLENGTH>
    <HASHCAPACITY>1000000</HASHCAPACITY>
  </PROPSROW>
  <PROPSROW DESC="CACHE">
    <SERVICE>CacheService</SERVICE>
    <SERVERS>localhost:2000,localhost:2001</SERVERS>
  </PROPSROW>
  <PROPSROW DESC="CACHEGROUP">
    <STARTTIME>2010-01-01</STARTTIME>
    <GROUP>localhost:2000,localhost:2001@2010-01-01;localhost:2002,localhost:2003@2010-05-01;localhost:2004,localhost:2005@2010-05-01</GROUP>
  </PROPSROW>
  <PROPSROW DESC="CACHEGROUP">
    <STARTTIME>2018-05-01</STARTTIME>
    <GROUP>localhost:2008,localhost:2009@2018-05-01;localhost:2010,localhost:2011@2018-05-01</GROUP>
  </PROPSROW>
  <PROPSROW DESC="CACHEFACADE">
    <SERVICE>CacheFacadeService</SERVICE>
    <SERVERS>localhost:1998</SERVERS>
    <TRYKEYSNUM>100</TRYKEYSNUM>
  </PROPSROW>
  <PROPSROW DESC="WORKER">
    <TIMEOUT DESC="FALSE">2</TIMEOUT>
    <SERVERS>localhost:2088</SERVERS>
    <SERVICE>false</SERVICE>
  </PROPSROW>
  <PROPSROW DESC="CTOR">
    <!-- <CTORSERVERS>localhost:1988</CTORSERVERS> -->
    <INITSERVICES>100</INITSERVICES>
    <MAXSERVICES>500</MAXSERVICES>
  </PROPSROW>
  <PROPSROW DESC="COMPUTEMODE">
    <MODE DESC="DEFAULT">0</MODE>
    <MODE>1</MODE>
  </PROPSROW>
  <PROPSROW DESC="FTTP">
    <SERVERS>localhost:2121</SERVERS>
  </PROPSROW>
  <PROPSROW DESC="WEBAPP">
    <SERVERS>localhost:9080</SERVERS>
    <USERS>admin:admin,guest:123456,test:test</USERS>
  </PROPSROW>
  <PROPSROW DESC="LOG">
    <LEVELNAME>ALL</LEVELNAME>
    <LEVELNAME>SEVERE</LEVELNAME>
    <LEVELNAME>WARNING</LEVELNAME>
    <LEVELNAME>INFO</LEVELNAME>
    <LEVELNAME>CONFIG</LEVELNAME>
    <LEVELNAME DESC="LOGLEVEL">FINE</LEVELNAME>
    <LEVELNAME>FINER</LEVELNAME>
    <LEVELNAME>FINEST</LEVELNAME>
    <LEVELNAME>OFF</LEVELNAME>
    <INFO>true</INFO>
    <FINE>false</FINE>
  </PROPSROW>
</PROPSTABLE>`
}