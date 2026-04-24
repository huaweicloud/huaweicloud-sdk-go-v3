package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcVncRequestBody 更新专项自主维护请求体
type UpdateDcVncRequestBody struct {

	// - default: 自动开启 - close: 关闭
	DcVncIp string `json:"dc_vnc_ip"`

	// 中心可用区的子网id，当dc_vnc_ip为default，且站点属于边缘小站时必传
	CenterSubnetId *string `json:"center_subnet_id,omitempty"`
}

func (o UpdateDcVncRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcVncRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDcVncRequestBody", string(data)}, " ")
}
