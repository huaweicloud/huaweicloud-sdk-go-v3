package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAppQuotaWithAppRequest Request Object
type DisassociateAppQuotaWithAppRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 客户端配额编号
	AppQuotaId string `json:"app_quota_id"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o DisassociateAppQuotaWithAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppQuotaWithAppRequest struct{}"
	}

	return strings.Join([]string{"DisassociateAppQuotaWithAppRequest", string(data)}, " ")
}
