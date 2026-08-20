package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopPathSummary topPath详情数据
type TopPathSummary struct {

	// **参数解释：** top100 path访问路径 **取值范围：** 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释：** top100 path访问次数 **取值范围：** 不涉及
	Value *int64 `json:"value,omitempty"`
}

func (o TopPathSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopPathSummary struct{}"
	}

	return strings.Join([]string{"TopPathSummary", string(data)}, " ")
}
