package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticCommitV3Response Response Object
type ShowStatisticCommitV3Response struct {
	Error *Error `json:"error,omitempty"`

	Result *CommitStatistic `json:"result,omitempty"`

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
