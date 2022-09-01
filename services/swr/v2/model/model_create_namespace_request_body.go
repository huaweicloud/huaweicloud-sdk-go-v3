package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNamespaceRequestBody struct {

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace" xml:"namespace"`
}

func (o CreateNamespaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNamespaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNamespaceRequestBody", string(data)}, " ")
}
