package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobListByProjectIdRequest Request Object
type ShowJobListByProjectIdRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 分页页码， 表示从此页开始查询， page_index大于等于0
	PageIndex int32 `json:"page_index"`

	// 每页显示的条目数量，page_size小于等于100
	PageSize int32 `json:"page_size"`
}

func (o ShowJobListByProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobListByProjectIdRequest struct{}"
	}

	return strings.Join([]string{"ShowJobListByProjectIdRequest", string(data)}, " ")
}
