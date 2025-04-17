package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRefreshResultResponse Response Object
type ShowRefreshResultResponse struct {

	// 刷新任务结果
	RefreshResults *[]RefreshResult `json:"refresh_results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowRefreshResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefreshResultResponse struct{}"
	}

	return strings.Join([]string{"ShowRefreshResultResponse", string(data)}, " ")
}
