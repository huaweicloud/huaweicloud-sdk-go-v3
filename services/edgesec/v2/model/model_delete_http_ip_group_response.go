package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpIpGroupResponse Response Object
type DeleteHttpIpGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpIpGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpIpGroupResponse", string(data)}, " ")
}
