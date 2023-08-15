package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceResponse Response Object
type DeleteServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceResponse", string(data)}, " ")
}
