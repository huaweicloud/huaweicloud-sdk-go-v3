package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业的资源信息
type AddCorpResDto struct {

	// 企业待添加的资源列表
	Resource *[]ResourceDto `json:"resource,omitempty" xml:"resource"`
}

func (o AddCorpResDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpResDto struct{}"
	}

	return strings.Join([]string{"AddCorpResDto", string(data)}, " ")
}
