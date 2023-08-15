package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAppsForAppQuotaRequest Request Object
type AssociateAppsForAppQuotaRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 客户端配额编号
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
