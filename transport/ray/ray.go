package ray

import (
	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
)

// OutboundRay is a transport interface for outbound connections.
type OutboundRay interface {
	// OutboundInput provides a stream for the input of the outbound connection.
	// The outbound connection shall write all the input until it is closed.
	OutboundInput() InputStream

	// OutboundOutput provides a stream to retrieve the response from the
	// outbound connection. The outbound connection shall close the channel
	// after all responses are received and put into the channel.
	OutboundOutput() OutputStream
}

// InboundRay is a transport interface for inbound connections.
type InboundRay interface {
	// InboundInput provides a stream to retrieve the request from client.
	// The inbound connection shall close the channel after the entire request
	// is received and put into the channel.
	InboundInput() OutputStream

	// InboundOutput provides a stream of data for the inbound connection to write
	// as response. The inbound connection shall write all the data from the
	// channel until it is closed.
	InboundOutput() InputStream
}

// Ray is an internal transport channel between inbound and outbound connection.
type Ray interface {
	InboundRay
	OutboundRay
}

type RayStream interface {
	common.Closable
	CloseError()
}

type InputStream interface {
	buf.Reader
	buf.TimeoutReader
	RayStream
	Peek(*buf.Buffer)
}

type OutputStream interface {
	buf.Writer
	RayStream
}
