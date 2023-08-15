package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunQueueActionReq struct {

	// 执行动作：restart：重启scale_out：扩容scale_in：缩容，目前只支持restart、scale_out、scale_in。
	Action string `json:"action"`

	// 是否强制重启，action为restart时可选择配置，默认为false。
	Force *bool `json:"force,omitempty"`

	// 队列的实际CU。
	CuCount *int32 `json:"cu_count,omitempty"`
}

func (o RunQueueActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueueActionReq struct{}"
	}

	return strings.Join([]string{"RunQueueActionReq", string(data)}, " ")
}
