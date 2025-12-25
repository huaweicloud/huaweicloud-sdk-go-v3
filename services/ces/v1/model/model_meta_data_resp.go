package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaDataResp 查询结果元数据信息，包括分页信息等。
type MetaDataResp struct {

	// **参数解释**： 当前返回结果条数。 **取值范围**: 0 - 2147483647
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 结果总条数。 **取值范围**： 0 - 2147483647
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 下一个开始的标记，用于分页。如本次查询10条数据，第十条alarm_id为al1441967036681YkazZ0deN，下次start配置为al1441967036681YkazZ0deN可从该alarm_id开始查询。 **取值范围**： 1 - 9999
	Marker *string `json:"marker,omitempty"`
}

func (o MetaDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDataResp struct{}"
	}

	return strings.Join([]string{"MetaDataResp", string(data)}, " ")
}
