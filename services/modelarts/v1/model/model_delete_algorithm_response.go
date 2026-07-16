package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlgorithmResponse Response Object
type DeleteAlgorithmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlgorithmResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlgorithmResponse", string(data)}, " ")
}
