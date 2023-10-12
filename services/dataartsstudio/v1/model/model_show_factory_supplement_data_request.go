package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactorySupplementDataRequest Request Object
type ShowFactorySupplementDataRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 排序字段:desc：创建时间按照降序展示asc ：创建时间按照升序展示默认值：desc
	Sort *string `json:"sort,omitempty"`

	// 分页列表的起始页，默认值为0。取值范围大于等于0。
	Page *string `json:"page,omitempty"`

	// 分页返回结果，指定每页最大记录数。默认值：10
	Size *string `json:"size,omitempty"`

	// 补数据名称
	Name *string `json:"name,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 实例状态：SUCCESS：成功RUNNING ：运行中CANCLE：取消
	Status *string `json:"status,omitempty"`

	// 查询作业的开始日期 13位时间戳
	StartDate *string `json:"start_date,omitempty"`

	// 查询作业的结束日期 13位时间戳
	EndDate *string `json:"end_date,omitempty"`
}

func (o ShowFactorySupplementDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactorySupplementDataRequest struct{}"
	}

	return strings.Join([]string{"ShowFactorySupplementDataRequest", string(data)}, " ")
}
