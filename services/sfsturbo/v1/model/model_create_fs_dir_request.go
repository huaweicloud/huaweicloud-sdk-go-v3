package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFsDirRequest Request Object
type CreateFsDirRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *CreateFsDirRequestBody `json:"body,omitempty"`
}

func (o CreateFsDirRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsDirRequest struct{}"
	}

	return strings.Join([]string{"CreateFsDirRequest", string(data)}, " ")
}
