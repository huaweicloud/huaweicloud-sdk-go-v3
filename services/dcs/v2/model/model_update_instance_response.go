package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceResponse Response Object
type UpdateInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceResponse", string(data)}, " ")
}
