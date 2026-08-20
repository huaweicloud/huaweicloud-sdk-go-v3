package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdLabelRequest Request Object
type CreateIpdLabelRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	Body *LabelParam `json:"body,omitempty"`
}

func (o CreateIpdLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdLabelRequest struct{}"
	}

	return strings.Join([]string{"CreateIpdLabelRequest", string(data)}, " ")
}
