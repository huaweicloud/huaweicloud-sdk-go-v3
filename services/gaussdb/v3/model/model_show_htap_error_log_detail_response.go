package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHtapErrorLogDetailResponse Response Object
type ShowHtapErrorLogDetailResponse struct {

	// **参数解释**： 错误日志列表。  **约束限制**： 不涉及。
	ErrorLogList   *[]HtapErrorLogDetailResponseErrorLogList `json:"error_log_list,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowHtapErrorLogDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapErrorLogDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowHtapErrorLogDetailResponse", string(data)}, " ")
}
