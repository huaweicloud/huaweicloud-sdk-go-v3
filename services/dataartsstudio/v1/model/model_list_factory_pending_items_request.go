package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryPendingItemsRequest Request Object
type ListFactoryPendingItemsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 提交人。
	SubmitUserName *string `json:"submit_user_name,omitempty"`

	// 任务名称。
	ItemName *string `json:"item_name,omitempty"`

	// 变更类型。
	UpdateType *string `json:"update_type,omitempty"`

	// 任务类型。
	TaskType *string `json:"task_type,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 分页返回结果，指定每页最大记录数。范围[1,100] 默认值：10
	Limit *int32 `json:"limit,omitempty"`

	// 分页列表的页数，默认值为0。取值范围大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFactoryPendingItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryPendingItemsRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryPendingItemsRequest", string(data)}, " ")
}
