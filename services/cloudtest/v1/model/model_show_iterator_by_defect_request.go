package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIteratorByDefectRequest Request Object
type ShowIteratorByDefectRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 缺陷id
	DefectId string `json:"defect_id"`
}

func (o ShowIteratorByDefectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIteratorByDefectRequest struct{}"
	}

	return strings.Join([]string{"ShowIteratorByDefectRequest", string(data)}, " ")
}
