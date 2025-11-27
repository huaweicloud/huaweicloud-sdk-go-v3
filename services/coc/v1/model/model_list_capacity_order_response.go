package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCapacityOrderResponse Response Object
type ListCapacityOrderResponse struct {

	// **参数解释：** 容量排名数据。 **取值范围：** 应用或者组件或者分组容量数据的排名。
	Data           *[]CapacityOrderResponseData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListCapacityOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCapacityOrderResponse struct{}"
	}

	return strings.Join([]string{"ListCapacityOrderResponse", string(data)}, " ")
}
