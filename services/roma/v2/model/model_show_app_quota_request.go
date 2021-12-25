package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAppQuotaRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 客户端配额编号

	AppQuotaId string `json:"app_quota_id"`
}

func (o ShowAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowAppQuotaRequest", string(data)}, " ")
}
