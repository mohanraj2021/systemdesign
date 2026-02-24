package main

import "fmt"

type ChannelInterface interface {
	Notify(videoName string)
	SubscribeInterface
}

type ObserverInterface interface {
	Update(message string)
}

type UploaderInterface interface {
	Uplaod(videoName string, videoURL string)
}

type SubscribeInterface interface {
	Subscribe(user User)
	Unsubscribe(user User)
}

type User struct {
	Name string
}

type Channel struct {
	Name      string
	Users     []ObserverInterface
	VideoName []string
	VideoURL  []string
}

func (c *Channel) Notify(videoName string) {
	for _, u := range c.Users {
		u.Update(fmt.Sprintf("New video uploaded: %s", videoName))
	}
}

func (c *Channel) Uplaod(videoName string, videoURL string) {
	c.VideoName = append(c.VideoName, videoName)
	c.VideoURL = append(c.VideoURL, videoURL)
	fmt.Printf("The video %s has been uploaded\n", videoName)
	c.Notify(videoName)
}

func (c *Channel) Subscribe(user ObserverInterface) {
	c.Users = append(c.Users, user)
	fmt.Printf("You Subscribed to channel %s\n", c.Name)
}

func (c *Channel) Unsubscribe(user ObserverInterface) {
	for i, u := range c.Users {
		if u == user {
			c.Users = append(c.Users[:i], c.Users[i+1:]...)
			break
		}
	}
	fmt.Printf("You Unsubscribed from channel %s\n", c.Name)

}

func (u User) Update(message string) {
	fmt.Printf("Mr.%s %s\n", u.Name, message)
}

func NewChannel(name string) *Channel {
	return &Channel{
		Name:  name,
		Users: make([]ObserverInterface, 0),
	}
}

func main() {
	channel := NewChannel("Golang Channel")
	user1 := User{Name: "Ali"}
	user2 := User{Name: "Ahmed"}

	channel.Subscribe(user1)
	channel.Subscribe(user2)
	videoName := "How to use Observer Pattern in Go"
	channel.Uplaod(videoName, "https://www.dummyvideo.com/watch?v=6r9hQGZzj8o")

	channel.Unsubscribe(user1)
	videoName2 := "How to use Factory Pattern in Go"
	channel.Uplaod(videoName2, "https://www.dummyvideo.com/watch?v=6r9hQGZzj8o")
}
