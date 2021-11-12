package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfApiGroupV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
}

func (o ShowDetailsOfApiGroupV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiGroupV2Request", string(data)}, " ")
}
