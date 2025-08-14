package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportResourceResponse Response Object
type ImportResourceResponse struct {

	// 返回结果
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportResourceResponse struct{}"
	}

	return strings.Join([]string{"ImportResourceResponse", string(data)}, " ")
}
