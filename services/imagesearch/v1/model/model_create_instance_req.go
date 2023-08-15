package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceReq 创建实例的body体
type CreateInstanceReq struct {

	// 实例名称。  - 仅支持小写字母（a-z）、数字，横杠和下划线。  - 以字母开头，长度在16位以内。
	Name string `json:"name"`

	// 模型名称，支持如下模型名称：  - common-search，通用图片搜索，适用于图片库中搜索相似内容或类别的图片。  - image-recommend，版权图片推荐，适用于版权摄影图片库中查找并推荐相同或相似版权图片。  - image-copyright，图片版权，适用于从海量图片库中快速识别侵权盗用图片。
	Model string `json:"model"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 规格，即实例的图片数量规格。默认为30000000（单位：张），当前仅支持默认规格。
	Level *int32 `json:"level,omitempty"`

	// 图片自定义标签，每个实例最多支持10个标签，自定义标签不支持英文字母大写。
	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceReq", string(data)}, " ")
}
