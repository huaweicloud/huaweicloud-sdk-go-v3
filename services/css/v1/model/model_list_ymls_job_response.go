package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListYmlsJobResponse Response Object
type ListYmlsJobResponse struct {

	// 历史修改配置列表。
	ConfigList *[]ConfigListRsp `json:"configList,omitempty"`

	// **参数解释**： 配置任务数量。 **取值范围**： 不涉及
	TotalSize      *int32 `json:"totalSize,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListYmlsJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListYmlsJobResponse struct{}"
	}

	return strings.Join([]string{"ListYmlsJobResponse", string(data)}, " ")
}
