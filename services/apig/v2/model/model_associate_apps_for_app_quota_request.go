package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAppsForAppQuotaRequest Request Object
type AssociateAppsForAppQuotaRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据配额编号
	AppQuotaId string `json:"app_quota_id"`

	Body *CreateAppQuotaBindingApp `json:"body,omitempty"`
}

func (o AssociateAppsForAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsForAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"AssociateAppsForAppQuotaRequest", string(data)}, " ")
}
