package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReleaseRequest Request Object
type DeleteReleaseRequest struct {

	// 模板实例名称
	Name string `json:"name"`

	// 模板实例所在的命名空间
	Namespace string `json:"namespace"`

	// **参数解释：** 是否展示模板实例的资源信息。 **约束限制：** 不涉及 **取值范围：** 指定为“true”时展示模板实例的资源信息，不指定该参数时默认不展示。 **默认取值：** 无
	ShowResources *string `json:"show_resources,omitempty"`

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o DeleteReleaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReleaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteReleaseRequest", string(data)}, " ")
}
