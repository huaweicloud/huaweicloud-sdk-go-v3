package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageRiskConfigsResponse Response Object
type ListImageRiskConfigsResponse struct {

	// **参数解释** 符合筛选条件的镜像安全配置检测结果总记录数 **取值范围** 取值0-2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 包含查询到的镜像安全配置检测结果详情，每个元素对应一个镜像的基线检测汇总信息 **取值范围** 数组长度0-limit（每页显示个数）
	DataList       *[]ImageRiskConfigsInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListImageRiskConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageRiskConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListImageRiskConfigsResponse", string(data)}, " ")
}
