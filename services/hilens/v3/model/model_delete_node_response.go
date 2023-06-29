package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeResponse Response Object
type DeleteNodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteNodeResponse", string(data)}, " ")
}
