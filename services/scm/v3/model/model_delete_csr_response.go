package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCsrResponse Response Object
type DeleteCsrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCsrResponse struct{}"
	}

	return strings.Join([]string{"DeleteCsrResponse", string(data)}, " ")
}
