package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlideVerifyCodeCheckDto struct {
	// 必须和发送验证码时带的用户身份信息相同 maxLength：255 minLength：1

	User string `json:"user"`
	// 登录客户端类型。 * 0：Web客户端类型； * 5：cloudlink pc； * 6：cloudlink mobile； * 16：workplace pc； * 18：workplace mobile

	ClientType int32 `json:"clientType"`
	// 校验类型。 * 0：登录； * 1：忘记密码; 默认值：0

	CheckType *int32 `json:"checkType,omitempty"`
	// 发送滑块验证码返回的token字符串 maxLength：255 minLength：1

	Token string `json:"token"`
	// 抠出图形的X轴坐标。

	PointX int32 `json:"pointX"`
	// 滑动时间，单位ms。

	SlideTime int32 `json:"slideTime"`
}

func (o SlideVerifyCodeCheckDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlideVerifyCodeCheckDto struct{}"
	}

	return strings.Join([]string{"SlideVerifyCodeCheckDto", string(data)}, " ")
}
