package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirtyDataResponse Response Object
type ShowDirtyDataResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 异常数据列表。
	DirtyDataList  *[]DirtyData `json:"dirty_data_list,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDirtyDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirtyDataResponse struct{}"
	}

	return strings.Join([]string{"ShowDirtyDataResponse", string(data)}, " ")
}
