package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipJob 批量创建返回的信息对象
type BatchCreateGlobalEipJob struct {

	// 全域弹性公网IP的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：全域弹性公网IP名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 请求完成的job id
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchCreateGlobalEipJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipJob struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipJob", string(data)}, " ")
}
