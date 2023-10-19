package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessClientResponse Response Object
type UpdateAccessClientResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAccessClientResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessClientResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessClientResponse", string(data)}, " ")
}
