package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePostalResponse Response Object
type UpdatePostalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePostalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostalResponse struct{}"
	}

	return strings.Join([]string{"UpdatePostalResponse", string(data)}, " ")
}
