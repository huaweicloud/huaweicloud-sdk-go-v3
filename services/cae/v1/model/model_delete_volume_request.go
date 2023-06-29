package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumeRequest Request Object
type DeleteVolumeRequest struct {

	// 卷id。
	Id string `json:"id"`

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`
}

func (o DeleteVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeRequest struct{}"
	}

	return strings.Join([]string{"DeleteVolumeRequest", string(data)}, " ")
}
