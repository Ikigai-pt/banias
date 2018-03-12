package types
import (
	"strconv"
	"strings"
)



//{"sender_id",:"id", "events": [ e1, e2, ... ] }
//
// where e1 is an object with all the information you need.
// {
//"type": { "event_version": ..., "event_name": ...},
//"payload": { the event payload itself goes here in whatever structure you need },
//}

type TrackRequest struct {
	SenderID string  `json:"sender_id" valid:"notempty,required"`
	Events   []Event `json:"events"`
}

type Event struct {
	TypeField    Type   `json:"type"`
	PayloadField Payload `json:"payload" valid:"-"`
}



// Type is an Event's metadata.
type Type struct {
	EventVersionField string `json:"event_version" valid:"notempty,required"`
	EventNameField    string `json:"event_name" valid:"notempty,required"`
}


type EventMsg struct  {
	SenderID string
	Event    Event
}
// Payload is the event's actual data inserted into data stores.
type Payload map[string]interface{}

// Error describes either an API error, or a rejected event.
// If rejected event, then Index field will be included.
type Error struct {
	Detail string `json:"detail"`
	// This is a pointer in order to marshal '0' value
	// but not if uninitialized.
	Index *uint64 `json:"index,omitempty"`
}
// NewError returns an initialized Error.
func NewError(index uint64, detail string) *Error {
	return &Error{Index: &index, Detail: detail}
}

func (err *Error) Error() string {
	return strings.Join(
		[]string{"item ", strconv.FormatUint(*err.Index, 10), ": ", err.Detail},
		"")
}
//Example
/*
	logger = log.With(logger, "caller", log.Caller(4))

{
	"sender_id": "my id",
	"types": [{
			"type": {
				"event_version": "1",
				"event_name": "transaction"
			},
			"payload": {
				"action": "buy",
				"price": 170,
				"date": "03/31/1967"
			}
		},
		{
			"type": {
				"event_version": "1",
				"event_name": "transaction"

			},
			"payload": {
				"action": "sell",
				"price": 170,
				"date": "03/31/1967"

			}
		},
		{
			"type": {
				"event_version": "2",
				"event_name": "click"

			},
			"payload": {
				"screen": "welcome"

			}
		}
	]
}

 */