package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEntityAttributeResponse Response Object
type UpdateEntityAttributeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEntityAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEntityAttributeResponse struct{}"
	}

	return strings.Join([]string{"UpdateEntityAttributeResponse", string(data)}, " ")
}
