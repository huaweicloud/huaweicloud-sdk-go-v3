package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BuildCommandRiskDetailListResponseInfo struct {

	// **参数解释** 镜像构建指令风险详情总数 **取值范围** 取值0-2147483547
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**: 镜像构建指令风险详情列表 **取值范围**: 取值0-2147483647
	RiskList *[]BuildCommandRiskDetailResponseInfo `json:"risk_list,omitempty"`
}

func (o BuildCommandRiskDetailListResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildCommandRiskDetailListResponseInfo struct{}"
	}

	return strings.Join([]string{"BuildCommandRiskDetailListResponseInfo", string(data)}, " ")
}
