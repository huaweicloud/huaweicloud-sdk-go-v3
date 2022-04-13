package v5

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/model"
)

type IoTDAClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIoTDAClient(hcClient *http_client.HcHttpClient) *IoTDAClient {
	return &IoTDAClient{HcClient: hcClient}
}

func IoTDAClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithDerivedAuthServiceName("iotdm")
	return builder
}

//接入凭证是用于客户端使用AMQP等协议与平台建链的一个认证凭据。只保留一条记录，如果重复调用只会重置接入凭证，使得之前的失效。
func (c *IoTDAClient) CreateAccessCode(request *model.CreateAccessCodeRequest) (*model.CreateAccessCodeResponse, error) {
	requestDef := GenReqDefForCreateAccessCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessCodeResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台创建一个AMQP队列。每个租户只能创建100个队列，若超过规格，则创建失败，若队列名称与已有的队列名称相同，则创建失败。
func (c *IoTDAClient) AddQueue(request *model.AddQueueRequest) (*model.AddQueueResponse, error) {
	requestDef := GenReqDefForAddQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddQueueResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中的AMQP队列信息列表。可通过队列名称作模糊查询，支持分页。
func (c *IoTDAClient) BatchShowQueue(request *model.BatchShowQueueRequest) (*model.BatchShowQueueResponse, error) {
	requestDef := GenReqDefForBatchShowQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowQueueResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台上删除指定AMQP队列。若当前队列正在使用，则会删除失败。
func (c *IoTDAClient) DeleteQueue(request *model.DeleteQueueRequest) (*model.DeleteQueueResponse, error) {
	requestDef := GenReqDefForDeleteQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueueResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中指定队列的详细信息。
func (c *IoTDAClient) ShowQueue(request *model.ShowQueueRequest) (*model.ShowQueueResponse, error) {
	requestDef := GenReqDefForShowQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueueResponse), nil
	}
}

//资源空间对应的是物联网平台原有的应用，在物联网平台的含义与应用一致，只是变更了名称。应用服务器可以调用此接口创建资源空间。
func (c *IoTDAClient) AddApplication(request *model.AddApplicationRequest) (*model.AddApplicationResponse, error) {
	requestDef := GenReqDefForAddApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddApplicationResponse), nil
	}
}

//删除指定资源空间。删除资源空间属于高危操作，删除资源空间后，该空间下的产品、设备等资源将不可用，请谨慎操作！
func (c *IoTDAClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

//资源空间对应的是物联网平台原有的应用，在物联网平台的含义与应用一致，只是变更了名称。应用服务器可以调用此接口查询指定资源空间详情。
func (c *IoTDAClient) ShowApplication(request *model.ShowApplicationRequest) (*model.ShowApplicationResponse, error) {
	requestDef := GenReqDefForShowApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationResponse), nil
	}
}

//资源空间对应的是物联网平台原有的应用，在物联网平台的含义与应用一致，只是变更了名称。应用服务器可以调用此接口查询资源空间列表。
func (c *IoTDAClient) ShowApplications(request *model.ShowApplicationsRequest) (*model.ShowApplicationsResponse, error) {
	requestDef := GenReqDefForShowApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationsResponse), nil
	}
}

//设备的产品模型中定义了物联网平台可向设备下发的命令，应用服务器可调用此接口向指定设备下发异步命令，以实现对设备的控制。平台负责将命令发送给设备，并将设备执行命令结果异步通知应用服务器。 命令执行结果支持灵活的数据流转，应用服务器通过调用物联网平台的创建规则触发条件（Resource:device.command.status，Event:update）、创建规则动作并激活规则后，当命令状态变更时，物联网平台会根据规则将结果发送到规则指定的服务器，如用户自定义的HTTP服务器，AMQP服务器，以及华为云的其他储存服务器等, 详情参考[设备命令状态变更通知](https://support.huaweicloud.com/api-iothub/iot_06_v5_01212.html)。注意：此接口适用于NB设备异步命令下发，暂不支持其他协议类型设备命令下发。
func (c *IoTDAClient) CreateAsyncCommand(request *model.CreateAsyncCommandRequest) (*model.CreateAsyncCommandResponse, error) {
	requestDef := GenReqDefForCreateAsyncCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAsyncCommandResponse), nil
	}
}

