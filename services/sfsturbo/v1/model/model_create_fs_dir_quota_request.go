package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFsDirQuotaRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	Body *CreateFsDirQuotaRequestBody `json:"body,omitempty"`
}

func (o CreateFsDirQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsDirQuotaRequest struct{}"
	}

	return strings.Join([]string{"CreateFsDirQuotaRequest", string(data)}, " ")
}
