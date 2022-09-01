package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例组信息。
type ListInstancesGroupResult struct {

	// 组ID。
	Id string `json:"id" xml:"id"`

	// 组状态。
	Status string `json:"status" xml:"status"`

	Volume *Volume `json:"volume" xml:"volume"`

	// 节点信息。
	Nodes []ListInstancesNodeResult `json:"nodes" xml:"nodes"`
}

func (o ListInstancesGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesGroupResult struct{}"
	}

	return strings.Join([]string{"ListInstancesGroupResult", string(data)}, " ")
}
