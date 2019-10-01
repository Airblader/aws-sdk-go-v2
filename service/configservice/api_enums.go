// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

type AggregatedSourceStatusType string

// Enum values for AggregatedSourceStatusType
const (
	AggregatedSourceStatusTypeFailed    AggregatedSourceStatusType = "FAILED"
	AggregatedSourceStatusTypeSucceeded AggregatedSourceStatusType = "SUCCEEDED"
	AggregatedSourceStatusTypeOutdated  AggregatedSourceStatusType = "OUTDATED"
)

func (enum AggregatedSourceStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AggregatedSourceStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AggregatedSourceType string

// Enum values for AggregatedSourceType
const (
	AggregatedSourceTypeAccount      AggregatedSourceType = "ACCOUNT"
	AggregatedSourceTypeOrganization AggregatedSourceType = "ORGANIZATION"
)

func (enum AggregatedSourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AggregatedSourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChronologicalOrder string

// Enum values for ChronologicalOrder
const (
	ChronologicalOrderReverse ChronologicalOrder = "Reverse"
	ChronologicalOrderForward ChronologicalOrder = "Forward"
)

func (enum ChronologicalOrder) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChronologicalOrder) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComplianceType string

// Enum values for ComplianceType
const (
	ComplianceTypeCompliant        ComplianceType = "COMPLIANT"
	ComplianceTypeNonCompliant     ComplianceType = "NON_COMPLIANT"
	ComplianceTypeNotApplicable    ComplianceType = "NOT_APPLICABLE"
	ComplianceTypeInsufficientData ComplianceType = "INSUFFICIENT_DATA"
)

func (enum ComplianceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComplianceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConfigRuleComplianceSummaryGroupKey string

// Enum values for ConfigRuleComplianceSummaryGroupKey
const (
	ConfigRuleComplianceSummaryGroupKeyAccountId ConfigRuleComplianceSummaryGroupKey = "ACCOUNT_ID"
	ConfigRuleComplianceSummaryGroupKeyAwsRegion ConfigRuleComplianceSummaryGroupKey = "AWS_REGION"
)

func (enum ConfigRuleComplianceSummaryGroupKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConfigRuleComplianceSummaryGroupKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConfigRuleState string

// Enum values for ConfigRuleState
const (
	ConfigRuleStateActive          ConfigRuleState = "ACTIVE"
	ConfigRuleStateDeleting        ConfigRuleState = "DELETING"
	ConfigRuleStateDeletingResults ConfigRuleState = "DELETING_RESULTS"
	ConfigRuleStateEvaluating      ConfigRuleState = "EVALUATING"
)

func (enum ConfigRuleState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConfigRuleState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConfigurationItemStatus string

// Enum values for ConfigurationItemStatus
const (
	ConfigurationItemStatusOk                         ConfigurationItemStatus = "OK"
	ConfigurationItemStatusResourceDiscovered         ConfigurationItemStatus = "ResourceDiscovered"
	ConfigurationItemStatusResourceNotRecorded        ConfigurationItemStatus = "ResourceNotRecorded"
	ConfigurationItemStatusResourceDeleted            ConfigurationItemStatus = "ResourceDeleted"
	ConfigurationItemStatusResourceDeletedNotRecorded ConfigurationItemStatus = "ResourceDeletedNotRecorded"
)

func (enum ConfigurationItemStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConfigurationItemStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryStatus string

// Enum values for DeliveryStatus
const (
	DeliveryStatusSuccess       DeliveryStatus = "Success"
	DeliveryStatusFailure       DeliveryStatus = "Failure"
	DeliveryStatusNotApplicable DeliveryStatus = "Not_Applicable"
)

func (enum DeliveryStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventSource string

// Enum values for EventSource
const (
	EventSourceAwsConfig EventSource = "aws.config"
)

func (enum EventSource) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventSource) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MaximumExecutionFrequency string

// Enum values for MaximumExecutionFrequency
const (
	MaximumExecutionFrequencyOneHour         MaximumExecutionFrequency = "One_Hour"
	MaximumExecutionFrequencyThreeHours      MaximumExecutionFrequency = "Three_Hours"
	MaximumExecutionFrequencySixHours        MaximumExecutionFrequency = "Six_Hours"
	MaximumExecutionFrequencyTwelveHours     MaximumExecutionFrequency = "Twelve_Hours"
	MaximumExecutionFrequencyTwentyFourHours MaximumExecutionFrequency = "TwentyFour_Hours"
)

func (enum MaximumExecutionFrequency) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MaximumExecutionFrequency) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MemberAccountRuleStatus string

// Enum values for MemberAccountRuleStatus
const (
	MemberAccountRuleStatusCreateSuccessful MemberAccountRuleStatus = "CREATE_SUCCESSFUL"
	MemberAccountRuleStatusCreateInProgress MemberAccountRuleStatus = "CREATE_IN_PROGRESS"
	MemberAccountRuleStatusCreateFailed     MemberAccountRuleStatus = "CREATE_FAILED"
	MemberAccountRuleStatusUpdateSuccessful MemberAccountRuleStatus = "UPDATE_SUCCESSFUL"
	MemberAccountRuleStatusUpdateFailed     MemberAccountRuleStatus = "UPDATE_FAILED"
	MemberAccountRuleStatusUpdateInProgress MemberAccountRuleStatus = "UPDATE_IN_PROGRESS"
	MemberAccountRuleStatusDeleteSuccessful MemberAccountRuleStatus = "DELETE_SUCCESSFUL"
	MemberAccountRuleStatusDeleteFailed     MemberAccountRuleStatus = "DELETE_FAILED"
	MemberAccountRuleStatusDeleteInProgress MemberAccountRuleStatus = "DELETE_IN_PROGRESS"
)

func (enum MemberAccountRuleStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MemberAccountRuleStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeConfigurationItemChangeNotification          MessageType = "ConfigurationItemChangeNotification"
	MessageTypeConfigurationSnapshotDeliveryCompleted       MessageType = "ConfigurationSnapshotDeliveryCompleted"
	MessageTypeScheduledNotification                        MessageType = "ScheduledNotification"
	MessageTypeOversizedConfigurationItemChangeNotification MessageType = "OversizedConfigurationItemChangeNotification"
)

func (enum MessageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MessageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrganizationConfigRuleTriggerType string

// Enum values for OrganizationConfigRuleTriggerType
const (
	OrganizationConfigRuleTriggerTypeConfigurationItemChangeNotification          OrganizationConfigRuleTriggerType = "ConfigurationItemChangeNotification"
	OrganizationConfigRuleTriggerTypeOversizedConfigurationItemChangeNotification OrganizationConfigRuleTriggerType = "OversizedConfigurationItemChangeNotification"
	OrganizationConfigRuleTriggerTypeScheduledNotification                        OrganizationConfigRuleTriggerType = "ScheduledNotification"
)

func (enum OrganizationConfigRuleTriggerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrganizationConfigRuleTriggerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrganizationRuleStatus string

// Enum values for OrganizationRuleStatus
const (
	OrganizationRuleStatusCreateSuccessful OrganizationRuleStatus = "CREATE_SUCCESSFUL"
	OrganizationRuleStatusCreateInProgress OrganizationRuleStatus = "CREATE_IN_PROGRESS"
	OrganizationRuleStatusCreateFailed     OrganizationRuleStatus = "CREATE_FAILED"
	OrganizationRuleStatusUpdateSuccessful OrganizationRuleStatus = "UPDATE_SUCCESSFUL"
	OrganizationRuleStatusUpdateFailed     OrganizationRuleStatus = "UPDATE_FAILED"
	OrganizationRuleStatusUpdateInProgress OrganizationRuleStatus = "UPDATE_IN_PROGRESS"
	OrganizationRuleStatusDeleteSuccessful OrganizationRuleStatus = "DELETE_SUCCESSFUL"
	OrganizationRuleStatusDeleteFailed     OrganizationRuleStatus = "DELETE_FAILED"
	OrganizationRuleStatusDeleteInProgress OrganizationRuleStatus = "DELETE_IN_PROGRESS"
)

func (enum OrganizationRuleStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrganizationRuleStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Owner string

// Enum values for Owner
const (
	OwnerCustomLambda Owner = "CUSTOM_LAMBDA"
	OwnerAws          Owner = "AWS"
)

func (enum Owner) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Owner) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecorderStatus string

// Enum values for RecorderStatus
const (
	RecorderStatusPending RecorderStatus = "Pending"
	RecorderStatusSuccess RecorderStatus = "Success"
	RecorderStatusFailure RecorderStatus = "Failure"
)

func (enum RecorderStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecorderStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RemediationExecutionState string

// Enum values for RemediationExecutionState
const (
	RemediationExecutionStateQueued     RemediationExecutionState = "QUEUED"
	RemediationExecutionStateInProgress RemediationExecutionState = "IN_PROGRESS"
	RemediationExecutionStateSucceeded  RemediationExecutionState = "SUCCEEDED"
	RemediationExecutionStateFailed     RemediationExecutionState = "FAILED"
)

func (enum RemediationExecutionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RemediationExecutionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RemediationExecutionStepState string

// Enum values for RemediationExecutionStepState
const (
	RemediationExecutionStepStateSucceeded RemediationExecutionStepState = "SUCCEEDED"
	RemediationExecutionStepStatePending   RemediationExecutionStepState = "PENDING"
	RemediationExecutionStepStateFailed    RemediationExecutionStepState = "FAILED"
)

func (enum RemediationExecutionStepState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RemediationExecutionStepState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RemediationTargetType string

// Enum values for RemediationTargetType
const (
	RemediationTargetTypeSsmDocument RemediationTargetType = "SSM_DOCUMENT"
)

func (enum RemediationTargetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RemediationTargetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceCountGroupKey string

// Enum values for ResourceCountGroupKey
const (
	ResourceCountGroupKeyResourceType ResourceCountGroupKey = "RESOURCE_TYPE"
	ResourceCountGroupKeyAccountId    ResourceCountGroupKey = "ACCOUNT_ID"
	ResourceCountGroupKeyAwsRegion    ResourceCountGroupKey = "AWS_REGION"
)

func (enum ResourceCountGroupKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceCountGroupKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeAwsEc2CustomerGateway                             ResourceType = "AWS::EC2::CustomerGateway"
	ResourceTypeAwsEc2Eip                                         ResourceType = "AWS::EC2::EIP"
	ResourceTypeAwsEc2Host                                        ResourceType = "AWS::EC2::Host"
	ResourceTypeAwsEc2Instance                                    ResourceType = "AWS::EC2::Instance"
	ResourceTypeAwsEc2InternetGateway                             ResourceType = "AWS::EC2::InternetGateway"
	ResourceTypeAwsEc2NetworkAcl                                  ResourceType = "AWS::EC2::NetworkAcl"
	ResourceTypeAwsEc2NetworkInterface                            ResourceType = "AWS::EC2::NetworkInterface"
	ResourceTypeAwsEc2RouteTable                                  ResourceType = "AWS::EC2::RouteTable"
	ResourceTypeAwsEc2SecurityGroup                               ResourceType = "AWS::EC2::SecurityGroup"
	ResourceTypeAwsEc2Subnet                                      ResourceType = "AWS::EC2::Subnet"
	ResourceTypeAwsCloudTrailTrail                                ResourceType = "AWS::CloudTrail::Trail"
	ResourceTypeAwsEc2Volume                                      ResourceType = "AWS::EC2::Volume"
	ResourceTypeAwsEc2Vpc                                         ResourceType = "AWS::EC2::VPC"
	ResourceTypeAwsEc2Vpnconnection                               ResourceType = "AWS::EC2::VPNConnection"
	ResourceTypeAwsEc2Vpngateway                                  ResourceType = "AWS::EC2::VPNGateway"
	ResourceTypeAwsEc2RegisteredHainstance                        ResourceType = "AWS::EC2::RegisteredHAInstance"
	ResourceTypeAwsEc2NatGateway                                  ResourceType = "AWS::EC2::NatGateway"
	ResourceTypeAwsEc2EgressOnlyInternetGateway                   ResourceType = "AWS::EC2::EgressOnlyInternetGateway"
	ResourceTypeAwsEc2Vpcendpoint                                 ResourceType = "AWS::EC2::VPCEndpoint"
	ResourceTypeAwsEc2VpcendpointService                          ResourceType = "AWS::EC2::VPCEndpointService"
	ResourceTypeAwsEc2FlowLog                                     ResourceType = "AWS::EC2::FlowLog"
	ResourceTypeAwsEc2VpcpeeringConnection                        ResourceType = "AWS::EC2::VPCPeeringConnection"
	ResourceTypeAwsIamGroup                                       ResourceType = "AWS::IAM::Group"
	ResourceTypeAwsIamPolicy                                      ResourceType = "AWS::IAM::Policy"
	ResourceTypeAwsIamRole                                        ResourceType = "AWS::IAM::Role"
	ResourceTypeAwsIamUser                                        ResourceType = "AWS::IAM::User"
	ResourceTypeAwsElasticLoadBalancingV2LoadBalancer             ResourceType = "AWS::ElasticLoadBalancingV2::LoadBalancer"
	ResourceTypeAwsAcmCertificate                                 ResourceType = "AWS::ACM::Certificate"
	ResourceTypeAwsRdsDbinstance                                  ResourceType = "AWS::RDS::DBInstance"
	ResourceTypeAwsRdsDbparameterGroup                            ResourceType = "AWS::RDS::DBParameterGroup"
	ResourceTypeAwsRdsDboptionGroup                               ResourceType = "AWS::RDS::DBOptionGroup"
	ResourceTypeAwsRdsDbsubnetGroup                               ResourceType = "AWS::RDS::DBSubnetGroup"
	ResourceTypeAwsRdsDbsecurityGroup                             ResourceType = "AWS::RDS::DBSecurityGroup"
	ResourceTypeAwsRdsDbsnapshot                                  ResourceType = "AWS::RDS::DBSnapshot"
	ResourceTypeAwsRdsDbcluster                                   ResourceType = "AWS::RDS::DBCluster"
	ResourceTypeAwsRdsDbclusterParameterGroup                     ResourceType = "AWS::RDS::DBClusterParameterGroup"
	ResourceTypeAwsRdsDbclusterSnapshot                           ResourceType = "AWS::RDS::DBClusterSnapshot"
	ResourceTypeAwsRdsEventSubscription                           ResourceType = "AWS::RDS::EventSubscription"
	ResourceTypeAwsS3Bucket                                       ResourceType = "AWS::S3::Bucket"
	ResourceTypeAwsS3AccountPublicAccessBlock                     ResourceType = "AWS::S3::AccountPublicAccessBlock"
	ResourceTypeAwsRedshiftCluster                                ResourceType = "AWS::Redshift::Cluster"
	ResourceTypeAwsRedshiftClusterSnapshot                        ResourceType = "AWS::Redshift::ClusterSnapshot"
	ResourceTypeAwsRedshiftClusterParameterGroup                  ResourceType = "AWS::Redshift::ClusterParameterGroup"
	ResourceTypeAwsRedshiftClusterSecurityGroup                   ResourceType = "AWS::Redshift::ClusterSecurityGroup"
	ResourceTypeAwsRedshiftClusterSubnetGroup                     ResourceType = "AWS::Redshift::ClusterSubnetGroup"
	ResourceTypeAwsRedshiftEventSubscription                      ResourceType = "AWS::Redshift::EventSubscription"
	ResourceTypeAwsSsmManagedInstanceInventory                    ResourceType = "AWS::SSM::ManagedInstanceInventory"
	ResourceTypeAwsCloudWatchAlarm                                ResourceType = "AWS::CloudWatch::Alarm"
	ResourceTypeAwsCloudFormationStack                            ResourceType = "AWS::CloudFormation::Stack"
	ResourceTypeAwsElasticLoadBalancingLoadBalancer               ResourceType = "AWS::ElasticLoadBalancing::LoadBalancer"
	ResourceTypeAwsAutoScalingAutoScalingGroup                    ResourceType = "AWS::AutoScaling::AutoScalingGroup"
	ResourceTypeAwsAutoScalingLaunchConfiguration                 ResourceType = "AWS::AutoScaling::LaunchConfiguration"
	ResourceTypeAwsAutoScalingScalingPolicy                       ResourceType = "AWS::AutoScaling::ScalingPolicy"
	ResourceTypeAwsAutoScalingScheduledAction                     ResourceType = "AWS::AutoScaling::ScheduledAction"
	ResourceTypeAwsDynamoDbTable                                  ResourceType = "AWS::DynamoDB::Table"
	ResourceTypeAwsCodeBuildProject                               ResourceType = "AWS::CodeBuild::Project"
	ResourceTypeAwsWafRateBasedRule                               ResourceType = "AWS::WAF::RateBasedRule"
	ResourceTypeAwsWafRule                                        ResourceType = "AWS::WAF::Rule"
	ResourceTypeAwsWafRuleGroup                                   ResourceType = "AWS::WAF::RuleGroup"
	ResourceTypeAwsWafWebAcl                                      ResourceType = "AWS::WAF::WebACL"
	ResourceTypeAwsWafregionalRateBasedRule                       ResourceType = "AWS::WAFRegional::RateBasedRule"
	ResourceTypeAwsWafregionalRule                                ResourceType = "AWS::WAFRegional::Rule"
	ResourceTypeAwsWafregionalRuleGroup                           ResourceType = "AWS::WAFRegional::RuleGroup"
	ResourceTypeAwsWafregionalWebAcl                              ResourceType = "AWS::WAFRegional::WebACL"
	ResourceTypeAwsCloudFrontDistribution                         ResourceType = "AWS::CloudFront::Distribution"
	ResourceTypeAwsCloudFrontStreamingDistribution                ResourceType = "AWS::CloudFront::StreamingDistribution"
	ResourceTypeAwsLambdaAlias                                    ResourceType = "AWS::Lambda::Alias"
	ResourceTypeAwsLambdaFunction                                 ResourceType = "AWS::Lambda::Function"
	ResourceTypeAwsElasticBeanstalkApplication                    ResourceType = "AWS::ElasticBeanstalk::Application"
	ResourceTypeAwsElasticBeanstalkApplicationVersion             ResourceType = "AWS::ElasticBeanstalk::ApplicationVersion"
	ResourceTypeAwsElasticBeanstalkEnvironment                    ResourceType = "AWS::ElasticBeanstalk::Environment"
	ResourceTypeAwsMobileHubProject                               ResourceType = "AWS::MobileHub::Project"
	ResourceTypeAwsXrayEncryptionConfig                           ResourceType = "AWS::XRay::EncryptionConfig"
	ResourceTypeAwsSsmAssociationCompliance                       ResourceType = "AWS::SSM::AssociationCompliance"
	ResourceTypeAwsSsmPatchCompliance                             ResourceType = "AWS::SSM::PatchCompliance"
	ResourceTypeAwsShieldProtection                               ResourceType = "AWS::Shield::Protection"
	ResourceTypeAwsShieldRegionalProtection                       ResourceType = "AWS::ShieldRegional::Protection"
	ResourceTypeAwsConfigResourceCompliance                       ResourceType = "AWS::Config::ResourceCompliance"
	ResourceTypeAwsLicenseManagerLicenseConfiguration             ResourceType = "AWS::LicenseManager::LicenseConfiguration"
	ResourceTypeAwsApiGatewayDomainName                           ResourceType = "AWS::ApiGateway::DomainName"
	ResourceTypeAwsApiGatewayMethod                               ResourceType = "AWS::ApiGateway::Method"
	ResourceTypeAwsApiGatewayStage                                ResourceType = "AWS::ApiGateway::Stage"
	ResourceTypeAwsApiGatewayRestApi                              ResourceType = "AWS::ApiGateway::RestApi"
	ResourceTypeAwsApiGatewayV2DomainName                         ResourceType = "AWS::ApiGatewayV2::DomainName"
	ResourceTypeAwsApiGatewayV2Stage                              ResourceType = "AWS::ApiGatewayV2::Stage"
	ResourceTypeAwsApiGatewayV2Api                                ResourceType = "AWS::ApiGatewayV2::Api"
	ResourceTypeAwsCodePipelinePipeline                           ResourceType = "AWS::CodePipeline::Pipeline"
	ResourceTypeAwsServiceCatalogCloudFormationProvisionedProduct ResourceType = "AWS::ServiceCatalog::CloudFormationProvisionedProduct"
	ResourceTypeAwsServiceCatalogCloudFormationProduct            ResourceType = "AWS::ServiceCatalog::CloudFormationProduct"
	ResourceTypeAwsServiceCatalogPortfolio                        ResourceType = "AWS::ServiceCatalog::Portfolio"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceValueType string

// Enum values for ResourceValueType
const (
	ResourceValueTypeResourceId ResourceValueType = "RESOURCE_ID"
)

func (enum ResourceValueType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceValueType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}