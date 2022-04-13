package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量删除边缘实例请求体
type DeleteInstancesRequestBody struct {
	// 边缘实例ID列表。 > IEC默认同步删除边缘实例的弹性公网IP和磁盘。

	Servers []BaseId `json:"servers"`
}

func (o DeleteInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteInstancesRequestBody", string(data)}, " ")
}
