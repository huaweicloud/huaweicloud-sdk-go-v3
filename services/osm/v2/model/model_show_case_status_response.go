package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCaseStatusResponse struct {
	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈

	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCaseStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCaseStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowCaseStatusResponse", string(data)}, " ")
}
