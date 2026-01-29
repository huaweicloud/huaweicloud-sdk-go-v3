package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCollectConfigResponseBodyCsvcList struct {

	// 云产品ID
	Csvc *string `json:"csvc,omitempty"`

	// 所有的日志类型
	SourceList *[]ListCollectConfigResponseBodySourceList `json:"source_list,omitempty"`
}

func (o ListCollectConfigResponseBodyCsvcList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyCsvcList struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyCsvcList", string(data)}, " ")
}
