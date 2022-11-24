package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceResponse", string(data)}, " ")
}
