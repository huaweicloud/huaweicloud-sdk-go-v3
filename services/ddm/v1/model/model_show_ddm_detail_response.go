package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmDetailResponse Response Object
type ShowDdmDetailResponse struct {

	// **参数解释**：  查询实例详情相关信息的返回集合。  **参数范围**：  不涉及。
	Instances      *[]InstanceDetail `json:"instances,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowDdmDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDdmDetailResponse", string(data)}, " ")
}
