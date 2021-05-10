package model

import (
	"encoding/json"

	"strings"
)

//
type V3NodeEipSpec struct {
	// 弹性IP类型，取值请参见“[[创建云服务器](https://support.huaweicloud.com/api-ecs/zh-cn_topic_0167957246.html)](tag:hws)[[创建云服务器](https://support.huaweicloud.com/intl/zh-cn/api-ecs/zh-cn_topic_0167957246.html)](tag:hws_hk) > eip字段数据结构说明”表中“iptype”参数的描述。

	Iptype *string `json:"iptype,omitempty"`

	Bandwidth *V3NodeBandwidth `json:"bandwidth,omitempty"`
}

func (o V3NodeEipSpec) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "V3NodeEipSpec struct{}"
	}

	return strings.Join([]string{"V3NodeEipSpec", string(data)}, " ")
}
