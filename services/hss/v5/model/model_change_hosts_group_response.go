package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeHostsGroupResponse Response Object
type ChangeHostsGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeHostsGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostsGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangeHostsGroupResponse", string(data)}, " ")
}
