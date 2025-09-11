package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashBoardInfo struct {

	// **参数描述**： 监控看板id **取值范围** 以db开头，包含22个字母和数字，长度为24个字符
	DashboardId *string `json:"dashboard_id,omitempty"`

	// **参数解释** 自定义监控看板名称 **取值范围** 长度为[1,128]个字符，只允许中文、英文、数字0-9、_和-
	DashboardName *string `json:"dashboard_name,omitempty"`

	// **参数解释** 企业项目ID **取值范围** 只能包含小写字母、数字、“-”、“_”，可以自定义企业项目ID，长度为36个字符。也可以为0（代表默认企业项目ID）。
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// **参数解释** 每行展示视图数量 **取值范围** - 0:表示自定义坐标 - 1:代表每行1个视图 - 2:代表每行2个视图 - 3:代表每行3个视图
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`

	// **参数解释** 监控看板是否标记收藏 **取值范围** - true: 收藏, - false: 未收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// **参数解释** 监控看板的创建用户名 **取值范围** 长度为[1,128]个字符，只允许中文、英文、数字0-9、_和-
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释** 监控看板创建时间 **取值范围** 最小值为1111111111111，最大值为9999999999999
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释** 看板下的视图总数 **取值范围** 最小值为0，最大值为50
	WidgetsNum *int32 `json:"widgets_num,omitempty"`

	// **参数解释** 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)” **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 子产品标识 **取值范围** 长度为[1,128]个字符
	SubProduct *string `json:"sub_product,omitempty"`

	// **参数解释** 监控大盘模板id **取值范围** 以mb开头，包含22个字母和数字，长度为24个字符
	DashboardTemplateId *string `json:"dashboard_template_id,omitempty"`
}

func (o DashBoardInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardInfo struct{}"
	}

	return strings.Join([]string{"DashBoardInfo", string(data)}, " ")
}
