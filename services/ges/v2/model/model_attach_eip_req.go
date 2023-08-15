package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachEipReq 绑定EIP请求体
type AttachEipReq struct {

	// 弹性公网IP的ID。
	EipId string `json:"eip_id"`
}

func (o AttachEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipReq struct{}"
	}

	return strings.Join([]string{"AttachEipReq", string(data)}, " ")
}
