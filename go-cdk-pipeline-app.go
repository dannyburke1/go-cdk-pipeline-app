package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type GoCdkPipelineAppStackProps struct {
	awscdk.StackProps
}

func NewGoCdkPipelineAppStack(scope constructs.Construct, id string, props *GoCdkPipelineAppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	var timeStamp = time.Now()
	var bucket = awss3.NewBucket(stack, jsii.String("CDKPipelinesGoBucket"+timeStamp.String()), &awss3.BucketProps{
		Versioned: jsii.Bool(false),
	})

	if bucket != nil {
		fmt.Printf("Bucket Object incorrectly called")
	}

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewGoCdkPipelineAppStack(app, "GoCdkPipelineAppStack", &GoCdkPipelineAppStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
