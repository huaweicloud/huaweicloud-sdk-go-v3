package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelGroupResponse Response Object
type DeleteModelGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteModelGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteModelGroupResponse", string(data)}, " ")
}
