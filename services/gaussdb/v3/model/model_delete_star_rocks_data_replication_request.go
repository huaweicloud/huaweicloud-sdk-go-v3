package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRocksDataReplicationRequest Request Object
type DeleteStarRocksDataReplicationRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	Body *DeleteStarRocksDataReplication `json:"body,omitempty"`
}

func (o DeleteStarRocksDataReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRocksDataReplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteStarRocksDataReplicationRequest", string(data)}, " ")
}
