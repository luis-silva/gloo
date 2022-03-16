package tg

import (
	"github.com/rotisserie/eris"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	aws2 "github.com/solo-io/gloo/projects/gloo/pkg/utils/aws"
)

func GetELBV2Client(cred *CredentialSpec, secrets v1.SecretList) (*elbv2.ELBV2, error) {
	regionConfig := &aws.Config{Region: aws.String(cred.Region())}
	secretRef := cred.SecretRef()
	sess, err := aws2.GetAwsSession(secretRef, secrets, regionConfig)
	if err != nil {
		if secretRef == nil {
			return nil, CreateSessionFromEnvError(err)
		}
		return nil, CreateSessionFromSecretError(err)
	}
	if cred.Arn() != "" {
		cred := stscreds.NewCredentials(sess, cred.Arn())
		config := &aws.Config{Credentials: cred}
		return elbv2.New(sess, config), nil
	}
	return elbv2.New(sess), nil
}

func GetTargetGroupsFromDescription(desc *elbv2.DescribeTargetGroupsOutput) []*elbv2.TargetGroup {
	var targetGroups []*elbv2.TargetGroup
	return append(targetGroups, desc.TargetGroups...)
}

var (
	CreateSessionFromEnvError = func(err error) error {
		return eris.Wrapf(err, "unable to create a session with credentials taken from env")
	}

	CreateSessionFromSecretError = func(err error) error {
		return eris.Wrapf(err, "unable to create a session with credentials taken from secret ref")
	}
)