//物联网平台可查询指定id的命令。
func (c *IoTDAClient) ShowAsyncDeviceCommand(request *model.ShowAsyncDeviceCommandRequest) (*model.ShowAsyncDeviceCommandResponse, error) {
	requestDef := GenReqDefForShowAsyncDeviceCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAsyncDeviceCommandResponse), nil
	}
}

//应用服务器可调用此接口为创建批量处理任务，对多个设备进行批量操作。当前支持批量软固件升级、批量创建设备、批量删除设备、批量冻结设备、批量解冻设备、批量创建命令、批量创建消息任务。
func (c *IoTDAClient) CreateBatchTask(request *model.CreateBatchTaskRequest) (*model.CreateBatchTaskResponse, error) {
	requestDef := GenReqDefForCreateBatchTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchTaskResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中批量任务列表，每一个任务又包括具体的任务内容、任务状态、任务完成情况统计等。
func (c *IoTDAClient) ListBatchTasks(request *model.ListBatchTasksRequest) (*model.ListBatchTasksResponse, error) {
	requestDef := GenReqDefForListBatchTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBatchTasksResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中指定批量任务的信息，包括任务内容、任务状态、任务完成情况统计以及子任务列表等。
func (c *IoTDAClient) ShowBatchTask(request *model.ShowBatchTaskRequest) (*model.ShowBatchTaskResponse, error) {
	requestDef := GenReqDefForShowBatchTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchTaskResponse), nil
	}
}

//应用服务器可调用此接口删除批量任务文件。
func (c *IoTDAClient) DeleteBatchTaskFile(request *model.DeleteBatchTaskFileRequest) (*model.DeleteBatchTaskFileResponse, error) {
	requestDef := GenReqDefForDeleteBatchTaskFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBatchTaskFileResponse), nil
	}
}

//应用服务器可调用此接口查询批量任务文件列表。
func (c *IoTDAClient) ListBatchTaskFiles(request *model.ListBatchTaskFilesRequest) (*model.ListBatchTaskFilesResponse, error) {
	requestDef := GenReqDefForListBatchTaskFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBatchTaskFilesResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台上传设备的CA证书
func (c *IoTDAClient) AddCertificate(request *model.AddCertificateRequest) (*model.AddCertificateResponse, error) {
	requestDef := GenReqDefForAddCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCertificateResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台验证设备的CA证书，目的是为了验证用户持有设备CA证书的私钥
func (c *IoTDAClient) CheckCertificate(request *model.CheckCertificateRequest) (*model.CheckCertificateResponse, error) {
	requestDef := GenReqDefForCheckCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckCertificateResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台删除设备的CA证书
func (c *IoTDAClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台获取设备的CA证书列表
func (c *IoTDAClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

//设备的产品模型中定义了物联网平台可向设备下发的命令，应用服务器可调用此接口向指定设备下发命令，以实现对设备的同步控制。平台负责将命令以同步方式发送给设备，并将设备执行命令结果同步返回, 如果设备没有响应，平台会返回给应用服务器超时，平台超时间是20秒。注意：此接口适用于MQTT设备同步命令下发，暂不支持NB-IoT设备命令下发。
func (c *IoTDAClient) CreateCommand(request *model.CreateCommandRequest) (*model.CreateCommandResponse, error) {
	requestDef := GenReqDefForCreateCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCommandResponse), nil
	}
}

//应用服务器可调用此接口新建设备组，一个华为云账号下最多可有1,000个分组，包括父分组和子分组。设备组的最大层级关系不超过5层，即群组形成的关系树最大深度不超过5。
func (c *IoTDAClient) AddDeviceGroup(request *model.AddDeviceGroupRequest) (*model.AddDeviceGroupResponse, error) {
	requestDef := GenReqDefForAddDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeviceGroupResponse), nil
	}
}

//应用服务器可调用此接口管理设备组中的设备。单个设备组内最多添加20,000个设备，一个设备最多可以被添加到10个设备组中。
func (c *IoTDAClient) CreateOrDeleteDeviceInGroup(request *model.CreateOrDeleteDeviceInGroupRequest) (*model.CreateOrDeleteDeviceInGroupResponse, error) {
	requestDef := GenReqDefForCreateOrDeleteDeviceInGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrDeleteDeviceInGroupResponse), nil
	}
}

