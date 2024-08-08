package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachServerAppInfo 分发软件信息。
type AttachServerAppInfo struct {

	// 分发软件版本ID。
	Id *string `json:"id,omitempty"`

	// 分发软件ID。
	AppId *string `json:"app_id,omitempty"`
}

func (o AttachServerAppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerAppInfo struct{}"
	}

	return strings.Join([]string{"AttachServerAppInfo", string(data)}, " ")
}
