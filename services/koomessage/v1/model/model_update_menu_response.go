package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateMenuResponse struct {
	Data           *UpdateMenuResponseModel `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateMenuResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMenuResponse struct{}"
	}

	return strings.Join([]string{"UpdateMenuResponse", string(data)}, " ")
}
