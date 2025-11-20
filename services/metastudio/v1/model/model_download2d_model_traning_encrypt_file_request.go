package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Download2dModelTraningEncryptFileRequest Request Object
type Download2dModelTraningEncryptFileRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 任务id
	JobId string `json:"job_id"`

	// 一次性token
	OnceToken string `json:"once_token"`
}

func (o Download2dModelTraningEncryptFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Download2dModelTraningEncryptFileRequest struct{}"
	}

	return strings.Join([]string{"Download2dModelTraningEncryptFileRequest", string(data)}, " ")
}
