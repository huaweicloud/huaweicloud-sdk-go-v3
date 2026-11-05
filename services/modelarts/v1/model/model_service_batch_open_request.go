package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceBatchOpenRequest struct {

	// **参数解释**：需要开通的服务列表。 **取值范围**：数组元素个数限制为16个。
	ServiceList *[]ServiceOpenItem `json:"service_list,omitempty"`
}

func (o ServiceBatchOpenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceBatchOpenRequest struct{}"
	}

	return strings.Join([]string{"ServiceBatchOpenRequest", string(data)}, " ")
}
