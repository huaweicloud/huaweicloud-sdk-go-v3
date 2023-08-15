package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppResponse Response Object
type PublishAppResponse struct {

	// 资产id
	Id *string `json:"id,omitempty"`

	// 资产版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppResponse struct{}"
	}

	return strings.Join([]string{"PublishAppResponse", string(data)}, " ")
}
