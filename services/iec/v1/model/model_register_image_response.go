package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterImageResponse Response Object
type RegisterImageResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId string `json:"job_id"`

	// 创建镜像的ID。
	ImageId        *string `json:"image_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageResponse struct{}"
	}

	return strings.Join([]string{"RegisterImageResponse", string(data)}, " ")
}
