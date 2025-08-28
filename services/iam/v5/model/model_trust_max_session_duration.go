package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrustMaxSessionDuration 信任委托最大会话时长，默认为3600秒，取值范围为[3600,43200]。
type TrustMaxSessionDuration struct {
}

func (o TrustMaxSessionDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustMaxSessionDuration struct{}"
	}

	return strings.Join([]string{"TrustMaxSessionDuration", string(data)}, " ")
}