//应用服务器可调用此接口删除指定设备组，如果该设备组存在子设备组或者该设备组中存在设备，必须先删除子设备组并将设备从该设备组移除，才能删除该设备组。
func (c *IoTDAClient) DeleteDeviceGroup(request *model.DeleteDeviceGroupRequest) (*model.DeleteDeviceGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceGroupResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中的设备组信息列表。
func (c *IoTDAClient) ListDeviceGroups(request *model.ListDeviceGroupsRequest) (*model.ListDeviceGroupsResponse, error) {
	requestDef := GenReqDefForListDeviceGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeviceGroupsResponse), nil
	}
}

//应用服务器可调用此接口查询指定设备组详情。
func (c *IoTDAClient) ShowDeviceGroup(request *model.ShowDeviceGroupRequest) (*model.ShowDeviceGroupResponse, error) {
	requestDef := GenReqDefForShowDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceGroupResponse), nil
	}
}

//应用服务器可调用此接口查询指定设备组下的设备列表。
func (c *IoTDAClient) ShowDevicesInGroup(request *model.ShowDevicesInGroupRequest) (*model.ShowDevicesInGroupResponse, error) {
	requestDef := GenReqDefForShowDevicesInGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDevicesInGroupResponse), nil
	}
}

//应用服务器可调用此接口修改物联网平台中指定设备组。
func (c *IoTDAClient) UpdateDeviceGroup(request *model.UpdateDeviceGroupRequest) (*model.UpdateDeviceGroupResponse, error) {
	requestDef := GenReqDefForUpdateDeviceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceGroupResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台创建一个设备，仅在创建后设备才可以接入物联网平台。  - 该接口支持使用gateway_id参数指定在父设备下创建一个子设备，并且支持多级子设备，当前最大支持二级子设备。 - 该接口同时还支持对设备进行初始配置，接口会读取创建设备请求参数product_id对应的产品详情，如果产品的属性有定义默认值，则会将该属性默认值写入该设备的设备影子中。 - 用户还可以使用创建设备请求参数shadow字段为设备指定初始配置，指定后将会根据service_id和desired设置的属性值与产品中对应属性的默认值比对，如果不同，则将以shadow字段中设置的属性值为准写入到设备影子中。
func (c *IoTDAClient) AddDevice(request *model.AddDeviceRequest) (*model.AddDeviceResponse, error) {
	requestDef := GenReqDefForAddDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeviceResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台上删除指定设备。若设备下连接了非直连设备，则必须把设备下的非直连设备都删除后，才能删除该设备。
func (c *IoTDAClient) DeleteDevice(request *model.DeleteDeviceRequest) (*model.DeleteDeviceResponse, error) {
	requestDef := GenReqDefForDeleteDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeviceResponse), nil
	}
}

//应用服务器可调用此接口冻结设备，设备冻结后不能再连接上线，可以通过解冻设备接口解除设备冻结。注意，当前仅支持冻结与平台直连的设备。
func (c *IoTDAClient) FreezeDevice(request *model.FreezeDeviceRequest) (*model.FreezeDeviceResponse, error) {
	requestDef := GenReqDefForFreezeDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezeDeviceResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中的设备信息列表。
func (c *IoTDAClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

//应用服务器可调用此接口重置设备密钥，携带指定密钥时平台将设备密钥重置为指定的密钥，不携带密钥时平台将自动生成一个新的随机密钥返回。
func (c *IoTDAClient) ResetDeviceSecret(request *model.ResetDeviceSecretRequest) (*model.ResetDeviceSecretResponse, error) {
	requestDef := GenReqDefForResetDeviceSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetDeviceSecretResponse), nil
	}
}

