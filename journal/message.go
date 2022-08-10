package journal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Message struct {
	Content string
}

func NewMessage(msg []string) *Message {
	content := strings.Join(msg, " ")
	return &Message{
		Content: content,
	}
}

func MessageToBytes(msg *Message) ([]byte, error) {
	return json.Marshal(msg)
}

func (j *Journal) Run(message []string) {

	defer j.DB.Close()
	//TODO: Function to take a plugin function

	b, err := MessageToBytes(NewMessage(message))
	if err != nil {
		log.Fatal(err)
	}

	key, err := j.StoreMessage(b)
	if err != nil {
		log.Fatalln(err)
	}

	// Default display operations
	err = j.UpdateHead(string(key))
	if err != nil {
		log.Fatalln(err)
	}

	err = j.Head()
	if err != nil {
		log.Fatalln(err)
	}

}
func (j *Journal) StoreMessage(msg []byte) ([]byte, error) {
	key := []byte(fmt.Sprintf("%v", time.Now().Unix()))
	err := j.DB.Put(key, msg)
	if err != nil {
		log.Fatal(err)
	}
	return key, nil
}
