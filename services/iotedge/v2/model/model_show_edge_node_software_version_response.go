package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeNodeSoftwareVersionResponse Response Object
type ShowEdgeNodeSoftwareVersionResponse struct {

	// 边缘应用id，只允许数字、英文小写、中划线，切必须以字母或数字结尾
	SoftwareVersion *string `json:"software_version,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowEdgeNodeSoftwareVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeSoftwareVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeSoftwareVersionResponse", string(data)}, " ")
}
