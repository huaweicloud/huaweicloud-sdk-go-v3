package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBrokersRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o ListBrokersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersRequest struct{}"
	}

	return strings.Join([]string{"ListBrokersRequest", string(data)}, " ")
}
