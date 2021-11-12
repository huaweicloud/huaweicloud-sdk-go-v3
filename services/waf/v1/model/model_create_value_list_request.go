package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateValueListRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateValueListRequestBody `json:"body,omitempty"`
}

func (o CreateValueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateValueListRequest struct{}"
	}

	return strings.Join([]string{"CreateValueListRequest", string(data)}, " ")
}
