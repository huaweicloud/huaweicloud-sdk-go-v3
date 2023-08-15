package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishDataResponse Response Object
type PublishDataResponse struct {

	// 资产id
	Id *string `json:"id,omitempty"`

	// 资产版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishDataResponse struct{}"
	}

	return strings.Join([]string{"PublishDataResponse", string(data)}, " ")
}
