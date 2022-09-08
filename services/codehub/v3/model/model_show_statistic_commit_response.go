package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStatisticCommitResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *CommitStatistic `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStatisticCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticCommitResponse", string(data)}, " ")
}
