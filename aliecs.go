package aliecs

import (
	"log"

	"github.com/BurntSushi/toml"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Config struct {
	AccessKey struct {
		ID     string `toml:"id"`
		Secret string `toml:"secret"`
	} `toml:"access_key"`

	ECS struct {
		InstanceID string `toml:"instance_id"`
		Endpoint   string `toml:"endpoint"`
		RegionID   string `toml:"region_id"`
		StopMode   string `toml:"stop_mode"`
	} `toml:"ecs"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func CreateClient(config *Config) (_result *ecs20140526.Client, _err error) {
	ECSconfig := &openapi.Config{
		AccessKeyId:     tea.String(config.AccessKey.ID),
		AccessKeySecret: tea.String(config.AccessKey.Secret),
	}
	ECSconfig.Endpoint = tea.String(config.ECS.Endpoint)
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(ECSconfig)
	return _result, _err
}

func GetStatus(config *Config, client *ecs20140526.Client) (instanceStatusList []*string, err error) {
	request := &ecs20140526.DescribeInstanceStatusRequest{
		RegionId:   tea.String(config.ECS.RegionID),
		InstanceId: []*string{tea.String(config.ECS.InstanceID)},
	}
	response, err := client.DescribeInstanceStatus(request)
	if err != nil {
		log.Fatal("ECS status query failed")
		return nil, err
	}
	if response.Body != nil && response.Body.InstanceStatuses != nil {
		statuses := response.Body.InstanceStatuses.InstanceStatus
		instanceStatusList = make([]*string, len(statuses))

		for i, status := range statuses {
			instanceStatusList[i] = status.Status
		}
	}
	return instanceStatusList, nil
}

func StartInstance(config *Config, client *ecs20140526.Client) (err error) {
	request := &ecs20140526.StartInstanceRequest{
		InstanceId: tea.String(config.ECS.InstanceID),
	}
	_, err = client.StartInstance(request)
	if err != nil {
		log.Fatal("ECS start query failed")
		return err
	}
	return nil
}

func StopInstance(config *Config, client *ecs20140526.Client) (err error) {
	request := &ecs20140526.StopInstanceRequest{
		InstanceId:  tea.String(config.ECS.InstanceID),
		StoppedMode: tea.String(config.ECS.StopMode),
	}
	_, err = client.StopInstance(request)
	if err != nil {
		log.Fatal("ECS stop query failed")
		return err
	}
	return nil
}
