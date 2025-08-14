package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpsamlConfig struct {

	// 服务提供商断言响应地址
	AcsUrl string `json:"acs_url"`

	// 服务提供商签发者
	Issuer string `json:"issuer"`

	// 服务提供商元数据
	Metadata string `json:"metadata"`
}

func (o SpsamlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpsamlConfig struct{}"
	}

	return strings.Join([]string{"SpsamlConfig", string(data)}, " ")
}
