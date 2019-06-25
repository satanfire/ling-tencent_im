package tencentim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type msgContent struct {
	Text interface{} `json:"Text"`
}

type msgBody struct {
	MsgType    string     `json:"MsgType"`
	MsgContent msgContent `json:"MsgContent"`
}

// AndroidInfo android 离线消息
type AndroidInfo struct {
	Sound string `json:"Sound,omitempty"` // 离线推送声音文件路径
}

// ApnsInfo ios离线消息
type ApnsInfo struct {
	Sound     string `json:"Sound,omitempty"`     // 离线推送声音文件路径
	BadgeMode int    `json:"BadgeMode,omitempty"` // 这个字段缺省或者为 0 表示需要计数，为 1 表示本条消息不需要计数，即右上角图标数字不增加
	Title     string `json:"Title,omitempty"`     // 该字段用于标识 APNs 推送的标题，若填写则会覆盖最上层 Title
	SubTitle  string `json:"SubTitle,omitempty"`  // 该字段用于标识 APNs 推送的子标题
	Image     string `json:"Image,omitempty"`     // 该字段用于标识 APNs 携带的图片地址，当客户端拿到该字段时，可以通过下载图片资源的方式将图片展示在弹窗上
}

// OfflinePushInfo 离线消息
type OfflinePushInfo struct {
	PushFlag    int          `json:"PushFlag,omitempty"` // 0表示推送，1表示不离线推送
	Title       string       `json:"Title,omitempty"`    // 离线推送标题。该字段为 iOS 和 Android 共用
	Desc        string       `json:"Desc,omitempty"`     // 离线推送内容
	Ext         string       `json:"Ext,omitempty"`      // 离线推送透传内容
	AndroidInfo *AndroidInfo `json:"AndroidInfo,omitempty"`
	ApnsInfo    *ApnsInfo    `json:"ApnsInfo,omitempty"`
}

type sendChatSendMsg struct {
	SyncOtherMachine int              `json:"SyncOtherMachine,omitempty"` //1：把消息同步到 From_Account 在线终端和漫游上； 2：消息不同步至 From_Account； 若不填写默认情况下会将消息存 From_Account 漫游
	FromAccount      string           `json:"From_Account,omitempty"`     // 管理员指定某一帐号向其它帐号发送消息
	ToAccount        string           `json:"To_Account"`
	MsgLifeTime      int              `json:"MsgLifeTime,omitempty"`
	MsgRandom        uint32           `json:"MsgRandom"`
	MsgTimeStamp     int64            `json:"MsgTimeStamp"`
	MsgBody          []msgBody        `json:"MsgBody"`
	OfflinePushInfo  *OfflinePushInfo `json:"OfflinePushInfo,omitempty"` // 同时设置离线推送信息
}

// SendMsg 单发单聊消息
type SendMsg struct {
	QueryStringParam QueryStringParam
	SendMsgBody      sendChatSendMsg
}

// NewSendMsg 创建单发单聊消息
func NewSendMsg(adminUserSig, toAccount string, content interface{}) *SendMsg {
	qsp := QueryStringParam{
		AppID:   appID,
		UserSig: adminUserSig,
	}

	msg := msgBody{
		MsgType: TIMTextElemMsgType,
		MsgContent: msgContent{
			Text: content.(string),
		},
	}
	return &SendMsg{
		QueryStringParam: qsp,
		SendMsgBody: sendChatSendMsg{
			SyncOtherMachine: SyncOtherMachineNoSync,
			ToAccount:        toAccount,
			MsgLifeTime:      60,
			MsgRandom:        rand.Uint32(),
			MsgTimeStamp:     time.Now().Unix(),
			MsgBody:          []msgBody{msg},
		},
	}
}

// SetFromAccount 设置FromAccount
func (sm *SendMsg) SetFromAccount(fromAccount string) {
	sm.SendMsgBody.FromAccount = fromAccount
}

// SetOfflinePushInfo 设置离线消息, 如果你要赋值这个整体过来也是可以的
func (sm *SendMsg) SetOfflinePushInfo(sopi *OfflinePushInfo) {
	sm.SendMsgBody.OfflinePushInfo = sopi
}

func (sm *SendMsg) checkOfflinePushInfo() {
	if sm.SendMsgBody.OfflinePushInfo == nil {
		sm.SendMsgBody.OfflinePushInfo = &OfflinePushInfo{}
	}
}

func (sm *SendMsg) checkAndroidApns() {
	if sm.SendMsgBody.OfflinePushInfo.ApnsInfo == nil {
		sm.SendMsgBody.OfflinePushInfo.ApnsInfo = &ApnsInfo{}
	}

	if sm.SendMsgBody.OfflinePushInfo.AndroidInfo == nil {
		sm.SendMsgBody.OfflinePushInfo.AndroidInfo = &AndroidInfo{}
	}
}

// SetTitle 设置推送标题
func (sm *SendMsg) SetTitle(title string) {
	sm.checkOfflinePushInfo()
	sm.SendMsgBody.OfflinePushInfo.Title = title
}

// SetDesc 设置推送内容
func (sm *SendMsg) SetDesc(desc string) {
	sm.checkOfflinePushInfo()
	sm.SendMsgBody.OfflinePushInfo.Desc = desc
}

// SetExt 设置透传内容
func (sm *SendMsg) SetExt(ext string) {
	sm.checkOfflinePushInfo()
	sm.SendMsgBody.OfflinePushInfo.Ext = ext
}

// SetSound 推送声音文件路
func (sm *SendMsg) SetSound(sound string) {
	sm.checkOfflinePushInfo()
	sm.checkAndroidApns()
	sm.SendMsgBody.OfflinePushInfo.AndroidInfo.Sound = sound
	sm.SendMsgBody.OfflinePushInfo.ApnsInfo.Sound = sound
}

// SenApnsImage 携带的图片地址
func (sm *SendMsg) SenApnsImage(image string) {
	sm.checkOfflinePushInfo()
	sm.checkAndroidApns()
	sm.SendMsgBody.OfflinePushInfo.ApnsInfo.Image = image
}

// QueryString 返回Query string
func (sm *SendMsg) QueryString() string {
	return sm.QueryStringParam.BuildQueryString()
}

// Name 获取消息类型
func (sm *SendMsg) Name() string {
	return "SendMsg"
}

// URI 返回chat对应的URI
func (sm *SendMsg) URI() string {
	return V4OpenIMSendMsg
}

// Body 返回SendMsg的msg body
func (sm *SendMsg) Body() (*bytes.Buffer, error) {
	bytesData, err := json.Marshal(sm.SendMsgBody)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bytesData))
	return bytes.NewBuffer(bytesData), nil
}
