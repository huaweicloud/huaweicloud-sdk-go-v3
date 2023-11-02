package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceOperationsResponse Response Object
type ListInstanceOperationsResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 支持的操作列表
	Operations     *[]Operations `json:"operations,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListInstanceOperationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceOperationsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceOperationsResponse", string(data)}, " ")
}
