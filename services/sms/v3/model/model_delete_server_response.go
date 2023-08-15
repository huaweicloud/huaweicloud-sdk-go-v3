package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerResponse Response Object
type DeleteServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerResponse", string(data)}, " ")
}