//应用服务器可调用此接口重置设备指纹。携带指定设备指纹时将之重置为指定值；不携带时将之置空，后续设备第一次接入时，该设备指纹的值将设置为第一次接入时的证书指纹。
func (c *IoTDAClient) ResetFingerprint(request *model.ResetFingerprintRequest) (*model.ResetFingerprintResponse, error) {
	requestDef := GenReqDefForResetFingerprint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetFingerprintResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中指定设备的详细信息。
func (c *IoTDAClient) ShowDevice(request *model.ShowDeviceRequest) (*model.ShowDeviceResponse, error) {
	requestDef := GenReqDefForShowDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceResponse), nil
	}
}

//应用服务器可调用此接口解冻设备，解除冻结后，设备可以连接上线。
func (c *IoTDAClient) UnfreezeDevice(request *model.UnfreezeDeviceRequest) (*model.UnfreezeDeviceResponse, error) {
	requestDef := GenReqDefForUnfreezeDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezeDeviceResponse), nil
	}
}

//应用服务器可调用此接口修改物联网平台中指定设备的基本信息。
func (c *IoTDAClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

//应用服务器可调用此接口查询指定设备的设备影子信息，包括对设备的期望属性信息（desired区）和设备最新上报的属性信息（reported区）。  设备影子介绍： 设备影子是一个用于存储和检索设备当前状态信息的JSON文档。 - 每个设备有且只有一个设备影子，由设备ID唯一标识 - 设备影子仅保存最近一次设备的上报数据和预期数据 - 无论该设备是否在线，都可以通过该影子获取和设置设备的属性 - 设备上线或者设备上报属性时，如果desired区和reported区存在差异，则将差异部分下发给设备，配置的预期属性需在产品模型中定义且method具有可写属性“W”才可下发
func (c *IoTDAClient) ShowDeviceShadow(request *model.ShowDeviceShadowRequest) (*model.ShowDeviceShadowResponse, error) {
	requestDef := GenReqDefForShowDeviceShadow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceShadowResponse), nil
	}
}

//应用服务器可调用此接口配置设备影子的预期属性（desired区），当设备上线或者设备上报属性时把属性下发给设备。  设备影子介绍： 设备影子是一个用于存储和检索设备当前状态信息的JSON文档。 - 每个设备有且只有一个设备影子，由设备ID唯一标识 - 设备影子仅保存最近一次设备的上报数据和预期数据 - 无论该设备是否在线，都可以通过该影子获取和设置设备的属性 - 设备上线或者设备上报属性时，如果desired区和reported区存在差异，则将差异部分下发给设备，配置的预期属性需在产品模型中定义且method具有可写属性“W”才可下发
func (c *IoTDAClient) UpdateDeviceShadowDesiredData(request *model.UpdateDeviceShadowDesiredDataRequest) (*model.UpdateDeviceShadowDesiredDataResponse, error) {
	requestDef := GenReqDefForUpdateDeviceShadowDesiredData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceShadowDesiredDataResponse), nil
	}
}

//物联网平台可向设备下发消息，应用服务器可调用此接口向指定设备下发消息，以实现对设备的控制。应用将消息下发给平台后，平台返回应用响应结果，平台再将消息发送给设备。注意：此接口适用于MQTT设备消息下发，暂不支持其他协议接入的设备消息下发。
func (c *IoTDAClient) CreateMessage(request *model.CreateMessageRequest) (*model.CreateMessageResponse, error) {
	requestDef := GenReqDefForCreateMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessageResponse), nil
	}
}

//物联网平台可查询指定设备下发的消息，平台为每个设备默认最多保存20条消息，超过20条后， 后续的消息会替换下发最早的消息。
func (c *IoTDAClient) ListDeviceMessages(request *model.ListDeviceMessagesRequest) (*model.ListDeviceMessagesResponse, error) {
	requestDef := GenReqDefForListDeviceMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeviceMessagesResponse), nil
	}
}

