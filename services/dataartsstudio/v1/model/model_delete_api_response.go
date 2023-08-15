package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApiResponse Response Object
type DeleteApiResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiResponse struct{}"
	}

	return strings.Join([]string{"DeleteApiResponse", string(data)}, " ")
}
