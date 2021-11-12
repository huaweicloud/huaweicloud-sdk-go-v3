package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorDetailInfo struct {
	// 批量操作失败的资源的详情信息

	Failed *[]ErrorCaseInfoBean `json:"failed,omitempty"`
}

func (o ErrorDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorDetailInfo struct{}"
	}

	return strings.Join([]string{"ErrorDetailInfo", string(data)}, " ")
}
