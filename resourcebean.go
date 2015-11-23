package fourinone

type ResourceBean struct {
	resourcesName string
}

func NewResourceBean(resourcesName string) *ResourceBean {
	return &ResourceBean{
		resourcesName: resourcesName,
	}
}
func (r *ResourceBean) GetString(key string) string {
	return ""
}
