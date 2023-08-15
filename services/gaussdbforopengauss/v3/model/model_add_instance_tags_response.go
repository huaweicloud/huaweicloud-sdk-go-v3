package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInstanceTagsResponse Response Object
type AddInstanceTagsResponse struct {

	// 添加标签的实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 添加标签的实例名称。
	InstanceName   *string `json:"instance_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"AddInstanceTagsResponse", string(data)}, " ")
}
