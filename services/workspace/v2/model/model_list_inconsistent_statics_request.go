package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInconsistentStaticsRequest Request Object
type ListInconsistentStaticsRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`
}

func (o ListInconsistentStaticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInconsistentStaticsRequest struct{}"
	}

	return strings.Join([]string{"ListInconsistentStaticsRequest", string(data)}, " ")
}
