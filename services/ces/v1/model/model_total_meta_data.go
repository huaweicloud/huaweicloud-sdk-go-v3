package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TotalMetaData **参数解释**： 查询结果元数据统计个数。
type TotalMetaData struct {

	// **参数解释**： 总条数。 **取值范围**： 在[0,2147483647]区间内。
	Total *int32 `json:"total,omitempty"`
}

func (o TotalMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalMetaData struct{}"
	}

	return strings.Join([]string{"TotalMetaData", string(data)}, " ")
}
