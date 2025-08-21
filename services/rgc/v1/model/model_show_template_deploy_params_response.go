package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateDeployParamsResponse Response Object
type ShowTemplateDeployParamsResponse struct {

	// 模板的部署参数。
	Variables      *[]TemplateParamVariable `json:"variables,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowTemplateDeployParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateDeployParamsResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateDeployParamsResponse", string(data)}, " ")
}
