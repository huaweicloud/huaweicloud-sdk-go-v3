package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIncidentRequestBody 更新事件请求body体
type ChangeIncidentRequestBody struct {

	// 更新事件的ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`

	DataObject *Incident `json:"data_object,omitempty"`
}

func (o ChangeIncidentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIncidentRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeIncidentRequestBody", string(data)}, " ")
}
