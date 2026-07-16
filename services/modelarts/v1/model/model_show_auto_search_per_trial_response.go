package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchPerTrialResponse Response Object
type ShowAutoSearchPerTrialResponse struct {

	// 超参搜索某个trial结果的字段信息。
	Header *[]string `json:"header,omitempty"`

	// 超参搜索某个trial结果的每条数据列表。
	Data           *[][]string `json:"data,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowAutoSearchPerTrialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchPerTrialResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchPerTrialResponse", string(data)}, " ")
}
