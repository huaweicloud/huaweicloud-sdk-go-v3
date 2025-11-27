package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInstanceToGroupRequest Request Object
type AddInstanceToGroupRequest struct {
	Body *AddInstanceToGroupRequestBody `json:"body,omitempty"`
}

func (o AddInstanceToGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceToGroupRequest struct{}"
	}

	return strings.Join([]string{"AddInstanceToGroupRequest", string(data)}, " ")
}
