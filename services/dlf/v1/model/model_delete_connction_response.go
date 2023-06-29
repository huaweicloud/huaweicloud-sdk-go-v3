package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnctionResponse Response Object
type DeleteConnctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConnctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnctionResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnctionResponse", string(data)}, " ")
}
