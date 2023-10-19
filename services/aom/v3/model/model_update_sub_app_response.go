package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubAppResponse Response Object
type UpdateSubAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSubAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubAppResponse", string(data)}, " ")
}
