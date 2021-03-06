package nsqkit

import (
	"fmt"
	"testing"
	"time"

	"github.com/klbud/nsq-std-foundation/nsq/message"
	"github.com/klbud/nsq-std-foundation/nsq/mock"
)

func TestPublisherServer_Publish(t *testing.T) {
	np := NewPublisher()
	quit := false
	go func() {
		time.Sleep(2 * time.Minute)
		np.StopPublisher()
		quit = true
	}()

	for {
		np.Publish(&message.Message{
			TopicName: mock.TopicName,
			Content:   []byte("hello"),
		})

		time.Sleep(1 * time.Millisecond)

		if quit {
			break
		}
	}

	t.Log("test publisher exit")
}

func TestPublisherServer_MultiPublish(t *testing.T) {
	muti := [][]byte{[]byte("{\"json\"}"), []byte("aaa")}
	fmt.Println(muti)
}
