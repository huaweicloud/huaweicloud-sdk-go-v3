package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubmissionsResponse Response Object
type ShowSubmissionsResponse struct {

	// 作业运行信息，详见submissions参数说明。
	Submissions *[]Submission `json:"submissions,omitempty"`

	// 查询该作业总的历史记录数。
	Total *int32 `json:"total,omitempty"`

	// 查询作业记录时，分页数。
	PageNo *int32 `json:"page_no,omitempty"`

	// 分页查询，每页返回的记录数。默认值：10。
	PageSize       *int32 `json:"page_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSubmissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubmissionsResponse struct{}"
	}

	return strings.Join([]string{"ShowSubmissionsResponse", string(data)}, " ")
}
