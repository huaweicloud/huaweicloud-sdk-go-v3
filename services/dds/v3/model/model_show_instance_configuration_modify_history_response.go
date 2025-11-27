package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceConfigurationModifyHistoryResponse Response Object
type ShowInstanceConfigurationModifyHistoryResponse struct {

	// 实例参数的修改历史列表。
	Histories *[]ConfigurationModifyHistoryInfo `json:"histories,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceConfigurationModifyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationModifyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationModifyHistoryResponse", string(data)}, " ")
}
