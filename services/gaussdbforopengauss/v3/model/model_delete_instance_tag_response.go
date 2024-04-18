package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceTagResponse Response Object
type DeleteInstanceTagResponse struct {

	// 处理结果
	Result *string `json:"result,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName   *string `json:"instance_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceTagResponse", string(data)}, " ")
}
