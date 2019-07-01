package tencentim

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	SetAppID("142742")
	sig := ""
	sendMsg := NewSendMsg(sig, "USER:", "xxx")
	send, err := Send(sendMsg)
	if err != nil {
		t.Fatal("test error", err)
	}
	t.Log(string(send))
}

func TestSendWithOfflinePushInfo(t *testing.T) {
	SetAppID("142742")
	sig := ""
	sendMsg := NewSendMsg(sig, "USER:", "xxx")
	sendMsg.SendMsgBody.SetTitle("test")
	send, err := Send(sendMsg)
	if err != nil {
		t.Fatal("test error", err)
	}
	t.Log(string(send))
}

func TestSendWithAndroidApns(t *testing.T) {
	SetAppID("142742")
	sig := ""
	sendMsg := NewSendMsg(sig, "USER:", "xxx")
	sendMsg.SendMsgBody.SetSound("www.baidu.mp3")
	send, err := Send(sendMsg)
	if err != nil {
		t.Fatal("test error", err)
	}
	t.Log(string(send))
}

func BenchmarkTestSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetAppID("1400202")
		sig := ""
		sendMsg := NewSendMsg(sig, "", "xxx")
		send, err := Send(sendMsg)
		if err != nil {
			b.Fatal("test error", err)
		}
		b.Log(string(send))
	}
}

func TestBatchSend(t *testing.T) {
	SetAppID("140025")
	sig := "eJxNjV1"
	var t1 []string
	for i := 1; i <= 501; i++ {
		t1 = append(t1, fmt.Sprintf("ROBOT:%d", i))
	}
	sendMsg := NewBatchSendMsg(sig, t1, "xxx")
	send, err := Send(sendMsg)
	if err != nil {
		t.Fatal("test error", err)
	}
	t.Log(string(send))
}
