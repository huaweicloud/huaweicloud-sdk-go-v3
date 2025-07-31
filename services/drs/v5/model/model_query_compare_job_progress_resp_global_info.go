package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryCompareJobProgressRespGlobalInfo 全局对比信息。
type QueryCompareJobProgressRespGlobalInfo struct {

	// 全局对比速率。
	SrcSpeed *string `json:"src_speed,omitempty"`
}

func (o QueryCompareJobProgressRespGlobalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCompareJobProgressRespGlobalInfo struct{}"
	}

	return strings.Join([]string{"QueryCompareJobProgressRespGlobalInfo", string(data)}, " ")
}
