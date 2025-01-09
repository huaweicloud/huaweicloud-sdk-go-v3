package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateAppRulesReq 批量操作规则(规则ID)。
type BatchOperateAppRulesReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 50]。
	Items []string `json:"items"`
}

func (o BatchOperateAppRulesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateAppRulesReq struct{}"
	}

	return strings.Join([]string{"BatchOperateAppRulesReq", string(data)}, " ")
}
