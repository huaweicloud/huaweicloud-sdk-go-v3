package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelServiceResponse Response Object
type UpdateModelServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateModelServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateModelServiceResponse", string(data)}, " ")
}
