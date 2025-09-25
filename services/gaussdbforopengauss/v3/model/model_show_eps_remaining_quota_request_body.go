package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowEpsRemainingQuotaRequestBody struct {

	// **参数解释**: 企业项目ID列表。 **约束限制**: 不涉及。
	EpsTags []string `json:"eps_tags"`
}

func (o ShowEpsRemainingQuotaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEpsRemainingQuotaRequestBody struct{}"
	}

	return strings.Join([]string{"ShowEpsRemainingQuotaRequestBody", string(data)}, " ")
}
