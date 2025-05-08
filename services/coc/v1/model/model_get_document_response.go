package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDocumentResponse Response Object
type GetDocumentResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业uuid
	DocumentId *string `json:"document_id,omitempty"`

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 作业内容，DSL语句
	Content *string `json:"content,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 作业版本，如v1
	Version *string `json:"version,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 修改人
	Modifier *string `json:"modifier,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 版本集合
	Versions *[]DocumentVersionVo `json:"versions,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDocumentResponse struct{}"
	}

	return strings.Join([]string{"GetDocumentResponse", string(data)}, " ")
}
