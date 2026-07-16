package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWorkspaceQuotasReqQuotas struct {

	// 资源标识。
	Resource string `json:"resource"`

	// 要修改的配额值。配额值为正整数或-1，-1代表不限制配额。配额值范围不能超过配额的最大值与最小值。可通过调用查询工作空间配额接口查询配额的最大值。
	Quota int32 `json:"quota"`
}

func (o UpdateWorkspaceQuotasReqQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceQuotasReqQuotas struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceQuotasReqQuotas", string(data)}, " ")
}
