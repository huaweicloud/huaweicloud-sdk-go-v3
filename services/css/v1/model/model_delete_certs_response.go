package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertsResponse Response Object
type DeleteCertsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertsResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertsResponse", string(data)}, " ")
}
