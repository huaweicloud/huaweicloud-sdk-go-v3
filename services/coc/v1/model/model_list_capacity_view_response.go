package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCapacityViewResponse Response Object
type ListCapacityViewResponse struct {

	// **参数解释：** 容量数据列表。 **取值范围：** 查询已选应用下的云服务容量数据组成列表，大小在0到500之间。
	Data           *[]CapacityOverviewResponseData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListCapacityViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCapacityViewResponse struct{}"
	}

	return strings.Join([]string{"ListCapacityViewResponse", string(data)}, " ")
}
