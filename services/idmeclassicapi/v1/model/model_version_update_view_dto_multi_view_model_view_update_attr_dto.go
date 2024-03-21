package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto struct {

	// 版本对象ID
	VersionId string `json:"versionId"`

	NewViewAttrs *MultiViewModelViewUpdateAttrDto `json:"newViewAttrs"`

	// 指定不复制的视图属性，其值将被设置为null。
	NeedSetNull *[]string `json:"needSetNull,omitempty"`
}

func (o VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto struct{}"
	}

	return strings.Join([]string{"VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto", string(data)}, " ")
}
