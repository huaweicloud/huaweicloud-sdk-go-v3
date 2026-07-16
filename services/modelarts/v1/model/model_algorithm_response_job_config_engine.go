package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmResponseJobConfigEngine 算法的引擎。
type AlgorithmResponseJobConfigEngine struct {

	// 算法选择的引擎规格ID。
	EngineId *string `json:"engine_id,omitempty"`

	// 算法选择的引擎版本名称。若填入engine_id则无需填写。
	EngineName *string `json:"engine_name,omitempty"`

	// 算法选择的引擎版本名称。若填入engine_id则无需填写。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 算法选择的自定义镜像地址。
	ImageUrl *string `json:"image_url,omitempty"`
}

func (o AlgorithmResponseJobConfigEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseJobConfigEngine struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseJobConfigEngine", string(data)}, " ")
}
