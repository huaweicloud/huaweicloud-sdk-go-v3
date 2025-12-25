package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTagValueResponse Response Object
type UpdateTagValueResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTagValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTagValueResponse struct{}"
	}

	return strings.Join([]string{"UpdateTagValueResponse", string(data)}, " ")
}
