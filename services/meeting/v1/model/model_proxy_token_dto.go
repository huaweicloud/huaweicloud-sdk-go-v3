package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 代理鉴权信息
type ProxyTokenDto struct {
	// 代理鉴权服务器的短token字符串

	AccessToken string `json:"accessToken"`
	// 代理鉴权服务器的长token字符串

	LongAccessToken *string `json:"longAccessToken,omitempty"`
	// token有效时长，单位：秒。

	ValidPeriod *int64 `json:"validPeriod,omitempty"`
	// 中台地址。

	MiddleEndUrl *string `json:"middleEndUrl,omitempty"`
	// 中台内网地址

	MiddleEndInnerUrl *string `json:"middleEndInnerUrl,omitempty"`
	// 是否开启二次路由

	EnableRerouting *bool `json:"enableRerouting,omitempty"`
}

func (o ProxyTokenDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyTokenDto struct{}"
	}

	return strings.Join([]string{"ProxyTokenDto", string(data)}, " ")
}
