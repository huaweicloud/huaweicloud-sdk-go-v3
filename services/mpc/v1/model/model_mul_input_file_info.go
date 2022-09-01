package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MulInputFileInfo struct {

	// 语言标签。
	Language *string `json:"language,omitempty" xml:"language"`

	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`
}

func (o MulInputFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MulInputFileInfo struct{}"
	}

	return strings.Join([]string{"MulInputFileInfo", string(data)}, " ")
}
