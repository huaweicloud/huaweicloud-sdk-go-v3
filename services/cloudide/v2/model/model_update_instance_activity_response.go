package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateInstanceActivityResponse struct {
	Result *ExpireVo `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceActivityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceActivityResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceActivityResponse", string(data)}, " ")
}
