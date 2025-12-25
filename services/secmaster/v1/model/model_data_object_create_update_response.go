package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataObjectCreateUpdateResponse 数据对象创建/更新返回体
type DataObjectCreateUpdateResponse struct {

	// 唯一标识ID
	Id *string `json:"id,omitempty"`

	// 唯一事件标识ID
	EventId *string `json:"event_id,omitempty"`
}

func (o DataObjectCreateUpdateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectCreateUpdateResponse struct{}"
	}

	return strings.Join([]string{"DataObjectCreateUpdateResponse", string(data)}, " ")
}
