package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupV5Response Response Object
type UpdateGroupV5Response struct {
	Group          *Group `json:"group,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateGroupV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupV5Response struct{}"
	}

	return strings.Join([]string{"UpdateGroupV5Response", string(data)}, " ")
}
