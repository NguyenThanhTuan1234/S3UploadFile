package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AWSConfigProvider struct {
	Profile string
}

func (p *AWSConfigProvider) LoadAWSConfig() (aws.Config, error) {
	if p.Profile == "" {
		return config.LoadDefaultConfig()
	}
	return config.LoadDefaultConfig(
		config.WithSharedConfigFiles([]string{p.Profile}),
	)
}
