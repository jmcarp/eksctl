package cloudformation

// AWSOpsWorksLayer_VolumeConfiguration AWS CloudFormation Resource (AWS::OpsWorks::Layer.VolumeConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html
type AWSOpsWorksLayer_VolumeConfiguration struct {

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-iops
	Iops int `json:"Iops,omitempty"`

	// MountPoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-mountpoint
	MountPoint *StringIntrinsic `json:"MountPoint,omitempty"`

	// NumberOfDisks AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-numberofdisks
	NumberOfDisks int `json:"NumberOfDisks,omitempty"`

	// RaidLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-raidlevel
	RaidLevel int `json:"RaidLevel,omitempty"`

	// Size AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-size
	Size int `json:"Size,omitempty"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-volumetype
	VolumeType *StringIntrinsic `json:"VolumeType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_VolumeConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.VolumeConfiguration"
}
