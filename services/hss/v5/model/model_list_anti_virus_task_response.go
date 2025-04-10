package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusTaskResponse Response Object
type ListAntiVirusTaskResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// data list
	DataList       *[]AntiVirusTaskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListAntiVirusTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAntiVirusTaskResponse", string(data)}, " ")
}
