package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceByIdResponse Response Object
type DeleteInstanceByIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceByIdResponse", string(data)}, " ")
}
