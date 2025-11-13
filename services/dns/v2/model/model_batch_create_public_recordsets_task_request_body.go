package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreatePublicRecordsetsTaskRequestBody struct {

	// 批量创建记录集列表。
	Recordsets []BatchCreatePublicRecordsetsTaskItem `json:"recordsets"`
}

func (o BatchCreatePublicRecordsetsTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicRecordsetsTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicRecordsetsTaskRequestBody", string(data)}, " ")
}
