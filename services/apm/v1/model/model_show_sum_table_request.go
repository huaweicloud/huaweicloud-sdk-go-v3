package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSumTableRequest Request Object
type ShowSumTableRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *SumTableParam `json:"body,omitempty"`
}

func (o ShowSumTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSumTableRequest struct{}"
	}

	return strings.Join([]string{"ShowSumTableRequest", string(data)}, " ")
}
