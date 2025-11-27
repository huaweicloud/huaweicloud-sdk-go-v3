package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceGroupRequest Request Object
type DeleteInstanceGroupRequest struct {
	Body *DeleteInstanceGroupRequestBody `json:"body,omitempty"`
}

func (o DeleteInstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceGroupRequest", string(data)}, " ")
}
