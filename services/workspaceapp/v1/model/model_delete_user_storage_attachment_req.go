package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserStorageAttachmentReq 删除个人存储及关联
type DeleteUserStorageAttachmentReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 200]
	Items *[]string `json:"items,omitempty"`
}

func (o DeleteUserStorageAttachmentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserStorageAttachmentReq struct{}"
	}

	return strings.Join([]string{"DeleteUserStorageAttachmentReq", string(data)}, " ")
}
