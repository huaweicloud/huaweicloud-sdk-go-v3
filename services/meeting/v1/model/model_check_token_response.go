package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckTokenResponse struct {
	// 接入token字符串。

	AccessToken *string `json:"accessToken,omitempty"`
	// 用户IP。

	TokenIp *string `json:"tokenIp,omitempty"`
	// token有效时长，单位：秒。

	ValidPeriod *int64 `json:"validPeriod,omitempty"`
	// token的失效时间戳，单位：秒。

	ExpireTime *int64 `json:"expireTime,omitempty"`
	// 业务token的创建时间戳，单位：毫秒。

	CreateTime *int64 `json:"createTime,omitempty"`

	User *UserInfo `json:"user,omitempty"`
	// 登录帐号类型。 * 72：API调用类型

	ClientType *int32 `json:"clientType,omitempty"`
	// 抢占登录标识 * 0：非抢占 * 1：抢占  未启用

	ForceLoginInd *int32 `json:"forceLoginInd,omitempty"`
	// 是否首次登录（说明：首次登录表示尚未修改过密码。首次登录时，系统会提醒用户需要修改密码），默认值：false。

	FirstLogin *bool `json:"firstLogin,omitempty"`
	// 密码是否过期，默认值：false。

	PwdExpired *bool `json:"pwdExpired,omitempty"`
	// 密码有效天数

	DaysPwdAvailable *int32 `json:"daysPwdAvailable,omitempty"`

	ProxyToken *ProxyTokenDto `json:"proxyToken,omitempty"`
	// 是否延时删除状态

	DelayDelete *bool `json:"delayDelete,omitempty"`
	// token类型 * 0：用户ACCESS TOKEN； * 1：会控TOKEN * 2：一次性TOKEN

	TokenType *int32 `json:"tokenType,omitempty"`
	// 刷新token字符串。

	RefreshToken *string `json:"refreshToken,omitempty"`
	// 刷新token有效时长，单位：秒。

	RefreshValidPeriod *int64 `json:"refreshValidPeriod,omitempty"`
	// 刷新token的失效时间戳，单位：秒。

	RefreshExpireTime *int64 `json:"refreshExpireTime,omitempty"`
	// 刷新token的创建时间戳，单位：毫秒。

	RefreshCreateTime *int64 `json:"refreshCreateTime,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o CheckTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTokenResponse struct{}"
	}

	return strings.Join([]string{"CheckTokenResponse", string(data)}, " ")
}
