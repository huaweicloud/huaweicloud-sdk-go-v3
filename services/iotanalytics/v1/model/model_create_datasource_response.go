package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatasourceResponse Response Object
type CreateDatasourceResponse struct {

	// 数据源id
	Id *string `json:"id,omitempty"`

	// 数据源名称
	Name *string `json:"name,omitempty"`

	// 数据源类型, 包括：IOTDA、API[、OBS、DIS、SMN、FUNCTION_GRAPH、MODEL_ARTS、DCS、KAFKA](tag:IoTA-Cloud-Only)、NODE。数据源不支持修改类型。
	Type *string `json:"type,omitempty"`

	Content *ContentDetailRsp `json:"content,omitempty"`

	// 创建时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	CreatedTime *string `json:"created_time,omitempty"`

	// 修改时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	ModifiedTime   *string `json:"modified_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatasourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceResponse struct{}"
	}

	return strings.Join([]string{"CreateDatasourceResponse", string(data)}, " ")
}
