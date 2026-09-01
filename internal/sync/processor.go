package sync

import "context"

type EventProcessor interface {
	Process(ctx context.Context, rawPayload []byte) error
}
