package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTbSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTbSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTbSessionResponse struct{}"
	}

	return strings.Join([]string{"DeleteTbSessionResponse", string(data)}, " ")
}
