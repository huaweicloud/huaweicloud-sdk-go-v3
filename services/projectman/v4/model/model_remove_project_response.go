package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveProjectResponse Response Object
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
