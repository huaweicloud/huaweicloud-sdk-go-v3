package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 添加设备结构体。
type AddDeviceRequestBody struct {

	// 设备标识码，通常使用IMEI、MAC地址或Serial No作为node_id。（注意:NB设备由于模组烧录信息后无法配置，所以NB设备会校验node_id全局唯一。）
	NodeId string `json:"node_id"`

	// 设备名称。
	DeviceName *string `json:"device_name,omitempty"`

	// 设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。
	ProductId string `json:"product_id"`

	AuthInfo *EdgeDeviceAuthInfo `json:"auth_info,omitempty"`

	// 设备的描述信息。
	Description *string `json:"description,omitempty"`

	// 父设备ID，用于标识设备所属的父设备。携带该参数时，表示在该父设备下创建一个子设备，这个子设备不与平台直连，此时必须保证这个父设备在平台已存在，创建成功后子设备的gateway_id等于该参数值；不携带该参数时，表示创建一个和平台直连的设备，创建成功后设备的device_id和gateway_id一致。
	GatewayId *string `json:"gateway_id,omitempty"`

	// 资源空间Id。此参数为非必选参数，用于兼容平台老用户存在多应用的场景。存在多应用的用户需要使用该接口时，必须携带该参数指定注册的设备归属到哪个应用下，否则接口会提示错误。如果用户存在多应用，同时又不想携带该参数，可以联系华为技术支持对用户数据做应用合并。
	SpaceId *string `json:"space_id,omitempty"`

	// 设备扩展信息。用户可以自定义任何想要的扩展信息，如果在创建设备时为子设备指定该字段，将会通过MQTT接口“平台通知网关子设备新增“将该信息通知给网关。字段值大小上限为1K。
	ExtensionInfo *interface{} `json:"extension_info,omitempty"`

	// 设备初始配置。用户使用该字段可以为设备指定初始配置，指定后将会根据service_id和desired设置的属性值与产品中对应属性的默认值比对，如果不同，则将以设置的属性值为准写入到设备配置中。
	Config *interface{} `json:"config,omitempty"`
}

func (o AddDeviceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeviceRequestBody struct{}"
	}

	return strings.Join([]string{"AddDeviceRequestBody", string(data)}, " ")
}
