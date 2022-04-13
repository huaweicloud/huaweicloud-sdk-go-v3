package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStatisticCommitV3Response struct {
	Error *Error `json:"error,omitempty"`
	// 代码增加和删除的行数

	Result *[]CommitStatistic `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStatisticCommitV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticCommitV3Response struct{}"
	}

	return strings.Join([]string{"ShowStatisticCommitV3Response", string(data)}, " ")
}
