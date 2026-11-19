package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryBinlogDownloadLinkRequestV3Body struct {

	// **参数解释**：  Binlog日志ID列表。  获取方法请参见[查询Binlog日志列表](https://support.huaweicloud.com/api-taurusdb/ListBinlogs.html)。  **约束限制**：  50条以内。
	Ids *[]string `json:"ids,omitempty"`
}

func (o QueryBinlogDownloadLinkRequestV3Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryBinlogDownloadLinkRequestV3Body struct{}"
	}

	return strings.Join([]string{"QueryBinlogDownloadLinkRequestV3Body", string(data)}, " ")
}
