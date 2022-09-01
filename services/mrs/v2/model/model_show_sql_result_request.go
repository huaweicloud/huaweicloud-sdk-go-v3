package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSqlResultRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// SQL的执行ID，即提交SQL语句返回结果中的sql_id。
	SqlId string `json:"sql_id" xml:"sql_id"`
}

func (o ShowSqlResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlResultRequest", string(data)}, " ")
}
