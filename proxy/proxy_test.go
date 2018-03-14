package proxy

import "testing"

func TestProxy(t *testing.T) {

	sub := &RealSubject{}
	proxy := NewProxy(sub)

	res := proxy.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}