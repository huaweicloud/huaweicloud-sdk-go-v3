package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClientCaResponse Response Object
type DeleteClientCaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClientCaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClientCaResponse struct{}"
	}

	return strings.Join([]string{"DeleteClientCaResponse", string(data)}, " ")
}
