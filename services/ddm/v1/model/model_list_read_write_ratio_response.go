package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListReadWriteRatioResponse struct {
	// DDM读写比例监控信息条数。

	TotalRecord *int32 `json:"totalRecord,omitempty"`
	// DDM实例读写次数信息列表的集合。

	ReadWriteRatioList *[]ReadWriteRatioList `json:"readWriteRatioList,omitempty"`
	HttpStatusCode     int                   `json:"-"`
}

func (o ListReadWriteRatioResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReadWriteRatioResponse struct{}"
	}

	return strings.Join([]string{"ListReadWriteRatioResponse", string(data)}, " ")
}
