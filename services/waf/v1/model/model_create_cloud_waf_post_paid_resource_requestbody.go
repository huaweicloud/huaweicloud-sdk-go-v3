package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 开通云模式按需请求体
type CreateCloudWafPostPaidResourceRequestbody struct {

	// 租户所在的站点，hec-hk：华为云国际站
	ConsoleArea string `json:"console_area"`
}

func (o CreateCloudWafPostPaidResourceRequestbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudWafPostPaidResourceRequestbody struct{}"
	}

	return strings.Join([]string{"CreateCloudWafPostPaidResourceRequestbody", string(data)}, " ")
}
