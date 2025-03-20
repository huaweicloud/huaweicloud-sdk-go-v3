package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupV5Response Response Object
type CreateGroupV5Response struct {
	Group          *Group `json:"group,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateGroupV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupV5Response struct{}"
	}

	return strings.Join([]string{"CreateGroupV5Response", string(data)}, " ")
}
