package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGraph2Response Response Object
type DeleteGraph2Response struct {

	// 删除图任务ID。请求失败时字段为空。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGraph2Response struct{}"
	}

	return strings.Join([]string{"DeleteGraph2Response", string(data)}, " ")
}
