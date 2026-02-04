package adapter

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
)

type RdsInstance struct {
	Name               string
	Engine             string
	EngineVersion      string
	InstanceClass      string
	StorageType        string
	StorageGB          int32
	MultiAZ            bool
	PubliclyAccessible bool
	MasterUserName     string
	InstanceSG         string
}

func AWSInstanceToConfig(instance rdsTypes.DBInstance) RdsInstance {
	return RdsInstance{
		Name:               aws.ToString(instance.DBInstanceIdentifier),
		Engine:             aws.ToString(instance.Engine),
		EngineVersion:      aws.ToString(instance.EngineVersion),
		InstanceClass:      aws.ToString(instance.DBInstanceClass),
		StorageType:        aws.ToString(instance.StorageType),
		StorageGB:          aws.ToInt32(instance.AllocatedStorage),
		MultiAZ:            aws.ToBool(instance.MultiAZ),
		PubliclyAccessible: aws.ToBool(instance.PubliclyAccessible),
		MasterUserName:     aws.ToString(instance.MasterUsername),
		InstanceSG:         aws.ToString(instance.VpcSecurityGroups[0].VpcSecurityGroupId),
	}
}

func BatchAwsToDomain(srcList []rdsTypes.DBInstance) []RdsInstance {
	result := make([]RdsInstance, len(srcList))
	for i, v := range srcList {
		result[i] = AWSInstanceToConfig(v)
	}
	return result
}
