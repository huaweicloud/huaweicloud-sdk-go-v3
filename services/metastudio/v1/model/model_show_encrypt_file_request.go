package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEncryptFileRequest Request Object
type ShowEncryptFileRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 任务id
	JobId string `json:"job_id"`

	// 一次性token
	OnceToken string `json:"once_token"`
}

func (o ShowEncryptFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEncryptFileRequest struct{}"
	}

	return strings.Join([]string{"ShowEncryptFileRequest", string(data)}, " ")
}
