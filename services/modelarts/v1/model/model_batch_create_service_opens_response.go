package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateServiceOpensResponse Response Object
type BatchCreateServiceOpensResponse struct {

	// **参数解释**：服务开通状态。 **取值范围**：数组元素个数限制为16个
	ServiceList    *[]ServiceStatusItem `json:"service_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchCreateServiceOpensResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServiceOpensResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateServiceOpensResponse", string(data)}, " ")
}
