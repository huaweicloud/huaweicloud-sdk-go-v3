package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageResponse Response Object
type UpdateImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageResponse struct{}"
	}

	return strings.Join([]string{"UpdateImageResponse", string(data)}, " ")
}
