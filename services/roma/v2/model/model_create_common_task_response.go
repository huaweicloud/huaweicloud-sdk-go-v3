package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCommonTaskResponse struct {
	Task *TaskBasicRsp `json:"task,omitempty"`
	// 参数类型为string，参数结构参照附录中“数据集成参数说明>RawFormDataResponse”章节

	TaskDetail     *string `json:"task_detail,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCommonTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommonTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateCommonTaskResponse", string(data)}, " ")
}
