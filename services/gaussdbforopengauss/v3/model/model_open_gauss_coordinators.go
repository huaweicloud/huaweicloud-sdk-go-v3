package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CN横向扩容时必填
type OpenGaussCoordinators struct {

	// 新增CN横向扩容每个节点的可用区。如果需要扩容多个CN，请分别填写待扩容CN所在的可用区。  不同区域的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)。  说明： 扩容后，实例中CN节点的数量必须小于或等于两倍的分片数量。
	AzCode string `json:"az_code" xml:"az_code"`
}

func (o OpenGaussCoordinators) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussCoordinators struct{}"
	}

	return strings.Join([]string{"OpenGaussCoordinators", string(data)}, " ")
}
