package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDocumentsRequest Request Object
type ListDocumentsRequest struct {

	// 分页参数：每页返回记录个数限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 作业名（模糊）
	NameLike *string `json:"name_like,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 执行的作业类型 枚举：PUBLIC/NORMAL 默认NORMAL
	DocumentType *string `json:"document_type,omitempty"`
}

func (o ListDocumentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDocumentsRequest struct{}"
	}

	return strings.Join([]string{"ListDocumentsRequest", string(data)}, " ")
}
