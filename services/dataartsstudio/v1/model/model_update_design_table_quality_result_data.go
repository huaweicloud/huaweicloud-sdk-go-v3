package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignTableQualityResultData 更新表的异常数据输出配置返回数据。
type UpdateDesignTableQualityResultData struct {
	Value *TableModelVo `json:"value,omitempty"`
}

func (o UpdateDesignTableQualityResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignTableQualityResultData struct{}"
	}

	return strings.Join([]string{"UpdateDesignTableQualityResultData", string(data)}, " ")
}
