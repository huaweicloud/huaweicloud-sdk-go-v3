package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResources 资源配额详情列表
type QuotaResources struct {

	// 资源配额列表
	Resources *[]QuotaDetail `json:"resources,omitempty"`
}

func (o QuotaResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResources struct{}"
	}

	return strings.Join([]string{"QuotaResources", string(data)}, " ")
}
