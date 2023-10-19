package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlIncrementalBackupListRequest Request Object
type ShowGaussMySqlIncrementalBackupListRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数。默认为10，可取范围：10、20、50。
	Limit *string `json:"limit,omitempty"`

	// 查询开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  “begin_time”有值时，“end_time”必选。
	BeginTime *string `json:"begin_time,omitempty"`

	// 查询结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  “end_time”有值时，“begin_time”必选。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowGaussMySqlIncrementalBackupListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlIncrementalBackupListRequest struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlIncrementalBackupListRequest", string(data)}, " ")
}
