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
	SyncOtherMachineSync   = 1
	SyncOtherMachineNoSync = 2
)

// QueryStringParam TIM uri query string param
type QueryStringParam struct {
	AppID   string `json:"-"`
	UserSig string `json:"-"`
}

func (qsp *QueryStringParam) BuildQueryString() string {
	return fmt.Sprintf("sdkappid=%s&identifier=%s&usersig=%s&random=%d&contenttype=json",
		qsp.AppID, "admin", qsp.UserSig, rand.Uint32())
}
