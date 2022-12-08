package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSasTokenRequestBody struct {

	// 用于为用户生成令牌的密钥  取值为primary或者secondary  每个用户只有2个apikey，primary的原则就是对比apiKey时间最早的那个。  如果只有一个apikey，primary和secondary都指同一个apiKey
	Keytype string `json:"keytype"`

	// 令牌到期UTC时间，格式如：2019-04-21T00:44:24Z   日期符合以下格式： yyyy-MM-ddTHH:mm:ssZ 由 ISO 8601 标准指定。  最小值不小于15min，最大值不超过24h
	Expiry string `json:"expiry"`
}

func (o CreateSasTokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSasTokenRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSasTokenRequestBody", string(data)}, " ")
}
