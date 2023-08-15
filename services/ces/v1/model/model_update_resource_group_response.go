package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupResponse Response Object
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
