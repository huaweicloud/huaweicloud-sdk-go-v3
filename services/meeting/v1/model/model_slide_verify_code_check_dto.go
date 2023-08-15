package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlideVerifyCodeCheckDto struct {

	// 必须和发送验证码时带的用户身份信息相同。
	User string `json:"user"`

	// 登录客户端类型。 * 0：Web客户端类型 * 5：PC客户端 * 6：移动客户端
	ClientType int32 `json:"clientType"`

	// 校验类型。默认值：0。 * 0：登录； * 1：忘记密码;
	CheckType *int32 `json:"checkType,omitempty"`

	// 验证码Token字符串。通过[[发送滑块验证码](https://support.huaweicloud.com/api-meeting/meeting_21_0100.html)](tag:hws)[[发送滑块验证码](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0100.html)](tag:hk) 接口获取。
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
