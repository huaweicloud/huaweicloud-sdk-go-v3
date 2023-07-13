package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSharedTagRequest Request Object
type CreateSharedTagRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 共享ID
	ShareId string `json:"share_id"`

	Body *CreateSharedTagRequestBody `json:"body,omitempty"`
}

func (o CreateSharedTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSharedTagRequest", string(data)}, " ")
}
