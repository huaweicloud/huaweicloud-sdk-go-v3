package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFavoriteRequest Request Object
type DeleteFavoriteRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 收藏ID。
	FavoriteId string `json:"favorite_id"`
}

func (o DeleteFavoriteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFavoriteRequest struct{}"
	}

	return strings.Join([]string{"DeleteFavoriteRequest", string(data)}, " ")
}
