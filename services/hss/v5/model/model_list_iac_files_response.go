package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFilesResponse Response Object
type ListIacFilesResponse struct {

	// 文件总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// IAC文件列表
	DataList       *[]ListIacFilesResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListIacFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFilesResponse struct{}"
	}

	return strings.Join([]string{"ListIacFilesResponse", string(data)}, " ")
}
