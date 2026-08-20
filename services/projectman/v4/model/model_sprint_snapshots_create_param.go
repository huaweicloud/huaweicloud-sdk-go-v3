package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SprintSnapshotsCreateParam 创建迭代快照请求对象
type SprintSnapshotsCreateParam struct {

	// 快照标题。
	Title string `json:"title"`

	// 计划唯一ID。可以通过IPD项目计划管理章节中发布/迭代计划列表查询接口获取，响应消息体中的id字段的值就是计划ID。
	IssueId string `json:"issue_id"`

	// 计划类别。可以通过IPD项目计划管理章节中发布/迭代计划列表查询接口获取，响应消息体中的category字段的值就是计划类别。
	Category string `json:"category"`
}

func (o SprintSnapshotsCreateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SprintSnapshotsCreateParam struct{}"
	}

	return strings.Join([]string{"SprintSnapshotsCreateParam", string(data)}, " ")
}
