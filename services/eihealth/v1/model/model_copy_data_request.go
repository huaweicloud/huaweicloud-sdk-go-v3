package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyDataRequest Request Object
type CopyDataRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CloneDataReq `json:"body,omitempty"`
}

func (o CopyDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyDataRequest struct{}"
	}

	return strings.Join([]string{"CopyDataRequest", string(data)}, " ")
}
