package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateShrinkNodesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateShrinkNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShrinkNodesResponse struct{}"
	}

	return strings.Join([]string{"UpdateShrinkNodesResponse", string(data)}, " ")
}
