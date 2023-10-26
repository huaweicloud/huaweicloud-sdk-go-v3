package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHpcShareRequest Request Object
type UpdateHpcShareRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *UpdateHpcShareRequestBody `json:"body,omitempty"`
}

func (o UpdateHpcShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHpcShareRequest struct{}"
	}

	return strings.Join([]string{"UpdateHpcShareRequest", string(data)}, " ")
}
