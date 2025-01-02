package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InterpreterAiDigitalInfo 传译员AI数字资产信息
type InterpreterAiDigitalInfo struct {

	// 数字资产类型：PUBLIC（系统公共）、PRIVATE（企业专用账号绑定）、LOCAL（企业本地通用）。
	Type *string `json:"type,omitempty"`

	// AI传译员场景下绑定使用的数字资产ID（数字人或TTS音色）。
	DigitalAccount *string `json:"digitalAccount,omitempty"`

	// 数字资产名称。
	DigitalName *string `json:"digitalName,omitempty"`

	// 专用数字资产绑定的发言人登录账号，翻译对象非匿名必填。
	PresenterAccount *string `json:"presenterAccount,omitempty"`

	// 专用数字资产绑定的发言人登录账号（匿名时），翻译对象匿名必填。
	PresenterRealNameAccount *string `json:"presenterRealNameAccount,omitempty"`

	// 专用数字资产绑定的发言人名称。
	PresenterName *string `json:"presenterName,omitempty"`

	// 发言人用户的userUUID。
	PresenterUserID *string `json:"presenterUserID,omitempty"`

	// 本地会议的会议id（第三方对接参数），数字资产为LOCAL时必填。
	LocalConfId *string `json:"localConfId,omitempty"`

	// 本地会议对接地址或域名。
	LocalConfAddr *string `json:"localConfAddr,omitempty"`

	// 本地会议对接鉴权信息。
	LocalAuthInfo *string `json:"localAuthInfo,omitempty"`

	// true：需要代理 false：不需要代理。
	LocalNeedProxy *bool `json:"localNeedProxy,omitempty"`

	// 本地会议获取动态鉴权信息Url。
	LocalAuthUrl *string `json:"localAuthUrl,omitempty"`

	// 本地会议鉴权AppId。
	LocalAuthAppId *string `json:"localAuthAppId,omitempty"`
}

func (o InterpreterAiDigitalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterpreterAiDigitalInfo struct{}"
	}

	return strings.Join([]string{"InterpreterAiDigitalInfo", string(data)}, " ")
}
