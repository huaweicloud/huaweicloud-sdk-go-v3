package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObjectMappingResponse Response Object
type ShowObjectMappingResponse struct {

	// 总数。
	Count *int64 `json:"count,omitempty"`

	// 同步映射数据列表。
	ObjectMappingList *[]DbObjectInfo `json:"object_mapping_list,omitempty"`
	HttpStatusCode    int             `json:"-"`
}

func (o ShowObjectMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectMappingResponse struct{}"
	}

	return strings.Join([]string{"ShowObjectMappingResponse", string(data)}, " ")
}
