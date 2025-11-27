package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FunctionGraphUrnRequiredPrimitiveTypeHolder struct {

	// FunctionGraph方法的统一资源标识，用于标识唯一的FunctionGraph方法。当前只支持和RFS同region的function_graph_urn，如果给予了关于其他region的，会报错400。  关于该参数的详细解释，请参考官方文档：https://support.huaweicloud.com/api-functiongraph/functiongraph_06_0102.html
	FunctionGraphUrn string `json:"function_graph_urn"`
}

func (o FunctionGraphUrnRequiredPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionGraphUrnRequiredPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"FunctionGraphUrnRequiredPrimitiveTypeHolder", string(data)}, " ")
}
