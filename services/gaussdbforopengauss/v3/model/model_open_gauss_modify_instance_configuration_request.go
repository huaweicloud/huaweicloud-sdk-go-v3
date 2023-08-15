package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenGaussModifyInstanceConfigurationRequest struct {

	// 参数值对象Map<String,String>，用户基于默认参数模板自定义的参数值。
	Values map[string]string `json:"values"`
}

func (o OpenGaussModifyInstanceConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussModifyInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"OpenGaussModifyInstanceConfigurationRequest", string(data)}, " ")
}
