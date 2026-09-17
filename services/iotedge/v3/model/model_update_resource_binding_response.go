package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceBindingResponse Response Object
type UpdateResourceBindingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateResourceBindingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceBindingResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceBindingResponse", string(data)}, " ")
}
