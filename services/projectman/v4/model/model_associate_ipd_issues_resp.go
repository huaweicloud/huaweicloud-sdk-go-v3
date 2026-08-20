package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIpdIssuesResp 关联/取消关联响应对象
type AssociateIpdIssuesResp struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 关联失败的原因。
	Message *string `json:"message,omitempty"`

	// 关联工作项的响应结果。
	Result map[string][]AssociateRespDetail `json:"result,omitempty"`
}

func (o AssociateIpdIssuesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIpdIssuesResp struct{}"
	}

	return strings.Join([]string{"AssociateIpdIssuesResp", string(data)}, " ")
}
