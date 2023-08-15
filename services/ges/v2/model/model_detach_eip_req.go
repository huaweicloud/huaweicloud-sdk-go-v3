package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachEipReq 解绑EIP请求体
type DetachEipReq struct {

	// 弹性公网IP的ID。
	EipId string `json:"eip_id"`
}

func (o DetachEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachEipReq struct{}"
	}

	return strings.Join([]string{"DetachEipReq", string(data)}, " ")
}
