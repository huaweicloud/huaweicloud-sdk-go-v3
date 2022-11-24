package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConfigurationAppliedHistoryResponse struct {

	// 参数模板应用历史列表
	Histories      *[]ApplyHistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowConfigurationAppliedHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAppliedHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAppliedHistoryResponse", string(data)}, " ")
}
