package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserAccessInfoRequest Request Object
type ShowUserAccessInfoRequest struct {

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o ShowUserAccessInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserAccessInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowUserAccessInfoRequest", string(data)}, " ")
}
