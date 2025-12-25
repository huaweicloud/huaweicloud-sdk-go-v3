package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePipeRequestBody struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 分区个数；默认创建1个，最大支持创建10个分区
	Shards *int32 `json:"shards,omitempty"`

	// 数据的保存时间，单位为天；默认30天，取值范围为1~3600
	StoragePeriod *int32 `json:"storage_period,omitempty"`
}

func (o UpdatePipeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePipeRequestBody", string(data)}, " ")
}
