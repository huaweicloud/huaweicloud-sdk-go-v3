package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubTasksByPostUsingRequest Request Object
type ListSubTasksByPostUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *SubTaskQueryByPageParams `json:"body,omitempty"`
}

func (o ListSubTasksByPostUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubTasksByPostUsingRequest struct{}"
	}

	return strings.Join([]string{"ListSubTasksByPostUsingRequest", string(data)}, " ")
}
