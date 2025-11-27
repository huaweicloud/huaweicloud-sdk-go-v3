package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigSetRequest Request Object
type UpdateConfigSetRequest struct {

	// 配置集合id
	Configsetid string `json:"configsetid"`

	Body *UpdateConfigSetRequestBody `json:"body,omitempty"`
}

func (o UpdateConfigSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigSetRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigSetRequest", string(data)}, " ")
}
