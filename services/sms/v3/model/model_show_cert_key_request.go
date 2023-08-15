package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertKeyRequest Request Object
type ShowCertKeyRequest struct {

	// 迁移任务ID
	TaskId string `json:"task_id"`
}

func (o ShowCertKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowCertKeyRequest", string(data)}, " ")
}
