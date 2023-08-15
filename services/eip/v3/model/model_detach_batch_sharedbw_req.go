package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachBatchSharedbwReq 共享带宽批量移出弹性公网IP请求参数
type DetachBatchSharedbwReq struct {

	// 共享带宽批量移出弹性公网IP请求对象
	Publicips *[]DetachBatchSharedbwReqPublicips `json:"publicips,omitempty"`
}

func (o DetachBatchSharedbwReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachBatchSharedbwReq struct{}"
	}

	return strings.Join([]string{"DetachBatchSharedbwReq", string(data)}, " ")
}
