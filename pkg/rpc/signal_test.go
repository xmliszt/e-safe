package rpc

import "testing"

// Expected to send to given PID with given data
func TestSendSignal(t *testing.T) {
	node := &Node{
		Pid: 999,
		RpcMap: make(map[int]chan *Data),
	}
	channel := make(chan *Data)
	node.RpcMap[555] = channel
	go func(){
		node.SendSignal(555, &Data{
			From: node.Pid,
			To: 555,
			Payload: map[string]interface{}{
				"type": "test",
			},
		})
	}()
	msg := <- channel
	if msg.From != 999 || msg.To != 555 || msg.Payload["type"] != "test" {
		t.Errorf("SendSignal has sending error! Received %v instead!", msg)
	}
}