package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstanceListNewResponse Response Object
type ExportInstanceListNewResponse struct {

	// 导出的实例总数
	Total *int64 `json:"total,omitempty"`

	// 实例信息列表
	InstanceInfos  *[]ExportInstanceInfo `json:"instance_infos,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ExportInstanceListNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceListNewResponse struct{}"
	}

	return strings.Join([]string{"ExportInstanceListNewResponse", string(data)}, " ")
}
