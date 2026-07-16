package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDynamicStorageRequest Request Object
type AttachDynamicStorageRequest struct {

	// **参数解释**：Notebook实例ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID），可通过调用[[查询Notebook实例列表接口](https://support.huaweicloud.com/api-modelarts/ListNotebooks.html#section0)](tag:hc)[[查询Notebook实例列表接口](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListNotebooks.html#section0)](tag:hk)获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	InstanceId string `json:"instance_id"`

	Body *DataVolumesReq `json:"body,omitempty"`
}

func (o AttachDynamicStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDynamicStorageRequest struct{}"
	}

	return strings.Join([]string{"AttachDynamicStorageRequest", string(data)}, " ")
}
