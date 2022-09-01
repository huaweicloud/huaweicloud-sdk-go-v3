package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResDatasourceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *CreateResDatasourceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateResDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResDatasourceRequest struct{}"
	}

	return strings.Join([]string{"CreateResDatasourceRequest", string(data)}, " ")
}
