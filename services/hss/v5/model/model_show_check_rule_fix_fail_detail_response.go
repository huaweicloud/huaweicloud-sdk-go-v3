package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckRuleFixFailDetailResponse Response Object
type ShowCheckRuleFixFailDetailResponse struct {

	// 修复失败原因列表
	FailDetailList *[]CheckRuleFixFailDetailInfo `json:"fail_detail_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowCheckRuleFixFailDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckRuleFixFailDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckRuleFixFailDetailResponse", string(data)}, " ")
}