//物联网平台可查询设备下发的指定消息id的消息。
func (c *IoTDAClient) ShowDeviceMessage(request *model.ShowDeviceMessageRequest) (*model.ShowDeviceMessageResponse, error) {
	requestDef := GenReqDefForShowDeviceMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceMessageResponse), nil
	}
}

//应用服务器可调用此接口创建产品。
func (c *IoTDAClient) CreateProduct(request *model.CreateProductRequest) (*model.CreateProductResponse, error) {
	requestDef := GenReqDefForCreateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProductResponse), nil
	}
}

//应用服务器可调用此接口删除已导入物联网平台的指定产品模型。
func (c *IoTDAClient) DeleteProduct(request *model.DeleteProductRequest) (*model.DeleteProductResponse, error) {
	requestDef := GenReqDefForDeleteProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProductResponse), nil
	}
}

//应用服务器可调用此接口查询已导入物联网平台的产品模型信息列表，了解产品模型的概要信息。
func (c *IoTDAClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

//应用服务器可调用此接口查询已导入物联网平台的指定产品模型详细信息，包括产品模型的服务、属性、命令等。
func (c *IoTDAClient) ShowProduct(request *model.ShowProductRequest) (*model.ShowProductResponse, error) {
	requestDef := GenReqDefForShowProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProductResponse), nil
	}
}

//应用服务器可调用此接口修改已导入物联网平台的指定产品模型，包括产品模型的服务、属性、命令等。
func (c *IoTDAClient) UpdateProduct(request *model.UpdateProductRequest) (*model.UpdateProductResponse, error) {
	requestDef := GenReqDefForUpdateProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProductResponse), nil
	}
}

//设备的产品模型中定义了物联网平台可向设备下发的属性，应用服务器可调用此接口向设备发送指令用以查询设备的实时属性, 并由设备将属性查询的结果同步返回给应用服务器。注意：此接口适用于MQTT设备，暂不支持NB-IoT设备。
func (c *IoTDAClient) ListProperties(request *model.ListPropertiesRequest) (*model.ListPropertiesResponse, error) {
	requestDef := GenReqDefForListProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPropertiesResponse), nil
	}
}

//设备的产品模型中定义了物联网平台可向设备下发的属性，应用服务器可调用此接口向指定设备下属性。平台负责将属性以同步方式发送给设备，并将设备执行属性结果同步返回。注意：此接口适用于MQTT设备，暂不支持NB-IoT设备。
func (c *IoTDAClient) UpdateProperties(request *model.UpdatePropertiesRequest) (*model.UpdatePropertiesResponse, error) {
	requestDef := GenReqDefForUpdateProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePropertiesResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台创建一条规则触发条件。
func (c *IoTDAClient) CreateRoutingRule(request *model.CreateRoutingRuleRequest) (*model.CreateRoutingRuleResponse, error) {
	requestDef := GenReqDefForCreateRoutingRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRoutingRuleResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台创建一条规则动作。
func (c *IoTDAClient) CreateRuleAction(request *model.CreateRuleActionRequest) (*model.CreateRuleActionResponse, error) {
	requestDef := GenReqDefForCreateRuleAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleActionResponse), nil
	}
}

//应用服务器可调用此接口删除物联网平台中的指定规则条件。
func (c *IoTDAClient) DeleteRoutingRule(request *model.DeleteRoutingRuleRequest) (*model.DeleteRoutingRuleResponse, error) {
	requestDef := GenReqDefForDeleteRoutingRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRoutingRuleResponse), nil
	}
}

