package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceShareResponse Response Object
type CreateResourceShareResponse struct {
	ResourceShare  *ResourceShare `json:"resource_share,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateResourceShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceShareResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceShareResponse", string(data)}, " ")
}
