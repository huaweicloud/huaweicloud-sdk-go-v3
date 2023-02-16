package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowQaPairsRequest struct {

	// 调用智能客服服务标志。
	XServiceKey string `json:"x-service-key"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// - HOT:  - RECOMMEND:
	SearchType ShowQaPairsRequestSearchType `json:"search_type"`

	Body *SearchQaPairsReq `json:"body,omitempty"`
}

func (o ShowQaPairsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQaPairsRequest struct{}"
	}

	return strings.Join([]string{"ShowQaPairsRequest", string(data)}, " ")
}

type ShowQaPairsRequestSearchType struct {
	value string
}

type ShowQaPairsRequestSearchTypeEnum struct {
	HOT       ShowQaPairsRequestSearchType
	RECOMMEND ShowQaPairsRequestSearchType
}

func GetShowQaPairsRequestSearchTypeEnum() ShowQaPairsRequestSearchTypeEnum {
	return ShowQaPairsRequestSearchTypeEnum{
		HOT: ShowQaPairsRequestSearchType{
			value: "HOT",
		},
		RECOMMEND: ShowQaPairsRequestSearchType{
			value: "RECOMMEND",
		},
	}
}

func (c ShowQaPairsRequestSearchType) Value() string {
	return c.value
}

func (c ShowQaPairsRequestSearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowQaPairsRequestSearchType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
