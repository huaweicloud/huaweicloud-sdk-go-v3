package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCheckPointResponse Response Object
type UpdateCheckPointResponse struct {

	// 状态码
	ErrorCode *string `json:"errorCode,omitempty"`

	// 状态信息
	ErrorMessage   *string `json:"errorMessage,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCheckPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCheckPointResponse struct{}"
	}

	return strings.Join([]string{"UpdateCheckPointResponse", string(data)}, " ")
}
