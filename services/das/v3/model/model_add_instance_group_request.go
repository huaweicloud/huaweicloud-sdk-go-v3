package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInstanceGroupRequest Request Object
type AddInstanceGroupRequest struct {
	Body *AddInstanceGroupRequestBody `json:"body,omitempty"`
}

func (o AddInstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"AddInstanceGroupRequest", string(data)}, " ")
}
