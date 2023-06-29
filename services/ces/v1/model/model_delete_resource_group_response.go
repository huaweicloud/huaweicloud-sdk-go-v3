package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceGroupResponse Response Object
type DeleteResourceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupResponse", string(data)}, " ")
}
