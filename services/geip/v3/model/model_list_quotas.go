package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotas 资源列表
type ListQuotas struct {

	// 资源列表
	Resources []ResourcesInfo `json:"resources"`
}

func (o ListQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotas struct{}"
	}

	return strings.Join([]string{"ListQuotas", string(data)}, " ")
}
