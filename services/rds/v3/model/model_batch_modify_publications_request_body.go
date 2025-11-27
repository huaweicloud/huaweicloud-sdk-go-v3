package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyPublicationsRequestBody 修改发布详情。
type BatchModifyPublicationsRequestBody struct {

	// 修改的发布id列表。
	PublicationIds []string `json:"publication_ids"`

	JobSchedule *OperateUsedJobSchedule `json:"job_schedule"`
}

func (o BatchModifyPublicationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyPublicationsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchModifyPublicationsRequestBody", string(data)}, " ")
}
