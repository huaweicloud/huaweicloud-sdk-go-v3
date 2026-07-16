package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInferServicesResponse Response Object
type BatchDeleteInferServicesResponse struct {

	// **参数解释：** 服务响应返回体。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceResponses *[]ServiceResponse `json:"service_responses,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o BatchDeleteInferServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInferServicesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInferServicesResponse", string(data)}, " ")
}
