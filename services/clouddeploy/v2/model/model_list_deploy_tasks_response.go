package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDeployTasksResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty" xml:"total_num"`

	// 返回结果
	Result         *[]TaskInfo `json:"result,omitempty" xml:"result"`
	HttpStatusCode int         `json:"-"`
}

func (o ListDeployTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployTasksResponse struct{}"
	}

	return strings.Join([]string{"ListDeployTasksResponse", string(data)}, " ")
}
