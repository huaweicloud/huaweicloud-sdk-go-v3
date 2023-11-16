package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFactoryJobNameResponse Response Object
type UpdateFactoryJobNameResponse struct {

	// 取值为true或者false
	IsSuccess *bool `json:"is_success,omitempty"`

	// 200表示成功返回
	StatusCode     *int32 `json:"status_code,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateFactoryJobNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFactoryJobNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateFactoryJobNameResponse", string(data)}, " ")
}
