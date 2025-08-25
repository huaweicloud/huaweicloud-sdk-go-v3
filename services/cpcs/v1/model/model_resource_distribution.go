package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDistribution struct {

	// 密钥管理服务密钥数量
	Kms int32 `json:"kms"`
}

func (o ResourceDistribution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDistribution struct{}"
	}

	return strings.Join([]string{"ResourceDistribution", string(data)}, " ")
}
