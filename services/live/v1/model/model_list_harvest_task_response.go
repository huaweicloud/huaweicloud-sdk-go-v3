package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHarvestTaskResponse Response Object
type ListHarvestTaskResponse struct {

	// 总任务数
	Total *int32 `json:"total,omitempty"`

	// 任务信息
	HarvestTasks   *[]HarvestTaskCreateSucRsp `json:"harvest_tasks,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListHarvestTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHarvestTaskResponse struct{}"
	}

	return strings.Join([]string{"ListHarvestTaskResponse", string(data)}, " ")
}
