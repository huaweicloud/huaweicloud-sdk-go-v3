package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataconnectionResponse Response Object
type DeleteDataconnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDataconnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataconnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataconnectionResponse", string(data)}, " ")
}
