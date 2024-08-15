package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceCompliantResponse Response Object
type ListInstanceCompliantResponse struct {

	// 总条数
	Count *int64 `json:"count,omitempty"`

	// 节点合规报告
	InstanceCompliant *[]InstanceCompliant `json:"instance_compliant,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ListInstanceCompliantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceCompliantResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceCompliantResponse", string(data)}, " ")
}
