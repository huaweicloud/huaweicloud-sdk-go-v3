package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPackageDataDetailRequest Request Object
type ShowPackageDataDetailRequest struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ShowPackageDataDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageDataDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowPackageDataDetailRequest", string(data)}, " ")
}
