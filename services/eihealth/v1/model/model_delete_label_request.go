package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLabelRequest Request Object
type DeleteLabelRequest struct {

	// 标签id
	LabelId string `json:"label_id"`
}

func (o DeleteLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLabelRequest struct{}"
	}

	return strings.Join([]string{"DeleteLabelRequest", string(data)}, " ")
}
