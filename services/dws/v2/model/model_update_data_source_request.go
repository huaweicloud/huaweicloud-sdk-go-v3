package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataSourceRequest Request Object
type UpdateDataSourceRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 数据源id
	ExtDataSourceId string `json:"ext_data_source_id"`

	Body *ReconfigureExtDataSourceActionReq `json:"body,omitempty"`
}

func (o UpdateDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataSourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataSourceRequest", string(data)}, " ")
}
