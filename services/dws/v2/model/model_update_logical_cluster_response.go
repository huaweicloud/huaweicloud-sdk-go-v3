package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterResponse Response Object
type UpdateLogicalClusterResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLogicalClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterResponse", string(data)}, " ")
}
