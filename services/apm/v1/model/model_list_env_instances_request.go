package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvInstancesRequest Request Object
type ListEnvInstancesRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *InstanceSearchParam `json:"body,omitempty"`
}

func (o ListEnvInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListEnvInstancesRequest", string(data)}, " ")
}
