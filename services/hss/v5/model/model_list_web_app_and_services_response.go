package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebAppAndServicesResponse Response Object
type ListWebAppAndServicesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 具有该WebAppAndService资产的主机及该资产信息列表
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
