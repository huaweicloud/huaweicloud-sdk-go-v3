package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RemoveProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveProjectResponse struct{}"
	}

	return strings.Join([]string{"RemoveProjectResponse", string(data)}, " ")
}
