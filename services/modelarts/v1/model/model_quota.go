package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Quota ModelArts资源管理服务中资源的配额信息。
type Quota struct {

	// **参数解释**： 资源的配额信息。
	Resources *[]ResourceQuota `json:"resources,omitempty"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
