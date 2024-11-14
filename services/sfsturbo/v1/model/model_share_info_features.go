package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShareInfoFeatures 文件系统的特性支持情况。
type ShareInfoFeatures struct {
	Backup *ShareInfoFeature `json:"backup,omitempty"`
}

func (o ShareInfoFeatures) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareInfoFeatures struct{}"
	}

	return strings.Join([]string{"ShareInfoFeatures", string(data)}, " ")
}
