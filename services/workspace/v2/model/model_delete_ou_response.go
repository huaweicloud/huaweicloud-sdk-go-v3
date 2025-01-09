package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOuResponse Response Object
type DeleteOuResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOuResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOuResponse struct{}"
	}

	return strings.Join([]string{"DeleteOuResponse", string(data)}, " ")
}
