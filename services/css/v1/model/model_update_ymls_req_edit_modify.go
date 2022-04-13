package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置文件信息。
type UpdateYmlsReqEditModify struct {
	// 参数配置列表。

	ElasticsearchYml *interface{} `json:"elasticsearch.yml"`
}

func (o UpdateYmlsReqEditModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsReqEditModify struct{}"
	}

	return strings.Join([]string{"UpdateYmlsReqEditModify", string(data)}, " ")
}
