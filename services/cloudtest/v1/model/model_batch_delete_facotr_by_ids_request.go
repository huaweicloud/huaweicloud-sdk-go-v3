package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFacotrByIdsRequest Request Object
type BatchDeleteFacotrByIdsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestListString `json:"body,omitempty"`
}

func (o BatchDeleteFacotrByIdsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFacotrByIdsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteFacotrByIdsRequest", string(data)}, " ")
}
