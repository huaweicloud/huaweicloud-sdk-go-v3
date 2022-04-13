package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstancesDetailsResponse struct {
	// 实例列表。

	Instances *[]ShowInstanceResp `json:"instances,omitempty"`
	// 实例个数。

	InstanceNum    *int32 `json:"instance_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesDetailsResponse", string(data)}, " ")
}
