package tencentim

import (
	"fmt"
	"math/rand"
)

const (
	// TIMTextElemMsgType TIMTextElem msg type
	TIMTextElemMsgType = "TIMTextElem"
)

const (
	// SyncOtherMachineSync 若希望将消息同步至 From_Account，则 SyncOtherMachine 填写1
	SyncOtherMachineSync = 1
	// SyncOtherMachineNoSync  若不希望将消息同步至 From_Account，则 SyncOtherMachine 填写2
	SyncOtherMachineNoSync = 2
)

// QueryStringParam TIM uri query string param
type QueryStringParam struct {
	AppID   string `json:"-"`
	UserSig string `json:"-"`
}

// BuildQueryString 返回QueryString
func (qsp *QueryStringParam) BuildQueryString() string {
	return fmt.Sprintf("sdkappid=%s&identifier=%s&usersig=%s&random=%d&contenttype=json",
		qsp.AppID, "admin", qsp.UserSig, rand.Uint32())
}
