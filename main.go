package main

import (
	"fmt"
)

type Observer interface {
	Update(messages []Message)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
	SendMessage(sender string, content string)
}

type ChatRoom struct {
	observers []Observer
	messages  []Message
}

type Message struct {
	Sender  string
	Content string
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) RegisterObserver(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ChatRoom) RemoveObserver(observer Observer) {
	for i, obs := range c.observers {
		if obs == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *ChatRoom) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update(c.messages)
	}
}

func (c *ChatRoom) SendMessage(sender string, content string) {
	c.messages = append(c.messages, Message{Sender: sender, Content: content})
	c.NotifyObservers()
}

type User struct {
	Name             string
	ReceivedMessages []Message
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func (u *User) Update(messages []Message) {
	u.ReceivedMessages = messages
}

func (u *User) SendMessage(chatRoom Subject, content string) {
	chatRoom.SendMessage(u.Name, content)
}

func (u *User) ReadMessages() {
	fmt.Printf("Messages for %s:\n", u.Name)
	for _, msg := range u.ReceivedMessages {
		fmt.Printf("%s: %s\n", msg.Sender, msg.Content)
	}
}

func main() {
	chatRoom := NewChatRoom()

	admin := NewUser("Admin")
	alice := NewUser("Alice")
	bob := NewUser("Bob")
	carol := NewUser("Carol")

	chatRoom.RegisterObserver(admin)
	chatRoom.RegisterObserver(alice)
	chatRoom.RegisterObserver(bob)
	chatRoom.RegisterObserver(carol)

	alice.SendMessage(chatRoom, "Hello, everyone!")
	bob.SendMessage(chatRoom, "Hey, Alice!")
	carol.SendMessage(chatRoom, "Hi, all!")

	chatRoom.RemoveObserver(bob)
	alice.SendMessage(chatRoom, "Bob has left the chat.")
	admin.ReadMessages()
}
