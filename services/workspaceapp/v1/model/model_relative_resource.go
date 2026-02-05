package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelativeResource 关联资源。
type RelativeResource struct {

	// 关联的资源ID。
	RelativeResourceId *string `json:"relative_resource_id,omitempty"`

	// 有资源名称的云资源，都需要返回resourceName，如果为空值，则返回“”。
	RelativeResourceName *string `json:"relative_resource_name,omitempty"`

	// * 关联关系类型，描述relativeResourceId和resourceId间的关联关系：    * 0：挂载(弱关联)，表示relativeResourceId和resourceId两者有关联关系，     * 但是两者可以独立使用、分别进行交易，且分别使用和交易都不影响整套云服务的使用；比如云主机和数据盘。 *     1：绑定(强关联)，表示relativeResourceId和resourceId是强绑定关系，      两者必须一起使用、一起进行交易，缺少其中一个会造成整套云服务不能使用；比如云主机和系统盘。           缺省值为0(挂载)。           subResources中的RelativeResource，此字段是必填；mainResources中的RelativeResource。
	RelativeType *int32 `json:"relative_type,omitempty"`

	// 依赖云服务类型： - hws.service.type.ec2
	RelativeCloudServiceType *string `json:"relative_cloud_service_type,omitempty"`

	// 依赖资源类型： - hws.resource.type.vm
	RelativeResourceType *string `json:"relative_resource_type,omitempty"`

	// 扩展信息，Key:Value格式。
	ExtendParams *string `json:"extend_params,omitempty"`
}

func (o RelativeResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelativeResource struct{}"
	}

	return strings.Join([]string{"RelativeResource", string(data)}, " ")
}
