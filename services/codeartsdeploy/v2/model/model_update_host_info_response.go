package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostInfoResponse Response Object
type UpdateHostInfoResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostInfoResponse", string(data)}, " ")
}
