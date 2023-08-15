package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildImageResponse Response Object
type RebuildImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RebuildImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildImageResponse struct{}"
	}

	return strings.Join([]string{"RebuildImageResponse", string(data)}, " ")
}
