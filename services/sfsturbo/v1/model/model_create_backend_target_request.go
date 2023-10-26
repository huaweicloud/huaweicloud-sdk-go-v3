package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackendTargetRequest Request Object
type CreateBackendTargetRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	Body *CreateBackendTargetRequestBody `json:"body,omitempty"`
}

func (o CreateBackendTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackendTargetRequest struct{}"
	}

	return strings.Join([]string{"CreateBackendTargetRequest", string(data)}, " ")
}
