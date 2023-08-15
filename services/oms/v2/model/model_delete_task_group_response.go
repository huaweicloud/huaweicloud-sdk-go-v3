package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskGroupResponse Response Object
type DeleteTaskGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTaskGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskGroupResponse", string(data)}, " ")
}
