package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息
type ListQuotasResult struct {

	// 配额列表
	Resources []Resources `json:"resources"`
}

func (o ListQuotasResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasResult struct{}"
	}

	return strings.Join([]string{"ListQuotasResult", string(data)}, " ")
}
