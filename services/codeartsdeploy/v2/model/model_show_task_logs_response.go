package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskLogsResponse Response Object
type ShowTaskLogsResponse struct {
	Result *QueryAppLogResult `json:"result,omitempty"`

	// 查询部署日志状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskLogsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskLogsResponse", string(data)}, " ")
}
