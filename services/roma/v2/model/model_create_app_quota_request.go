package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppQuotaRequest Request Object
type CreateAppQuotaRequest struct {

	// 实例ID
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
