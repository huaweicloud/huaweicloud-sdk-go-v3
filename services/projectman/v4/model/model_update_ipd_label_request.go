package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpdLabelRequest Request Object
type UpdateIpdLabelRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 标签ID，标签唯一标识。 可以通过查询标签列表接口获取，响应消息体中的id字段的值就是标签ID。
	LabelId string `json:"label_id"`

	Body *LabelParam `json:"body,omitempty"`
}

func (o UpdateIpdLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpdLabelRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpdLabelRequest", string(data)}, " ")
}
