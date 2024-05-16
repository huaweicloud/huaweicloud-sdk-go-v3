package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGroupRequest Request Object
type AddGroupRequest struct {
	Body *AddGroupRequestBody `json:"body,omitempty"`
}

func (o AddGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGroupRequest struct{}"
	}

	return strings.Join([]string{"AddGroupRequest", string(data)}, " ")
}
