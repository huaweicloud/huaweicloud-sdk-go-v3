package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataPolicyReq 允许/禁止删除数据策略
type DataPolicyReq struct {
	DeletePolicy *PolicyType `json:"delete_policy"`
}

func (o DataPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPolicyReq struct{}"
	}

	return strings.Join([]string{"DataPolicyReq", string(data)}, " ")
}
