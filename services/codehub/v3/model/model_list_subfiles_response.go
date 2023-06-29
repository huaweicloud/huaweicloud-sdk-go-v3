package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubfilesResponse Response Object
type ListSubfilesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *LogsTreeList `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubfilesResponse struct{}"
	}

	return strings.Join([]string{"ListSubfilesResponse", string(data)}, " ")
}
