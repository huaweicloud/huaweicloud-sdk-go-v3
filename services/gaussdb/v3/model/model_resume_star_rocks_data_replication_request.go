package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeStarRocksDataReplicationRequest Request Object
type ResumeStarRocksDataReplicationRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	Body *ResumeStarRocksDataReplication `json:"body,omitempty"`
}

func (o ResumeStarRocksDataReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeStarRocksDataReplicationRequest struct{}"
	}

	return strings.Join([]string{"ResumeStarRocksDataReplicationRequest", string(data)}, " ")
}
