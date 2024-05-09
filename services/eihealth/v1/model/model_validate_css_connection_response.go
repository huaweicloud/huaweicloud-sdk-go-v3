package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateCssConnectionResponse Response Object
type ValidateCssConnectionResponse struct {

	// 是否连通
	Connected *bool `json:"connected,omitempty"`

	// 失败原因
	Reason         *string `json:"reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateCssConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateCssConnectionResponse struct{}"
	}

	return strings.Join([]string{"ValidateCssConnectionResponse", string(data)}, " ")
}
