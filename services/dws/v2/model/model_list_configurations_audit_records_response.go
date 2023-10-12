package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsAuditRecordsResponse Response Object
type ListConfigurationsAuditRecordsResponse struct {

	// 记录
	Records *[]ConfigurationRecordResp `json:"records,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigurationsAuditRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsAuditRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsAuditRecordsResponse", string(data)}, " ")
}
