package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询结果元数据信息，包括分页信息等。
type MetaData struct {

	// 当前返回结果条数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 下一个开始的标记，用于分页，null表示无更多数据。
	Start *int64 `json:"start,omitempty" xml:"start"`

	// 总条数。
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o MetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaData struct{}"
	}

	return strings.Join([]string{"MetaData", string(data)}, " ")
}
