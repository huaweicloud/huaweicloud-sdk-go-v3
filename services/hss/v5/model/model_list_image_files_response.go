package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageFilesResponse Response Object
type ListImageFilesResponse struct {

	// **参数解释**: 总数 **取值范围**: 字符长度1-65535
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]ImageFileInfo `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListImageFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageFilesResponse struct{}"
	}

	return strings.Join([]string{"ListImageFilesResponse", string(data)}, " ")
}
