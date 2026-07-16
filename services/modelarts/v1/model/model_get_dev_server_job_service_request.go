package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerJobServiceRequest Request Object
type GetDevServerJobServiceRequest struct {

	// **参数解释**：部署服务的id。 **约束限制**：字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o GetDevServerJobServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerJobServiceRequest struct{}"
	}

	return strings.Join([]string{"GetDevServerJobServiceRequest", string(data)}, " ")
}
