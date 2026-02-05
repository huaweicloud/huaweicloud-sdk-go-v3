package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIndicatorsResponse Response Object
type ImportIndicatorsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *ImportIndicatorsResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ImportIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ImportIndicatorsResponse", string(data)}, " ")
}
