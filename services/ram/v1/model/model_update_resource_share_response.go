package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceShareResponse Response Object
type UpdateResourceShareResponse struct {
	ResourceShare  *ResourceShare `json:"resource_share,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateResourceShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceShareResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceShareResponse", string(data)}, " ")
}
