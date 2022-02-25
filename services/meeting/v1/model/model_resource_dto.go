package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDto struct {
	// 资源标识，不携带则后台自动生成

	Id *string `json:"id,omitempty"`
	// 资源类型，前台通过查询接口返回该sp支持售卖的资源在界面上做相应屏蔽，当前为枚举类型. - VMR        - 云会议室 - CONF_CALL  - 会议并发数 - HARD_1080P - 1080P硬终端 - HARD_720P  - 720P硬终端 - SOFT       - 软终端用户数 - ROOM       - 大屏软终端 - LIVE       - 直播推流 - RECORD     - 录播空间 - HARD_THIRD_PARTY - 第三方硬终端账号 - HUAWEI_VISION -智慧屏 - IDEA_HUB   - ideahub

	Type string `json:"type"`
	// 类型标识，比如资源类型为vmr，vmr又分为5方，10方等，该参数为vmrPkgId，用来区分子类别，详见如下： - vmr10:ff808081699b56d40169c410d5080179 - vmr50:ff808081699b56cb0169c411a0980152 - vmr100:ff808081699b56cb0169c41167850151 - vmr200:ff808081699b56d40169c410913d0178 - vmr25:ff808081699b56d40169c4111fe5017a - vmr300:ff8080816b9ec3ab016bdff237962e83 - vmr400:ff8080816b9ec475016bdff37efc279f - vmr500:ff8080816b9ec3ab016bdff338542e84

	TypeId *string `json:"typeId,omitempty"`
	// 资源数量

	Count int32 `json:"count"`
	// 到期时间,utc时间戳

	ExpireDate int64 `json:"expireDate"`
}

func (o ResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
