package fourinone

type ObjValue struct {
	data map[string]string
}

func NewObjValue() *ObjValue {
	return &ObjValue{
		data: make(map[string]string),
	}
}
func (o *ObjValue) Add(key, value string) {
	if _, ok := o.data[key]; !ok {
		o.data[key] = value
	}
}
func (o *ObjValue) Update(key, value string) {
	if _, ok := o.data[key]; ok {
		o.data[key] = value
	}
}
func(o *ObjValue)GetData(key,def string)string{
	if v,ok:=o.data[key];ok{
		return v
	}else{
		return def
	}
}
