package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModClusterProtectionEventRequest 修改安全事件
type ModClusterProtectionEventRequest struct {

	// 总数
	TotalNum int32 `json:"total_num"`

	// 操作类型，包含以下几种： - ignore：忽略 - handle：处理 - addWhiteImage：加白
	Opr string `json:"opr"`

	// 事件ID列表
	DataList []string `json:"data_list"`
}

func (o ModClusterProtectionEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModClusterProtectionEventRequest struct{}"
	}

	return strings.Join([]string{"ModClusterProtectionEventRequest", string(data)}, " ")
}
