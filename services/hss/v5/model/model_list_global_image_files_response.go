package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalImageFilesResponse Response Object
type ListGlobalImageFilesResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值65535
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]ImageFileResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListGlobalImageFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalImageFilesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalImageFilesResponse", string(data)}, " ")
}
