package enums

type Priority int

const (
	CriticalPriority Priority = 1
	HighPriority     Priority = 2
	MediumPriority   Priority = 3
	LowPriority      Priority = 4
)

type Status string

const (
	PendingStatus   Status = "Pending"
	SentStatus      Status = "Sent"
	FailedStatus    Status = "Failed"
	DeliveredStatus Status = "Delivered"
)

type ChannelType string

const (
	EmailChannelType ChannelType = "Email"
	SMSChannelType   ChannelType = "SMS"
	PushChannelType  ChannelType = "Push"
)

type DecoratorType string

const (
	SimpleDecoratorType DecoratorType = "Simple"
	UrgentDecoratorType DecoratorType = "Urgent"
)
