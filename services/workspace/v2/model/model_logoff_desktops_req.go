package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffDesktopsReq 注销桌面的请求。
type LogoffDesktopsReq struct {

	// 计算机id列表。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 下发注销桌面任务时，给用户发送的提示信息。
	Message *string `json:"message,omitempty"`
}

func (o LogoffDesktopsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffDesktopsReq struct{}"
	}

	return strings.Join([]string{"LogoffDesktopsReq", string(data)}, " ")
}
