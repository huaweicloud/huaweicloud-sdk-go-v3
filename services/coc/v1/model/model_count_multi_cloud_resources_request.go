package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountMultiCloudResourcesRequest Request Object
type CountMultiCloudResourcesRequest struct {

	// **参数解释：** 云厂商类型。 **约束限制：** 不涉及。 **取值范围：** - AWS：亚马逊。 - AZURE：微软。 - ALI：阿里云。 - HCS：Huawei Cloud Stack。 **默认取值：** 不涉及。
	Vendor string `json:"vendor"`

	// **参数解释：** 资源类型。 **约束限制：** 不涉及。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 用户选择的资源id组成的列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`

	// **参数解释：** 资源名称。 **约束限制：** 不涉及。 **取值范围：** 列表，可参考：裸金属服务器BMS。 **默认取值：** 不涉及。
	NameList *[]string `json:"name_list,omitempty"`

	// **参数解释：** 关联的区域region的id组成的列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RegionIdList *[]string `json:"region_id_list,omitempty"`
}

func (o CountMultiCloudResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountMultiCloudResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountMultiCloudResourcesRequest", string(data)}, " ")
}
