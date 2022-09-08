package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteApiGroupV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`
}

func (o DeleteApiGroupV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiGroupV2Request", string(data)}, " ")
}
