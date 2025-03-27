package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSpeedResponse Response Object
type UpdateSpeedResponse struct {

	// 设置迁移限速规则成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateSpeedResponse", string(data)}, " ")
}
