package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostMassageReq 修改主机信息的请求body体。
type UpdateHostMassageReq struct {

	// 用户自定义主机信息，最大支持2万条记录。内容填空表示清除所有已配置的主机信息。
	Hosts []EnhancedConnectionsHost `json:"hosts"`
}

func (o UpdateHostMassageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostMassageReq struct{}"
	}

	return strings.Join([]string{"UpdateHostMassageReq", string(data)}, " ")
}
