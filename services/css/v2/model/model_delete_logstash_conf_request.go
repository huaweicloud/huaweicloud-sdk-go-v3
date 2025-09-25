package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogstashConfRequest Request Object
type DeleteLogstashConfRequest struct {

	// 指定删除配置文件的集群ID。获取方法请参见[获取集群ID](css_03_0101.xml)。
	ClusterId string `json:"cluster_id"`

	Body *DeleteConfReqNew `json:"body,omitempty"`
}

func (o DeleteLogstashConfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogstashConfRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogstashConfRequest", string(data)}, " ")
}
