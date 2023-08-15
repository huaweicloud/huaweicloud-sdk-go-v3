package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpGroupResponse Response Object
type DeleteIpGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupResponse", string(data)}, " ")
}
