package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncidentRequestBody 删除事件请求body体
type DeleteIncidentRequestBody struct {

	// 删除事件的ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o DeleteIncidentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncidentRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteIncidentRequestBody", string(data)}, " ")
}
