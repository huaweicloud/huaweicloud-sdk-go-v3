package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationGroupResponse Response Object
type DeleteApplicationGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationGroupResponse", string(data)}, " ")
}
