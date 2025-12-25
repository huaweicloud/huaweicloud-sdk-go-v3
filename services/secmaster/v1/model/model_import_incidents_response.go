package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIncidentsResponse Response Object
type ImportIncidentsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *ImportIncidentsResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ImportIncidentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIncidentsResponse struct{}"
	}

	return strings.Join([]string{"ImportIncidentsResponse", string(data)}, " ")
}
