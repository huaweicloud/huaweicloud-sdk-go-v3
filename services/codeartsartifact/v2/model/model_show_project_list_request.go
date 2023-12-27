package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectListRequest Request Object
type ShowProjectListRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`
}

func (o ShowProjectListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectListRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectListRequest", string(data)}, " ")
}
