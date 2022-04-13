package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateWebinarResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWebinarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebinarResponse struct{}"
	}

	return strings.Join([]string{"UpdateWebinarResponse", string(data)}, " ")
}
