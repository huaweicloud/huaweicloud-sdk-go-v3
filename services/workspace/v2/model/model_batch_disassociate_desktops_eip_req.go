package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateDesktopsEipReq 批量桌面解绑EIP请求体。
type BatchDisassociateDesktopsEipReq struct {

	// 待解绑EIP的桌面id。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`
}

func (o BatchDisassociateDesktopsEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateDesktopsEipReq struct{}"
	}

	return strings.Join([]string{"BatchDisassociateDesktopsEipReq", string(data)}, " ")
}
