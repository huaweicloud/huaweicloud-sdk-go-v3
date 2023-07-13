package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveLayoutResponse Response Object
type SaveLayoutResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SaveLayoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveLayoutResponse struct{}"
	}

	return strings.Join([]string{"SaveLayoutResponse", string(data)}, " ")
}
