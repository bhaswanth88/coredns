package dnstap

import (
	"context"
	"time"

	"github.com/bhaswanth88/coredns/plugin/dnstap/msg"
	"github.com/bhaswanth88/coredns/request"

	tap "github.com/dnstap/golang-dnstap"
	"github.com/miekg/dns"
)

// ResponseWriter captures the client response and logs the query to dnstap.
type ResponseWriter struct {
	queryTime time.Time
	query     *dns.Msg
	ctx       context.Context
	dns.ResponseWriter
	Dnstap
}

// WriteMsg writes back the response to the client and THEN works on logging the request and response to dnstap.
func (w *ResponseWriter) WriteMsg(resp *dns.Msg) error {
	err := w.ResponseWriter.WriteMsg(resp)
	if err != nil {
		return err
	}

	r := new(tap.Message)
	msg.SetQueryTime(r, w.queryTime)
	msg.SetResponseTime(r, time.Now())
	msg.SetQueryAddress(r, w.RemoteAddr())

	if w.IncludeRawMessage {
		buf, _ := resp.Pack()
		r.ResponseMessage = buf
	}

	msg.SetType(r, tap.Message_CLIENT_RESPONSE)
	state := request.Request{W: w.ResponseWriter, Req: w.query}
	w.TapMessageWithMetadata(w.ctx, r, state)
	return nil
}
