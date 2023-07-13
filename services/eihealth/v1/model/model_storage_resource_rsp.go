package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageResourceRsp struct {
	Spec *SpecDto `json:"spec"`

	// 使用量
	Size int64 `json:"size"`

	// 计费模式
	ChargeMode string `json:"charge_mode"`

	// 购买时间
	CreateTime string `json:"create_time"`

	// 状态
	Status string `json:"status"`
}

func (o StorageResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageResourceRsp struct{}"
	}

	return strings.Join([]string{"StorageResourceRsp", string(data)}, " ")
}
