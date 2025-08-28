package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateMembersRequestBody This is a auto create Body Object
type BatchCreateMembersRequestBody struct {

	// **参数解释**：批量添加member请求参数。  **约束限制**：不涉及
	Members []BatchCreateMembersOption `json:"members"`
}

func (o BatchCreateMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersRequestBody", string(data)}, " ")
}
