package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishModelResponse Response Object
type PublishModelResponse struct {

	// **参数解释**： 生成的模型ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ModelId        *string `json:"model_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishModelResponse struct{}"
	}

	return strings.Join([]string{"PublishModelResponse", string(data)}, " ")
}
