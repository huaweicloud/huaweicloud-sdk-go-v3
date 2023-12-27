package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchByChecksumRequest Request Object
type SearchByChecksumRequest struct {

	// checksum
	Checksum string `json:"checksum"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 仓库类型
	Format *string `json:"format,omitempty"`

	// 是否在项目中
	InProject *string `json:"in_project,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o SearchByChecksumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchByChecksumRequest struct{}"
	}

	return strings.Join([]string{"SearchByChecksumRequest", string(data)}, " ")
}
