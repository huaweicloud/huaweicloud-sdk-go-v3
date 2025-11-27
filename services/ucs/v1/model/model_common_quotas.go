package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonQuotas struct {

	// 配额资源列表
	Resources *[]QuotaResource `json:"resources,omitempty"`
}

func (o CommonQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonQuotas struct{}"
	}

	return strings.Join([]string{"CommonQuotas", string(data)}, " ")
}
