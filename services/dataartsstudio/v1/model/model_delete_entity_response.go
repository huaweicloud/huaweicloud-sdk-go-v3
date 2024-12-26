package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEntityResponse Response Object
type DeleteEntityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEntityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEntityResponse struct{}"
	}

	return strings.Join([]string{"DeleteEntityResponse", string(data)}, " ")
}
