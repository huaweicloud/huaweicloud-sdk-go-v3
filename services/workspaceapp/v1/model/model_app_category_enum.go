package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppCategoryEnum 应用分类： * `GAME` - 游戏 * `BUSSINESS_INTELLIGENCE` - 商业智能 * `SECURE_STORAGE` - 安全与存储 * `MULTIMEDIA_AND_CODING` - 多媒体与编解码 * `PROJECT_MANAGEMENT` - 项目管理 * `PRODUCTIVITY_AND_COLLABORATION` - 生产力与协作 * `WEB_ADN_APPLICATION` - 网页与应用开发 * `GRAPHIC_DESIGN` - 图形设计 * `OTHER` - 其它
type AppCategoryEnum struct {
	value string
}

type AppCategoryEnumEnum struct {
	GAME                           AppCategoryEnum
	BUSSINESS_INTELLIGENCE         AppCategoryEnum
	SECURE_STORAGE                 AppCategoryEnum
	MULTIMEDIA_AND_CODING          AppCategoryEnum
	PROJECT_MANAGEMENT             AppCategoryEnum
	PRODUCTIVITY_AND_COLLABORATION AppCategoryEnum
	WEB_ADN_APPLICATION            AppCategoryEnum
	GRAPHIC_DESIGN                 AppCategoryEnum
	OTHER                          AppCategoryEnum
}

func GetAppCategoryEnumEnum() AppCategoryEnumEnum {
	return AppCategoryEnumEnum{
		GAME: AppCategoryEnum{
			value: "GAME",
		},
		BUSSINESS_INTELLIGENCE: AppCategoryEnum{
			value: "BUSSINESS_INTELLIGENCE",
		},
		SECURE_STORAGE: AppCategoryEnum{
			value: "SECURE_STORAGE",
		},
		MULTIMEDIA_AND_CODING: AppCategoryEnum{
			value: "MULTIMEDIA_AND_CODING",
		},
		PROJECT_MANAGEMENT: AppCategoryEnum{
			value: "PROJECT_MANAGEMENT",
		},
		PRODUCTIVITY_AND_COLLABORATION: AppCategoryEnum{
			value: "PRODUCTIVITY_AND_COLLABORATION",
		},
		WEB_ADN_APPLICATION: AppCategoryEnum{
			value: "WEB_ADN_APPLICATION",
		},
		GRAPHIC_DESIGN: AppCategoryEnum{
			value: "GRAPHIC_DESIGN",
		},
		OTHER: AppCategoryEnum{
			value: "OTHER",
		},
	}
}

func (c AppCategoryEnum) Value() string {
	return c.value
}

func (c AppCategoryEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppCategoryEnum) UnmarshalJSON(b []byte) error {
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
