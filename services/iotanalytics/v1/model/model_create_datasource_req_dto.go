package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDatasourceReqDto struct {

	// 数据源名称
	Name string `json:"name"`

	// 数据源类型, 包括：IOTDA、API[、OBS、DIS、SMN、FUNCTION_GRAPH、MODEL_ARTS、DCS、KAFKA](tag:IoTA-Cloud-Only)、NODE
	Type string `json:"type"`

	Content *ContentDetailReq `json:"content"`
}

func (o CreateDatasourceReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceReqDto struct{}"
	}

	return strings.Join([]string{"CreateDatasourceReqDto", string(data)}, " ")
}
