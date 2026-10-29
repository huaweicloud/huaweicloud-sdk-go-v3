package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRecordResponse Response Object
type CheckRecordResponse struct {

	// **参数解释：** 每次扫描的各个问题级别的问题数量。
	Data *[]CheckRecordDataInfo `json:"data,omitempty"`

	// **参数解释：** 扫描次数总数。 **取值范围：** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecordResponse struct{}"
	}

	return strings.Join([]string{"CheckRecordResponse", string(data)}, " ")
}
