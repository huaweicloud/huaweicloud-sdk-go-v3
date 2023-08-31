package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSupplementdataResponse Response Object
type StopSupplementdataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopSupplementdataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSupplementdataResponse struct{}"
	}

	return strings.Join([]string{"StopSupplementdataResponse", string(data)}, " ")
}
