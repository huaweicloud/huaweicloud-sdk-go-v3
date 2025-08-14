package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationInstanceResponse Response Object
type DeleteApplicationInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationInstanceResponse", string(data)}, " ")
}
