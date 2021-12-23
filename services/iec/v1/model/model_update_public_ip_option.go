package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新弹性公网IP参数
type UpdatePublicIpOption struct {
	// vip对应的port的ID可为空。当为空的时候，代表解绑原有eip的关系。不为空时，代表绑定eip和vip。

	PortId string `json:"port_id"`
}

func (o UpdatePublicIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicIpOption struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpOption", string(data)}, " ")
}
