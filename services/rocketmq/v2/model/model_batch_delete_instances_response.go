package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstancesResponse Response Object
type BatchDeleteInstancesResponse struct {

	// **参数解释**： 修改实例的结果。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Results        *[]BatchDeleteInstanceRespResults `json:"results,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchDeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesResponse", string(data)}, " ")
}
