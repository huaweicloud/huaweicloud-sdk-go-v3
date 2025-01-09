package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSysprepInfoResponse Response Object
type ShowSysprepInfoResponse struct {

	// sysprep版本。
	SysprepVersion *string `json:"sysprep_version,omitempty"`

	// 是否支持创建镜像。
	SupportCreateImage *bool `json:"support_create_image,omitempty"`
	HttpStatusCode     int   `json:"-"`
}

func (o ShowSysprepInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSysprepInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSysprepInfoResponse", string(data)}, " ")
}
