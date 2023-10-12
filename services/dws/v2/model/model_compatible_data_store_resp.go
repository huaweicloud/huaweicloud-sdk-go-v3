package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompatibleDataStoreResp struct {

	// 类型
	Type *string `json:"type,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o CompatibleDataStoreResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleDataStoreResp struct{}"
	}

	return strings.Join([]string{"CompatibleDataStoreResp", string(data)}, " ")
}
