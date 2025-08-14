package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudStorageAttachmentReq 删除个人存储及关联。
type DeleteCloudStorageAttachmentReq struct {

	// 用户id，请求数量区间 [1, 200]。
	Items *[]string `json:"items,omitempty"`
}

func (o DeleteCloudStorageAttachmentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudStorageAttachmentReq struct{}"
	}

	return strings.Join([]string{"DeleteCloudStorageAttachmentReq", string(data)}, " ")
}
