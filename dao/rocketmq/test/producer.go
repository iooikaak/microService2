package test

import (
	"context"
	"fmt"
)

const (
	ProducerTopicTest = "INFRA-TOPIC-DEV-GOTEST"
)

func (d *Dao) SendToTest(ctx context.Context, msg, tag, shardingKey string) error {
	r, err := d.Producer.Send(ctx, ProducerTopicTest, []byte(msg), 2, tag, shardingKey)
	if err != nil {
		fmt.Printf("Producer Send r: %#v err: %v", r, err)
		return err
	}
	return nil
}
