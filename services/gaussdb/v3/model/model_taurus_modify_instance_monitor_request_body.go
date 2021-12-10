package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 秒级监控修改请求
type TaurusModifyInstanceMonitorRequestBody struct {
	// 采集周期。  取值： 0表示实例秒级监控关闭； 1表示实例秒级监控开启，采集周期为1s； 5表示实例秒级监控开启，采集周期为5s。

	Period string `json:"period"`
}

func (o TaurusModifyInstanceMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusModifyInstanceMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"TaurusModifyInstanceMonitorRequestBody", string(data)}, " ")
}
