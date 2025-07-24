package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFsDirQuotaRequest Request Object
type UpdateFsDirQuotaRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *UpdateFsDirQuotaRequestBody `json:"body,omitempty"`
}

func (o UpdateFsDirQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFsDirQuotaRequest struct{}"
	}

	return strings.Join([]string{"UpdateFsDirQuotaRequest", string(data)}, " ")
}
