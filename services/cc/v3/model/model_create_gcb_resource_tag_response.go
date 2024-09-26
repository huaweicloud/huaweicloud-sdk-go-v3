package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGcbResourceTagResponse Response Object
type CreateGcbResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateGcbResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGcbResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateGcbResourceTagResponse", string(data)}, " ")
}
