package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApiResponse Response Object
type UpdateApiResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiResponse struct{}"
	}

	return strings.Join([]string{"UpdateApiResponse", string(data)}, " ")
}
