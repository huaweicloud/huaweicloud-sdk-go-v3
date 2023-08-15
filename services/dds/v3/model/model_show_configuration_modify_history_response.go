package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationModifyHistoryResponse Response Object
type ShowConfigurationModifyHistoryResponse struct {

	// 参数模板的修改历史列表。
	Histories      *[]HistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowConfigurationModifyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationModifyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationModifyHistoryResponse", string(data)}, " ")
}
