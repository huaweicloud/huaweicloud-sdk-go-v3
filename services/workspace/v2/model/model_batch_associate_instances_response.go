package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateInstancesResponse Response Object
type BatchAssociateInstancesResponse struct {

	// 分配的维度，当前支持“用户为维度”、“桌面为维度”两种。
	AssignDimension *string `json:"assign_dimension,omitempty"`

	// 桌面维度结果，如果assign_dimension为DEKSTOP，那么取这个结果。
	Desktop *[]DesktopDimensionAttachInfo `json:"desktop,omitempty"`

	// 用户维度结果，如果assign_dimension为USER，那么取这个结果。
	User *[]UserDimensionAttachInfo `json:"user,omitempty"`

	// 桌面分配关系。
	DesktopsAttachInfos *[]AttachInstancesDesktopInfo `json:"desktops_attach_infos,omitempty"`
	HttpStatusCode      int                           `json:"-"`
}

func (o BatchAssociateInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchAssociateInstancesResponse", string(data)}, " ")
}
