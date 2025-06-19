package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableTheJobResponse Response Object
type DisableTheJobResponse struct {

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableTheJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableTheJobResponse struct{}"
	}

	return strings.Join([]string{"DisableTheJobResponse", string(data)}, " ")
}
