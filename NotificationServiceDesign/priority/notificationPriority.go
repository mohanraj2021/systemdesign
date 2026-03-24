package priority

import "plugplay.com/notification/enums"

type NotificationChannel interface {
	Send() error
	GetChannelType() string
}

type NotificationPriority struct {
	ID        int
	OrderId   int
	Priority  enums.Priority
	Timestamp int
	Type      string
	Message   string
	Status    enums.Status
	Channels  []enums.ChannelType
}

type NotificationPriorityList []NotificationPriority

func (npl NotificationPriorityList) Len() int {
	return len(npl)
}

func (npl NotificationPriorityList) Less(i, j int) bool {
	return npl[i].Priority < npl[j].Priority
}

func (npl NotificationPriorityList) Swap(i, j int) {
	npl[i], npl[j] = npl[j], npl[i]
}

func (npl *NotificationPriorityList) Pop() interface{} {
	n := (*npl)[len(*npl)-1]
	*npl = (*npl)[:len(*npl)-1]
	return n
}

func (npl *NotificationPriorityList) Push(x interface{}) {
	*npl = append(*npl, x.(NotificationPriority))
}
