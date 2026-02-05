package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAppQuotaWithAppRequest Request Object
type DisassociateAppQuotaWithAppRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据配额编号
	AppQuotaId string `json:"app_quota_id"`

	// 凭据编号
	AppId string `json:"app_id"`
}

func (o DisassociateAppQuotaWithAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppQuotaWithAppRequest struct{}"
	}

	return strings.Join([]string{"DisassociateAppQuotaWithAppRequest", string(data)}, " ")
}
