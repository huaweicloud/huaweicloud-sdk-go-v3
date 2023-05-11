package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteFsDirQuotaRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	Body *DeleteFsDirQuotaRequestBody `json:"body,omitempty"`
}

func (o DeleteFsDirQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsDirQuotaRequest struct{}"
	}

	return strings.Join([]string{"DeleteFsDirQuotaRequest", string(data)}, " ")
}
