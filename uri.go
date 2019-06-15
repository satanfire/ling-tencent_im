package tencentim

const schema = "https://console.tim.qq.com"

// 账号管理
const (
	// V4IMOpenLoginSvcAccountImport 导入单个帐号
	V4IMOpenLoginSvcAccountImport = "v4/im_open_login_svc/account_import"
	// V4IMOpenLoginSvcMultiAccountImport 导入批量帐号
	V4IMOpenLoginSvcMultiAccountImport = "v4/im_open_login_svc/multiaccount_import"
	// V4IMOpenLoginSvcKick 失效帐号登录态
	V4IMOpenLoginSvcKick = "v4/im_open_login_svc/kick"
)

// 在线状态
const (
	// V4OpenIMQueryState 获取用户在线状态
	V4OpenIMQueryState = "v4/openim/querystate"
)

// 资料管理
const (
	// V4ProfilePortraitGet 拉取资料
	V4ProfilePortraitGet = "v4/profile/portrait_get"
	// V4ProfilePortraitSet 设置资料
	V4ProfilePortraitSet = "v4/profile/portrait_set"
)

// chat
const (
	// V4OpenIMSendMsg 单发单聊消息
	V4OpenIMSendMsg = "v4/openim/sendmsg"
	// V4OpenIMBatchSendMsg 批量发单聊消息
	V4OpenIMBatchSendMsg = "v4/openim/batchsendmsg"
	// V4OpenIMImportMsg 导入单聊消息
	V4OpenIMImportMsg = "v4/openim/importmsg"
)
