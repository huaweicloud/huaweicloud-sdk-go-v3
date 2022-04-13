package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDatasourceReqDto struct {
	// 数据源名称

	Name string `json:"name"`
	// 数据源类型, 包括：IOTDA、API[、OBS、DIS、SMN、FUNCTION_GRAPH、MODEL_ARTS、DCS、KAFKA](tag:IoTA-Cloud-Only)、NODE。数据源不支持修改类型。

	Type string `json:"type"`

	Content *ContentDetailReq `json:"content,omitempty"`
}

func (o UpdateDatasourceReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatasourceReqDto struct{}"
	}

	return strings.Join([]string{"UpdateDatasourceReqDto", string(data)}, " ")
}
