package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceRemainderData 配额剩余数量信息。
type ResourceRemainderData struct {

	// 资源类型。
	Type *string `json:"type,omitempty"`

	// 剩余数量。
	Remainder *int64 `json:"remainder,omitempty"`

	// 所需资源。
	Need *int64 `json:"need,omitempty"`

	// 磁盘类型。type为volume_gigabytes时，会返回该字段。
	VolumeType *string `json:"volume_type,omitempty"`
}

func (o ResourceRemainderData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceRemainderData struct{}"
	}

	return strings.Join([]string{"ResourceRemainderData", string(data)}, " ")
}
