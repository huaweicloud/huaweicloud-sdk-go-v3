package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppQuotaRequest Request Object
type CreateAppQuotaRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *AppQuotaCreate `json:"body,omitempty"`
}

func (o CreateAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"CreateAppQuotaRequest", string(data)}, " ")
}
