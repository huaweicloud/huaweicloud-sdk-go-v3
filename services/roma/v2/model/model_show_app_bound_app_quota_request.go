package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppBoundAppQuotaRequest Request Object
type ShowAppBoundAppQuotaRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o ShowAppBoundAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppBoundAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowAppBoundAppQuotaRequest", string(data)}, " ")
}
