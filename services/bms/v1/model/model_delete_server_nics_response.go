package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerNicsResponse Response Object
type DeleteServerNicsResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServerNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerNicsResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerNicsResponse", string(data)}, " ")
}
