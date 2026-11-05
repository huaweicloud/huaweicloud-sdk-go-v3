package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceStatusesResponse Response Object
type ListServiceStatusesResponse struct {

	// **参数解释**：服务开通状态。 **取值范围**：数组元素个数限制为16个
	ServiceList    *[]ServiceStatusItem `json:"service_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListServiceStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceStatusesResponse struct{}"
	}

	return strings.Join([]string{"ListServiceStatusesResponse", string(data)}, " ")
}
