package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResourceList 资源列表
type QuotaResourceList struct {

	// 资源列表
	Resources *[]QuotaResource `json:"resources,omitempty"`
}

func (o QuotaResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourceList struct{}"
	}

	return strings.Join([]string{"QuotaResourceList", string(data)}, " ")
}
