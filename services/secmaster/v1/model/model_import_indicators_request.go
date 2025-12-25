package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIndicatorsRequest Request Object
type ImportIndicatorsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ImportIndicatorsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportIndicatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIndicatorsRequest struct{}"
	}

	return strings.Join([]string{"ImportIndicatorsRequest", string(data)}, " ")
}
