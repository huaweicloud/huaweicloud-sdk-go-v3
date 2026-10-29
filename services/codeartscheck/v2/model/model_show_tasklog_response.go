package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTasklogResponse Response Object
type ShowTasklogResponse struct {
	ParamInfo *ParamInfo `json:"param_info,omitempty"`

	// **参数解释：** 失败日志详情。
	LogInfo        *[]LogInfo `json:"log_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowTasklogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasklogResponse struct{}"
	}

	return strings.Join([]string{"ShowTasklogResponse", string(data)}, " ")
}
