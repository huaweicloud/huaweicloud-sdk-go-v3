package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationResponse Response Object
type UpdateApplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationResponse", string(data)}, " ")
}
