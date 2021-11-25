package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteShareResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareResponse struct{}"
	}

	return strings.Join([]string{"DeleteShareResponse", string(data)}, " ")
}
