package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeQuotaNewResponse Response Object
type ChangeQuotaNewResponse struct {

	// 实例列表
	Instances *[]DasCommonInstanceDto `json:"instances,omitempty"`

	// 总记录数
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ChangeQuotaNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeQuotaNewResponse struct{}"
	}

	return strings.Join([]string{"ChangeQuotaNewResponse", string(data)}, " ")
}
