package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEcnWithErResponse Response Object
type DeleteEcnWithErResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEcnWithErResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEcnWithErResponse struct{}"
	}

	return strings.Join([]string{"DeleteEcnWithErResponse", string(data)}, " ")
}
