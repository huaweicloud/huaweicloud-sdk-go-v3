package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEncryptTaskResponse struct {

	// 任务列表
	TaskArray *[]EachEncryptRsp `json:"task_array,omitempty" xml:"task_array"`

	// 是否截断
	IsTruncated *int32 `json:"is_truncated,omitempty" xml:"is_truncated"`

	// 查询结果数量
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEncryptTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncryptTaskResponse struct{}"
	}

	return strings.Join([]string{"ListEncryptTaskResponse", string(data)}, " ")
}
