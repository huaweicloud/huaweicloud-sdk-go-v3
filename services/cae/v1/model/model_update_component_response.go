package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComponentResponse Response Object
type UpdateComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentResponse struct{}"
	}

	return strings.Join([]string{"UpdateComponentResponse", string(data)}, " ")
}
