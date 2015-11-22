package fourinone

import (
	"github.com/qtzheng/fourinone/utils"
)

type AsyncTask interface {
	Task()
}
type AsyncExector struct {
	Task AsyncTask
}

func (a *AsyncExector) Run() {
	defer func() {
		err := recover()
		utils.RecoverError(err)
	}()
	if a.Task != nil {
		go a.Task.Task()
	} else {
		utils.LogInfo(utils.Info, "AsyncExector not have Task")
	}
}
