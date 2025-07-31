package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostTagInfoList struct {

	// 资源标签列表
	DataList *[]HostTagValuesInfo `json:"data_list,omitempty"`

	// 资源标签总数
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o HostTagInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostTagInfoList struct{}"
	}

	return strings.Join([]string{"HostTagInfoList", string(data)}, " ")
}
