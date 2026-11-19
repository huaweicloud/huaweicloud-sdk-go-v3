package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceClusterDetectRisk2Response Response Object
type ListCceClusterDetectRisk2Response struct {

	// **参数解释**： 集群总数 **取值范围**： 0~200
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 集群信息列表 **取值范围**： 数组长度0-200
	DataList       *[]NodeDetectRiskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListCceClusterDetectRisk2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceClusterDetectRisk2Response struct{}"
	}

	return strings.Join([]string{"ListCceClusterDetectRisk2Response", string(data)}, " ")
}
