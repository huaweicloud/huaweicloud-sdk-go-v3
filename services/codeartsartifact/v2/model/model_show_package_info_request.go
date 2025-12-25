package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPackageInfoRequest Request Object
type ShowPackageInfoRequest struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ShowPackageInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowPackageInfoRequest", string(data)}, " ")
}
