package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandShareRequest Request Object
type ExpandShareRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *ExpandShareRequestBody `json:"body,omitempty"`
}

func (o ExpandShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandShareRequest struct{}"
	}

	return strings.Join([]string{"ExpandShareRequest", string(data)}, " ")
}
