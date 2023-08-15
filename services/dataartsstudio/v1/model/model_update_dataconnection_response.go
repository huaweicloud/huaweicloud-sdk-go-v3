package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataconnectionResponse Response Object
type UpdateDataconnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDataconnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataconnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataconnectionResponse", string(data)}, " ")
}
