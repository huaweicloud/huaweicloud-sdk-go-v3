package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIsolatedFileResponse Response Object
type ListIsolatedFileResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 已隔离文件详情
	DataList       *[]IsolatedFileResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListIsolatedFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIsolatedFileResponse struct{}"
	}

	return strings.Join([]string{"ListIsolatedFileResponse", string(data)}, " ")
}
