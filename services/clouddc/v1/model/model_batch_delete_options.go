package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteOptions 批量删除请求参数
type BatchDeleteOptions struct {

	// instance id set
	InstanceIdSet *[]string `json:"instance_id_set,omitempty"`
}

func (o BatchDeleteOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOptions struct{}"
	}

	return strings.Join([]string{"BatchDeleteOptions", string(data)}, " ")
}
