package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiViewModelMasterIdModifierDto struct {

	// 主对象ID。
	MasterId string `json:"masterId"`

	// 处理人。
	Modifier *string `json:"modifier,omitempty"`

	Item *ObjectReferenceParamDto `json:"item,omitempty"`
}

func (o MultiViewModelMasterIdModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelMasterIdModifierDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelMasterIdModifierDto", string(data)}, " ")
}
