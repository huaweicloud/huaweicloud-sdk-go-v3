package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPackageSpecCountDownResponse Response Object
type ShowPackageSpecCountDownResponse struct {

	// **参数解释**： 接口响应状态。 **取值范围**： ● success：表示接口调用成功。 ● fail：表示接口调用失败。
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 单元测试报告列表
	Result         *[]CountdownList `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowPackageSpecCountDownResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageSpecCountDownResponse struct{}"
	}

	return strings.Join([]string{"ShowPackageSpecCountDownResponse", string(data)}, " ")
}
