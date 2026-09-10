package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionHostReq 修改主机信息的请求body体。
type UpdateConnectionHostReq struct {

	// 用户自定义主机信息，最大支持2万条记录。内容填空表示清除所有已配置的主机信息。
	Hosts []ConnectionsHost `json:"hosts"`
}

func (o UpdateConnectionHostReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionHostReq struct{}"
	}

	return strings.Join([]string{"UpdateConnectionHostReq", string(data)}, " ")
}
