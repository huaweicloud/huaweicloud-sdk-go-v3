package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceIdResponse Response Object
type ListInstanceIdResponse struct {

	// 实例ID列表
	InstanceIds    *[]string `json:"instance_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListInstanceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceIdResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceIdResponse", string(data)}, " ")
}
