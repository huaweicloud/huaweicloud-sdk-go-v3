package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProgressRequest Request Object
type ShowProgressRequest struct {

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 异步操作uri
	OperationUri string `json:"operation_uri"`
}

func (o ShowProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowProgressRequest", string(data)}, " ")
}
