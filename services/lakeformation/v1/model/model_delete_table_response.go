package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableResponse Response Object
type DeleteTableResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableResponse", string(data)}, " ")
}
