package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebAppAndServicesResponse Response Object
type ListWebAppAndServicesResponse struct {

	// **参数解释** 总数 **取值范围** 最小值0，最大值300000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 具有该web应用、web服务、数据库资产的主机及该资产信息列表 **取值范围** 最小值0，最大值10000
	DataList       *[]WebAppAndServiceResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListWebAppAndServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebAppAndServicesResponse struct{}"
	}

	return strings.Join([]string{"ListWebAppAndServicesResponse", string(data)}, " ")
}
