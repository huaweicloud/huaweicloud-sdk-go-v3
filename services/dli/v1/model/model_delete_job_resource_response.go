package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobResourceResponse Response Object
type DeleteJobResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteJobResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobResourceResponse", string(data)}, " ")
}
