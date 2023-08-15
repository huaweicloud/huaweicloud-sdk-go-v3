package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInitPasswordResponse Response Object
type UpdateInitPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInitPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInitPasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdateInitPasswordResponse", string(data)}, " ")
}
