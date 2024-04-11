package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesignModelsResultData 返回的数据信息。
type ExportDesignModelsResultData struct {

	// 导入接口返回的唯一标识。
	Uuid *string `json:"uuid,omitempty"`
}

func (o ExportDesignModelsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesignModelsResultData struct{}"
	}

	return strings.Join([]string{"ExportDesignModelsResultData", string(data)}, " ")
}
