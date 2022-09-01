package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSharedTagRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 共享ID
	ShareId string `json:"share_id" xml:"share_id"`

	Body *CreateSharedTagRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateSharedTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSharedTagRequest", string(data)}, " ")
}
