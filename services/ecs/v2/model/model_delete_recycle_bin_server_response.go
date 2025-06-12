package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecycleBinServerResponse Response Object
type DeleteRecycleBinServerResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRecycleBinServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecycleBinServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecycleBinServerResponse", string(data)}, " ")
}
