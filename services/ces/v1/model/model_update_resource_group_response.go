package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateResourceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupResponse", string(data)}, " ")
}
