package fourinone

type ArrayAdapter struct {
	objinit  int
	objindex int
	abjarr   []interface{}
}

func (a *ArrayAdapter) auto() {
	if a.objindex == len(a.abjarr) {
		objarrnew := make([]interface{}, a.objindex+a.objinit)
		copy(objarrnew, a.abjarr)
		a.abjarr = objarrnew
	}
}
func NewArrayAdapter() *ArrayAdapter {
	ad := &ArrayAdapter{
		objinit:  0x40,
		objindex: 0,
	}
	ad.abjarr = make([]interface{}, ad.objinit)
	return ad
}
func (a *ArrayAdapter) Add(obj interface{}) {
	a.abjarr[a.objindex] = obj
	a.objindex++
}
