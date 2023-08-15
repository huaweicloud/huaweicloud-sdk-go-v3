package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigHistoriesRequest Request Object
type ListConfigHistoriesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示条数，最小值为1，最大值为1000，若不设置该参数，则为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigHistoriesRequest", string(data)}, " ")
}
