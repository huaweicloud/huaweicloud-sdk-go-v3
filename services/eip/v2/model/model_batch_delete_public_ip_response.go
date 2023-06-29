package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicIpResponse Response Object
type BatchDeletePublicIpResponse struct {

	// job_id列表，此job信息不保存在数据库中，不能同过组件查询到。
	JobIds         *[]string `json:"job_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeletePublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicIpResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicIpResponse", string(data)}, " ")
}
