package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertKeyRequest Request Object
type ShowCertKeyRequest struct {

	// 迁移任务ID
	TaskId string `json:"task_id"`

	// 是否生成ca证书
	EnableCaCert *bool `json:"enable_ca_cert,omitempty"`
}

func (o ShowCertKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowCertKeyRequest", string(data)}, " ")
}
