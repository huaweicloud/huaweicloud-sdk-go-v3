package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAimMsgAppResponse Response Object
type UpdateAimMsgAppResponse struct {

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用ID。
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAimMsgAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAimMsgAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAimMsgAppResponse", string(data)}, " ")
}
