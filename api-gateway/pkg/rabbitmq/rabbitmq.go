package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn *amqp.Connection
	ch *amqp.Channel

	url string
}

func New(url string) (*Rabbit, error) {
	rabbit := &Rabbit{
		url: url,
	}

	err := rabbit.connect()
	if err != nil {
		return nil, fmt.Errorf("connect: %w",err)
	}

	return rabbit,nil
}

func (r *Rabbit) connect () error{
	r.Close()

	var err error
	r.conn, err = amqp.Dial(r.url)
	if err != nil {
		return fmt.Errorf("amqp.Dial: %w",err)
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w",err)
	}
	return nil
}

func (r *Rabbit) Close() {
	if !r.conn.IsClosed() {
		r.conn.Close()
	}
	
	if !r.ch.IsClosed(){
		r.ch.Close()
	}
}