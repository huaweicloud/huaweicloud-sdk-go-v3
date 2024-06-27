package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveIssuesInfo 从迭代中移除需求API Body信息
type RemoveIssuesInfo struct {

	// 关联需求
	WorkitemList []WorkItemInfo `json:"workitem_list"`

	// 是否删除需求关联的用例
	IsDeleteCase *bool `json:"is_delete_case,omitempty"`
}

func (o RemoveIssuesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIssuesInfo struct{}"
	}

	return strings.Join([]string{"RemoveIssuesInfo", string(data)}, " ")
}
