package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceTempCredentialRequest Request Object
type CreateInstanceTempCredentialRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`
}

func (o CreateInstanceTempCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTempCredentialRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceTempCredentialRequest", string(data)}, " ")
}
