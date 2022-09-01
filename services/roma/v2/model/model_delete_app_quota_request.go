package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAppQuotaRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 客户端配额编号
	AppQuotaId string `json:"app_quota_id" xml:"app_quota_id"`
}

func (o DeleteAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppQuotaRequest", string(data)}, " ")
}
