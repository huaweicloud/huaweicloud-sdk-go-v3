package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishVersionVoSearchResultDataValue 返回的数据信息。
type PublishVersionVoSearchResultDataValue struct {

	// 查询到的版本值对象（PublishVersionVO）数组。
	Records *[]PublishVersionVo `json:"records,omitempty"`

	// 符合搜索条件的记录总数。
	Total *int32 `json:"total,omitempty"`
}

func (o PublishVersionVoSearchResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishVersionVoSearchResultDataValue struct{}"
	}

	return strings.Join([]string{"PublishVersionVoSearchResultDataValue", string(data)}, " ")
}
