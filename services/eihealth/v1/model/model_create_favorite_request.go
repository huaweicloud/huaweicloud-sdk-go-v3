package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFavoriteRequest Request Object
type CreateFavoriteRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateFavoriteReq `json:"body,omitempty"`
}

func (o CreateFavoriteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFavoriteRequest struct{}"
	}

	return strings.Join([]string{"CreateFavoriteRequest", string(data)}, " ")
}
