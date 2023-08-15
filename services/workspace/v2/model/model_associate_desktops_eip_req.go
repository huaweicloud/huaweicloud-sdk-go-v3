package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateDesktopsEipReq 桌面绑定EIP请求体。
type AssociateDesktopsEipReq struct {

	// 桌面绑定的Eip的id。
	EipId string `json:"eip_id"`

	// 桌面id。
	DesktopId string `json:"desktop_id"`
}

func (o AssociateDesktopsEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateDesktopsEipReq struct{}"
	}

	return strings.Join([]string{"AssociateDesktopsEipReq", string(data)}, " ")
}
