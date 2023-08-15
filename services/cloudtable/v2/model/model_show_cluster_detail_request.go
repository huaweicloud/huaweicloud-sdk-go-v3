package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterDetailRequest Request Object
type ShowClusterDetailRequest struct {

	// 语言类型
	XLanguage string `json:"X-Language"`

	// 集群ID。  获取方法：在ClooudTable控制台，单击要查询的集群名称进入集群详情页，获取“集群ID\"。
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailRequest", string(data)}, " ")
}
