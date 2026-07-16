package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopNotebookRequest Request Object
type StopNotebookRequest struct {

	// **参数解释**：Notebook实例ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID），可通过调用[[查询Notebook实例列表接口](https://support.huaweicloud.com/api-modelarts/ListNotebooks.html#section0)](tag:hc)[[查询Notebook实例列表接口](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListNotebooks.html#section0)](tag:hk)获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o StopNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopNotebookRequest struct{}"
	}

	return strings.Join([]string{"StopNotebookRequest", string(data)}, " ")
}
