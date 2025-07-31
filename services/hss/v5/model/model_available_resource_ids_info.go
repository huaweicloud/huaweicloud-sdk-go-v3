package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableResourceIdsInfo struct {

	// **参数解释**： 资源ID **取值范围**： 字符长度1-256位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 当前时间 **取值范围**： 字符长度1-64位
	CurrentTime *string `json:"current_time,omitempty"`

	// **参数解释**： 是否共享配额 **取值范围**：   - shared：共享的   - unshared：非共享的
	SharedQuota *string `json:"shared_quota,omitempty"`
}

func (o AvailableResourceIdsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableResourceIdsInfo struct{}"
	}

	return strings.Join([]string{"AvailableResourceIdsInfo", string(data)}, " ")
}
