package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeShareNameRequest Request Object
type ChangeShareNameRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *ChangeShareNameReq `json:"body,omitempty"`
}

func (o ChangeShareNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareNameRequest struct{}"
	}

	return strings.Join([]string{"ChangeShareNameRequest", string(data)}, " ")
}
