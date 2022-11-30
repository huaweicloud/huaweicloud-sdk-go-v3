package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyEpsQuotaRequestBody struct {

	// 需要修改的企业配额列表。
	EpsQuotas []EpsQuotasOption `json:"eps_quotas"`
}

func (o ModifyEpsQuotaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEpsQuotaRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyEpsQuotaRequestBody", string(data)}, " ")
}
