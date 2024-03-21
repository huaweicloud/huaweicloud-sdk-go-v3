package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiViewModelViewUpdateAttrDto struct {
	Item *ObjectReferenceParamDto `json:"item,omitempty"`
}

func (o MultiViewModelViewUpdateAttrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelViewUpdateAttrDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelViewUpdateAttrDto", string(data)}, " ")
}
