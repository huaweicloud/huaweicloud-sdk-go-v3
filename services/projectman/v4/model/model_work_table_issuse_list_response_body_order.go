package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 优先级顺序信息
type WorkTableIssuseListResponseBodyOrder struct {

	// 优先级顺序id
	Id *int32 `json:"id,omitempty"`

	// 优先级顺序名称
	Name *string `json:"name,omitempty"`
}

func (o WorkTableIssuseListResponseBodyOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyOrder struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyOrder", string(data)}, " ")
}
