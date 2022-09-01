package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandShareRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id" xml:"share_id"`

	Body *ExpandShareRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ExpandShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandShareRequest struct{}"
	}

	return strings.Join([]string{"ExpandShareRequest", string(data)}, " ")
}