//应用服务器可调用此接口删除物联网平台中的指定规则动作。
func (c *IoTDAClient) DeleteRuleAction(request *model.DeleteRuleActionRequest) (*model.DeleteRuleActionResponse, error) {
	requestDef := GenReqDefForDeleteRuleAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleActionResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中设置的规则条件列表。
func (c *IoTDAClient) ListRoutingRules(request *model.ListRoutingRulesRequest) (*model.ListRoutingRulesResponse, error) {
	requestDef := GenReqDefForListRoutingRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoutingRulesResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中设置的规则动作列表。
func (c *IoTDAClient) ListRuleActions(request *model.ListRuleActionsRequest) (*model.ListRuleActionsResponse, error) {
	requestDef := GenReqDefForListRuleActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleActionsResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中指定规则条件的配置信息。
func (c *IoTDAClient) ShowRoutingRule(request *model.ShowRoutingRuleRequest) (*model.ShowRoutingRuleResponse, error) {
	requestDef := GenReqDefForShowRoutingRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRoutingRuleResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中指定规则动作的配置信息。
func (c *IoTDAClient) ShowRuleAction(request *model.ShowRuleActionRequest) (*model.ShowRuleActionResponse, error) {
	requestDef := GenReqDefForShowRuleAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleActionResponse), nil
	}
}

//应用服务器可调用此接口修改物联网平台中指定规则条件的配置参数。
func (c *IoTDAClient) UpdateRoutingRule(request *model.UpdateRoutingRuleRequest) (*model.UpdateRoutingRuleResponse, error) {
	requestDef := GenReqDefForUpdateRoutingRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoutingRuleResponse), nil
	}
}

//应用服务器可调用此接口修改物联网平台中指定规则动作的配置。
func (c *IoTDAClient) UpdateRuleAction(request *model.UpdateRuleActionRequest) (*model.UpdateRuleActionResponse, error) {
	requestDef := GenReqDefForUpdateRuleAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleActionResponse), nil
	}
}

//应用服务器可调用此接口修改物联网平台中指定规则的状态，激活或者去激活规则。
func (c *IoTDAClient) ChangeRuleStatus(request *model.ChangeRuleStatusRequest) (*model.ChangeRuleStatusResponse, error) {
	requestDef := GenReqDefForChangeRuleStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeRuleStatusResponse), nil
	}
}

//应用服务器可调用此接口在物联网平台创建一条规则。
func (c *IoTDAClient) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
	requestDef := GenReqDefForCreateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRuleResponse), nil
	}
}

//应用服务器可调用此接口删除物联网平台中的指定规则。
func (c *IoTDAClient) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
	requestDef := GenReqDefForDeleteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRuleResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中设置的规则列表。
func (c *IoTDAClient) ListRules(request *model.ListRulesRequest) (*model.ListRulesResponse, error) {
	requestDef := GenReqDefForListRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRulesResponse), nil
	}
}

//应用服务器可调用此接口查询物联网平台中指定规则的配置信息。
func (c *IoTDAClient) ShowRule(request *model.ShowRuleRequest) (*model.ShowRuleResponse, error) {
	requestDef := GenReqDefForShowRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRuleResponse), nil
	}
}

//应用服务器可调用此接口修改物联网平台中指定规则的配置。
func (c *IoTDAClient) UpdateRule(request *model.UpdateRuleRequest) (*model.UpdateRuleResponse, error) {
	requestDef := GenReqDefForUpdateRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRuleResponse), nil
	}
}

//应用服务器可调用此接口查询绑定了指定标签的资源。当前支持标签的资源有Device(设备)。
func (c *IoTDAClient) ListResourcesByTags(request *model.ListResourcesByTagsRequest) (*model.ListResourcesByTagsResponse, error) {
	requestDef := GenReqDefForListResourcesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesByTagsResponse), nil
	}
}

//应用服务器可调用此接口为指定资源绑定标签。当前支持标签的资源有Device(设备)。
func (c *IoTDAClient) TagDevice(request *model.TagDeviceRequest) (*model.TagDeviceResponse, error) {
	requestDef := GenReqDefForTagDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagDeviceResponse), nil
	}
}

//应用服务器可调用此接口为指定资源解绑标签。当前支持标签的资源有Device(设备)。
func (c *IoTDAClient) UntagDevice(request *model.UntagDeviceRequest) (*model.UntagDeviceResponse, error) {
	requestDef := GenReqDefForUntagDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UntagDeviceResponse), nil
	}
}
