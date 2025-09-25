package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageFilesResponse Response Object
type ListImageFilesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
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
