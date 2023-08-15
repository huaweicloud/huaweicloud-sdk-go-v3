package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProcessResponse Response Object
type DeleteProcessResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProcessResponse struct{}"
	}

	return strings.Join([]string{"DeleteProcessResponse", string(data)}, " ")
}
