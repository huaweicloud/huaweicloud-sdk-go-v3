package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventRequest Request Object
type ListEventRequest struct {

	// 语言，默认值为en-us。zh-cn（中文）/en-us（英文）
	XLanguage *string `json:"X-Language,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 查询日志的时间范围（不能和from、to同时使用，同时使用以recent为准），且recent参数与from、to必须使用其中一个。当同时使用recent参数与from、to时，以recent参数为准 **约束限制：** 不涉及 **取值范围：**  - yesterday：昨天  - today：今天  - 3days：命令注入攻击   - 1week：XSS攻击   - 1month：恶意爬虫   **默认取值：** 不涉及
	Recent *ListEventRequestRecent `json:"recent,omitempty"`

	// 起始时间(13位时间戳)，需要和to同时使用，不能和recent参数同时使用
	From *int64 `json:"from,omitempty"`

	// 结束时间(13位时间戳)，需要和from同时使用，不能和recent参数同时使用
	To *int64 `json:"to,omitempty"`

	// **参数解释：** 攻击类型 **约束限制：** 不涉及 **取值范围：**  - sqli：sql注入攻击   - lfi：本地文件包含  - cmdi：命令注入攻击   - xss：XSS攻击   - robot：恶意爬虫   - rfi：远程文件包含   - custom_custom：精准防护   - cc: cc攻击   - webshell：网站木马   - custom_whiteblackip：黑白名单拦截   - custom_geoip：地理访问控制拦截   - antitamper：防篡改   - anticrawler：反爬虫    - leakage：网站信息防泄漏   - illegal：非法请求  - antiscan_high_freq_scan：高频扫描封禁  - antiscan_dir_traversal：目录遍历防护  - vuln：除上述攻击类型外的其他漏洞攻击  **默认取值：** 不涉及
	Attacks *[]string `json:"attacks,omitempty"`

	// 域名id，从获取防护网站列表（ListHost）接口获取域名id
	Hosts *[]string `json:"hosts,omitempty"`

	// 源ip，Web访问者的IP地址（攻击者IP地址）
	Sips *[]string `json:"sips,omitempty"`

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *int32 `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventRequest struct{}"
	}

	return strings.Join([]string{"ListEventRequest", string(data)}, " ")
}

type ListEventRequestRecent struct {
	value string
}

type ListEventRequestRecentEnum struct {
	YESTERDAY ListEventRequestRecent
	TODAY     ListEventRequestRecent
	E_3DAYS   ListEventRequestRecent
	E_1WEEK   ListEventRequestRecent
	E_1MONTH  ListEventRequestRecent
}

func GetListEventRequestRecentEnum() ListEventRequestRecentEnum {
	return ListEventRequestRecentEnum{
		YESTERDAY: ListEventRequestRecent{
			value: "yesterday",
		},
		TODAY: ListEventRequestRecent{
			value: "today",
		},
		E_3DAYS: ListEventRequestRecent{
			value: "3days",
		},
		E_1WEEK: ListEventRequestRecent{
			value: "1week",
		},
		E_1MONTH: ListEventRequestRecent{
			value: "1month",
		},
	}
}

func (c ListEventRequestRecent) Value() string {
	return c.value
}

func (c ListEventRequestRecent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestRecent) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
