package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDeploymentsRequest struct {
	// 偏移量。 当前偏移量，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询返回部署计划列表当前页面的数量。

	Limit *int32 `json:"limit,omitempty"`
	// 查询条件，部署计划状态，现只包含如下值： - open:部署计划处于未执行状态，可执行部署计划进行部署 - closed:部署计划已关闭，不可部署。

	Status *string `json:"status,omitempty"`
	// 查询条件，部署计划ID。

	Id *string `json:"id,omitempty"`
	// 查询条件，边缘业务ID。

	EdgecloudId *string `json:"edgecloud_id,omitempty"`
}

func (o ListDeploymentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeploymentsRequest struct{}"
	}

	return strings.Join([]string{"ListDeploymentsRequest", string(data)}, " ")
}
