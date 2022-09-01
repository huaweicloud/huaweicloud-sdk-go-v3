package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 负载均衡器规格信息。
type Flavor struct {

	// 规格ID。
	Id string `json:"id" xml:"id"`

	Info *FlavorInfo `json:"info" xml:"info"`

	// 规格名称。
	Name string `json:"name" xml:"name"`

	// 是否公共规格。取值：  true表示公共规格，所有租户可见。 false表示私有规格，为当前租户所有。
	Shared bool `json:"shared" xml:"shared"`

	// 项目ID。
	ProjectId string `json:"project_id" xml:"project_id"`

	// 规格类别。取值： - L4和L7 表示四层和七层flavor。 - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。 - L4_elastic_max和L7_elastic_max 表示弹性扩缩容实例的上限规格。
	Type string `json:"type" xml:"type"`

	// 是否售罄。取值： - true：已售罄，将无法购买该规格的LB。 - false：未售罄，可购买该规格的LB。
	FlavorSoldOut bool `json:"flavor_sold_out" xml:"flavor_sold_out"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
