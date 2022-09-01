package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDataSourceResponse struct {

	// 数据源id
	Id *string `json:"id,omitempty" xml:"id"`

	// 数据源名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 数据源类型, 包括：IOTDA、API[、OBS、DIS、SMN、FUNCTION_GRAPH、MODEL_ARTS、DCS、KAFKA](tag:IoTA-Cloud-Only)、NODE。数据源不支持修改类型。
	Type *string `json:"type,omitempty" xml:"type"`

	Content *ContentDetailRsp `json:"content,omitempty" xml:"content"`

	// 创建时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 修改时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	ModifiedTime   *string `json:"modified_time,omitempty" xml:"modified_time"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataSourceResponse struct{}"
	}

	return strings.Join([]string{"ShowDataSourceResponse", string(data)}, " ")
}
