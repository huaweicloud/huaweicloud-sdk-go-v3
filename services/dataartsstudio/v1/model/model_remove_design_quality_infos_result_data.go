package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDesignQualityInfosResultData 清空质量规则返回数据。
type RemoveDesignQualityInfosResultData struct {

	// 是否清除成功。
	Value *bool `json:"value,omitempty"`
}

func (o RemoveDesignQualityInfosResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDesignQualityInfosResultData struct{}"
	}

	return strings.Join([]string{"RemoveDesignQualityInfosResultData", string(data)}, " ")
}
