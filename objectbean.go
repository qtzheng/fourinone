package fourinone

type Serializable interface{}
type ObjectBean interface{
	ToObject()interface{}
	GetName()string
	GetDomain()string
	GetNode() string
}