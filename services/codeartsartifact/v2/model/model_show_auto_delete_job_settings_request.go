package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoDeleteJobSettingsRequest Request Object
type ShowAutoDeleteJobSettingsRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`
}

func (o ShowAutoDeleteJobSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoDeleteJobSettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoDeleteJobSettingsRequest", string(data)}, " ")
}
