package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssuesInfo 对外需求相关API Body信息
type IssuesInfo struct {

	// 关联需求
	WorkitemList []WorkItemInfo `json:"workitem_list"`
}

func (o IssuesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssuesInfo struct{}"
	}

	return strings.Join([]string{"IssuesInfo", string(data)}, " ")
}
