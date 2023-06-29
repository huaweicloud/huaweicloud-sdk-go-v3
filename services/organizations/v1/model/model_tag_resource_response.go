package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceResponse Response Object
type TagResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o TagResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceResponse struct{}"
	}

	return strings.Join([]string{"TagResourceResponse", string(data)}, " ")
}
