package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IndependentBodyReq struct {

	// 规格id，该参数通过[获取实例规格列表](ListFlavors.xml)接口获取，根据集群版本选择所需要的规格id
	FlavorRef string `json:"flavor_ref"`

	// 要独立节点个数。 - 如果路径参数type取值为“ess-master”即新增独立master节点，节点个数必须为大于等于三且小于等于10的奇数。 - 如果路径参数type取值为“ess-client”即新增独立client节点，节点个数要求大于等于1小于等于32。
	NodeSize int32 `json:"node_size"`

	// 节点存储类型：取值为ULTRAHIGH，COMMON，HIGH。
	VolumeType *string `json:"volume_type,omitempty"`

	// **参数解释**： 节点存储大小。 **约束限制**： - flavor_ref参数是本地盘规格时，不能设置此参数。 - 必须为大于0且为4和10的公倍数，单位：GB。 **取值范围**： 磁盘规格大小可以通过过[获取实例规格列表](ListFlavors.xml)中diskrange属性获得。 **默认取值**： flavor_ref参数不是本地盘规格时： - 新增独立冷数据节点：默认100G与节点规格支持的最小磁盘容量取较大者。 - 新增独立master或client节点：默认大小为40G，且不可更改。  >新增独立冷数据节点，推荐大于100G。
	VolumeSize *int32 `json:"volume_size,omitempty"`
}

func (o IndependentBodyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndependentBodyReq struct{}"
	}

	return strings.Join([]string{"IndependentBodyReq", string(data)}, " ")
}
