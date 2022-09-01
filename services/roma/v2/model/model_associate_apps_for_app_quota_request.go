package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateAppsForAppQuotaRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 客户端配额编号
	AppQuotaId string `json:"app_quota_id" xml:"app_quota_id"`

	Body *CreateAppQuotaBindingApp `json:"body,omitempty" xml:"body"`
}

func (o AssociateAppsForAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsForAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"AssociateAppsForAppQuotaRequest", string(data)}, " ")
}
