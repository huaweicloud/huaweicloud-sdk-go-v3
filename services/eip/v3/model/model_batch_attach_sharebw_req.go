package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachSharebwReq 共享带宽批量加入弹性公网IP请求参数
type BatchAttachSharebwReq struct {

	// - 功能说明：共享带宽数据 - 约束：共享带宽批量加入多个弹性公网IP时，请求参数publicips中的bandwidth_id必须为同一个共享带宽的id
	Publicips *[]BatchAttachSharebwDict `json:"publicips,omitempty"`
}

func (o BatchAttachSharebwReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharebwReq struct{}"
	}

	return strings.Join([]string{"BatchAttachSharebwReq", string(data)}, " ")
}
