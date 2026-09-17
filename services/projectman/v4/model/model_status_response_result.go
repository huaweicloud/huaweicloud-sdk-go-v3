package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusResponseResult 工作项状态查询接口返回状态数据
type StatusResponseResult struct {

	// **参数解释**： 状态对象。  **取值范围**： 不涉及。
	Status *[]StatusVoIpd `json:"status,omitempty"`
}

func (o StatusResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusResponseResult struct{}"
	}

	return strings.Join([]string{"StatusResponseResult", string(data)}, " ")
}
