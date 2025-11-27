package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationGroupResponse Response Object
type UpdateApplicationGroupResponse struct {

	// 分组id。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateApplicationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationGroupResponse", string(data)}, " ")
}
