package gogram

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// HandlerFunc represents a handler function, which is
// used to handle actual endpoints.
type HandlerFunc func(Context) error

// NewContext returns a new native context object,
// field by the passed update.
func NewContext(b API, u Update) Context {
	return &nativeContext{
		b: b,
		u: u,
	}
}

// Context wraps an update and represents the context of current event.
type Context interface {
	// Bot returns the bot instance.
	Bot() API

	// Update returns the original update.
	Update() Update

	// Message returns stored message if such presented.
	Message() *Message

	// Callback returns stored callback if such presented.
	Callback() *Callback

	// Query returns stored query if such presented.
	Query() *Query

	// InlineResult returns stored inline result if such presented.
	InlineResult() *InlineResult

	// ShippingQuery returns stored shipping query if such presented.
	ShippingQuery() *ShippingQuery

	// PreCheckoutQuery returns stored pre checkout query if such presented.
	PreCheckoutQuery() *PreCheckoutQuery

	// Payment returns payment instance.
	Payment() *Payment

	// Poll returns stored poll if such presented.
	Poll() *Poll

	// PollAnswer returns stored poll answer if such presented.
	PollAnswer() *PollAnswer

	// ChatMember returns chat member changes.
	ChatMember() *ChatMemberUpdate

	// ChatJoinRequest returns the chat join request.
	ChatJoinRequest() *ChatJoinRequest

	// Migration returns both migration from and to chat IDs.
	Migration() (int64, int64)

	// Topic returns the topic changes.
	Topic() *Topic

	// Boost returns the boost instance.
	Boost() *BoostUpdated

	// BoostRemoved returns the boost removed from a chat instance.
	BoostRemoved() *BoostRemoved

	// Sender returns the current recipient, depending on the context type.
	// Returns nil if user is not presented.
	Sender() *User

	// Chat returns the current chat, depending on the context type.
	// Returns nil if chat is not presented.
	Chat() *Chat
	// Recipient combines both Sender and Chat functions. If there is no user
	// the chat will be returned. The native context cannot be without sender,
	// but it is useful in the case when the context created intentionally
	// by the NewContext constructor and have only Chat field inside.
	Recipient() Recipient

	// Text returns the message text, depending on the context type.
	// In the case when no related data presented, returns an empty string.
	Text() string

	// Entities returns the message entities, whether it's media caption's or the text's.
	// In the case when no entities presented, returns a nil.
	Entities() Entities

	// Data returns the current data, depending on the context type.
	// If the context contains command, returns its arguments string.
	// If the context contains payment, returns its payload.
	// In the case when no related data presented, returns an empty string.
	Data() string

	// Args returns the raw arguments from a command or callback as a slice of strings.
	// Command arguments are split by spaces, and callback arguments are split by the "?" symbol.
	//
	// Example:
	// Command: "/command arg1 arg2" -> Args() returns ["arg1", "arg2"]
	// Callback: "/callback?param=value" -> Args() returns ["param=value"].
	Args() []string

	// ParseArgs parses command or callback arguments and assigns them to the provided variables.
	// It expects a variable number of arguments, each being a pointer to a supported type (e.g., *int, *bool, *string).
	//
	// The function converts the raw string arguments into the specified types.
	// Supported types include:
	// - bool
	// - int, int8, int16, int32, int64
	// - uint, uint8, uint16, uint32, uint64
	// - float32, float64
	// - string
	//
	// If the number of arguments does not match the number of provided variables,
	// or if any conversion fails, an error is returned with details about the issue.
	//
	// Example usage:
	// var name string
	// var age int
	// var active bool
	// err := c.ParseArgs(&name, &age, &active)
	// If successful, the variables will be populated with the parsed values.
	ParseArgs(vars ...interface{}) error

	// Send sends a message to the current recipient.
	// See Send from bot.go.
	Send(what interface{}, opts ...interface{}) error

	// SendAlbum sends an album to the current recipient.
	// See SendAlbum from bot.go.
	SendAlbum(a Album, opts ...interface{}) error

	// Reply replies to the current message.
	// See Reply from bot.go.
	Reply(what interface{}, opts ...interface{}) error

	// Forward forwards the given message to the current recipient.
	// See Forward from bot.go.
	Forward(msg Editable, opts ...interface{}) error

	// ForwardTo forwards the current message to the given recipient.
	// See Forward from bot.go
	ForwardTo(to Recipient, opts ...interface{}) error

	// Edit edits the current message.
	// See Edit from bot.go.
	Edit(what interface{}, opts ...interface{}) error

	// EditCaption edits the caption of the current message.
	// See EditCaption from bot.go.
	EditCaption(caption string, opts ...interface{}) error

	// EditOrSend edits the current message if the update is callback,
	// otherwise the content is sent to the chat as a separate message.
	EditOrSend(what interface{}, opts ...interface{}) error

	// EditOrReply edits the current message if the update is callback,
	// otherwise the content is replied as a separate message.
	EditOrReply(what interface{}, opts ...interface{}) error

	// EditAfter schedules an Edit operation to be performed after a specified duration (d).
	// It returns a *time.Timer that can be used to manage or cancel the scheduled action.
	EditAfter(d time.Duration, what interface{}, opts ...interface{}) *time.Timer

	// EditAfterSignal waits for a message from whatChan, then performs an Edit operation with the received data.
	// It blocks until the channel provides a value, and then proceeds to edit using that value.
	EditAfterSignal(whatChan chan interface{}, opts ...interface{})

	// EditAfterSignalWithTimeout waits for a message from whatChan or a timeout.
	// If a message is received from whatChan before the timeout, it performs an Edit with the received data.
	// If the timeout occurs first, it performs an Edit using the timeoutWhat data.
	EditAfterSignalWithTimeout(timeout time.Duration, whatChan chan interface{}, timeoutWhat interface{}, opts ...interface{})

	// Delete removes the current message.
	// See Delete from bot.go.
	Delete() error

	// DeleteAfter waits for the duration to elapse and then removes the
	// message. It handles an error automatically using b.OnError callback.
	// It returns a Timer that can be used to cancel the call using its Stop method.
	DeleteAfter(d time.Duration) *time.Timer

	// DeleteAfterSignal waits for a boolean signal from the specified channel (Chan).
	// When a signal is received, it performs a Delete operation.
	// If the Delete operation encounters an error, it invokes the bot's error handler (OnError) to handle the error gracefully.
	// This function blocks until a signal is received, making it suitable for use in situations where
	// a delete action should occur in response to an external trigger.
	DeleteAfterSignal(Chan chan bool)

	// Notify updates the chat action for the current recipient.
	// See Notify from bot.go.
	Notify(action ChatAction) error

	// Ship replies to the current shipping query.
	// See Ship from bot.go.
	Ship(what ...interface{}) error

	// Accept finalizes the current deal.
	// See Accept from bot.go.
	Accept(errorMessage ...string) error

	// Answer sends a response to the current inline query.
	// See Answer from bot.go.
	Answer(resp *QueryResponse) error

	// Respond sends a response for the current callback query.
	// See Respond from bot.go.
	Respond(resp ...*CallbackResponse) error

	// RespondText sends a popup response for the current callback query.
	RespondText(text string) error

	// RespondAlert sends an alert response for the current callback query.
	RespondAlert(text string) error

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})
}

// nativeContext is a native implementation of the Context interface.
// "context" is taken by context package, maybe there is a better name.
type nativeContext struct {
	b     API
	u     Update
	lock  sync.RWMutex
	store map[string]interface{}
}

func (c *nativeContext) Bot() API {
	return c.b
}

func (c *nativeContext) Update() Update {
	return c.u
}

func (c *nativeContext) Message() *Message {
	switch {
	case c.u.Message != nil:
		return c.u.Message
	case c.u.Callback != nil:
		return c.u.Callback.Message
	case c.u.EditedMessage != nil:
		return c.u.EditedMessage
	case c.u.ChannelPost != nil:
		if c.u.ChannelPost.PinnedMessage != nil {
			return c.u.ChannelPost.PinnedMessage
		}
		return c.u.ChannelPost
	case c.u.EditedChannelPost != nil:
		return c.u.EditedChannelPost
	default:
		return nil
	}
}

func (c *nativeContext) Callback() *Callback {
	return c.u.Callback
}

func (c *nativeContext) Query() *Query {
	return c.u.Query
}

func (c *nativeContext) InlineResult() *InlineResult {
	return c.u.InlineResult
}

func (c *nativeContext) ShippingQuery() *ShippingQuery {
	return c.u.ShippingQuery
}

func (c *nativeContext) PreCheckoutQuery() *PreCheckoutQuery {
	return c.u.PreCheckoutQuery
}

func (c *nativeContext) Payment() *Payment {
	if c.u.Message == nil {
		return nil
	}
	return c.u.Message.Payment
}

func (c *nativeContext) ChatMember() *ChatMemberUpdate {
	switch {
	case c.u.ChatMember != nil:
		return c.u.ChatMember
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember
	default:
		return nil
	}
}

func (c *nativeContext) ChatJoinRequest() *ChatJoinRequest {
	return c.u.ChatJoinRequest
}

func (c *nativeContext) Poll() *Poll {
	return c.u.Poll
}

func (c *nativeContext) PollAnswer() *PollAnswer {
	return c.u.PollAnswer
}

func (c *nativeContext) Migration() (int64, int64) {
	m := c.u.Message
	if m == nil {
		return 0, 0
	}
	return m.MigrateFrom, m.MigrateTo
}

func (c *nativeContext) Topic() *Topic {
	m := c.u.Message
	if m == nil {
		return nil
	}
	switch {
	case m.TopicCreated != nil:
		return m.TopicCreated
	case m.TopicReopened != nil:
		return m.TopicReopened
	case m.TopicEdited != nil:
		return m.TopicEdited
	}
	return nil
}

func (c *nativeContext) Boost() *BoostUpdated {
	return c.u.Boost
}

func (c *nativeContext) BoostRemoved() *BoostRemoved {
	return c.u.BoostRemoved
}

func (c *nativeContext) Sender() *User {
	switch {
	case c.u.Callback != nil:
		return c.u.Callback.Sender
	case c.Message() != nil:
		return c.Message().Sender
	case c.u.Query != nil:
		return c.u.Query.Sender
	case c.u.InlineResult != nil:
		return c.u.InlineResult.Sender
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Sender
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Sender
	case c.u.PollAnswer != nil:
		return c.u.PollAnswer.Sender
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.Sender
	case c.u.ChatMember != nil:
		return c.u.ChatMember.Sender
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Sender
	case c.u.Boost != nil:
		if b := c.u.Boost.Boost; b != nil && b.Source != nil {
			return b.Source.Booster
		}
	case c.u.BoostRemoved != nil:
		if b := c.u.BoostRemoved; b.Source != nil {
			return b.Source.Booster
		}
	}
	return nil
}

func (c *nativeContext) Chat() *Chat {
	switch {
	case c.Message() != nil:
		return c.Message().Chat
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.Chat
	case c.u.ChatMember != nil:
		return c.u.ChatMember.Chat
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Chat
	default:
		return nil
	}
}

func (c *nativeContext) Recipient() Recipient {
	chat := c.Chat()
	if chat != nil {
		return chat
	}
	return c.Sender()
}

func (c *nativeContext) Text() string {
	m := c.Message()
	if m == nil {
		return ""
	}
	if m.Caption != "" {
		return m.Caption
	}
	return m.Text
}

func (c *nativeContext) Entities() Entities {
	m := c.Message()
	if m == nil {
		return nil
	}
	if len(m.CaptionEntities) > 0 {
		return m.CaptionEntities
	}
	return m.Entities
}

func (c *nativeContext) Data() string {
	switch {
	case c.u.Message != nil:
		m := c.u.Message
		if m.Payment != nil {
			return m.Payment.Payload
		}
		return m.Payload
	case c.u.Callback != nil:
		return c.u.Callback.Data
	case c.u.Query != nil:
		return c.u.Query.Text
	case c.u.InlineResult != nil:
		return c.u.InlineResult.Query
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Payload
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Payload
	default:
		return ""
	}
}

func (c *nativeContext) Args() []string {
	m := c.u.Message
	switch {
	case m != nil && m.Payment != nil:
		return strings.Split(m.Payment.Payload, "?")
	case m != nil:
		payload := strings.Trim(m.Payload, " ")
		if payload != "" {
			return strings.Fields(payload)
		}
	case c.u.Callback != nil:
		return strings.Split(c.u.Callback.Data, "?")
	case c.u.Query != nil:
		return strings.Split(c.u.Query.Text, " ")
	case c.u.InlineResult != nil:
		return strings.Split(c.u.InlineResult.Query, " ")
	}
	return nil
}

func (c *nativeContext) ParseArgs(vars ...interface{}) error {
	args := c.Args()

	// Check if the number of args matches the number of variables
	if len(args) != len(vars) {
		return fmt.Errorf("number of arguments (%d) does not match the expected number of variables (%d)", len(args), len(vars))
	}

	// Iterate over the arguments and variables
	for i, v := range vars {
		switch variable := v.(type) {
		case *string:
			*variable = args[i]
		case *bool:
			boolVal, err := strconv.ParseBool(args[i])
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to bool: %v", i+1, args[i], err)
			}
			*variable = boolVal
		case *int:
			intVal, err := strconv.Atoi(args[i])
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to int: %v", i+1, args[i], err)
			}
			*variable = intVal
		case *int8:
			intVal, err := strconv.ParseInt(args[i], 10, 8)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to int8: %v", i+1, args[i], err)
			}
			*variable = int8(intVal)
		case *int16:
			intVal, err := strconv.ParseInt(args[i], 10, 16)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to int16: %v", i+1, args[i], err)
			}
			*variable = int16(intVal)
		case *int32:
			intVal, err := strconv.ParseInt(args[i], 10, 32)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to int32: %v", i+1, args[i], err)
			}
			*variable = int32(intVal)
		case *int64:
			intVal, err := strconv.ParseInt(args[i], 10, 64)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to int64: %v", i+1, args[i], err)
			}
			*variable = intVal
		case *uint:
			uintVal, err := strconv.ParseUint(args[i], 10, 64)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to uint: %v", i+1, args[i], err)
			}
			*variable = uint(uintVal)
		case *uint8:
			uintVal, err := strconv.ParseUint(args[i], 10, 8)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to uint8: %v", i+1, args[i], err)
			}
			*variable = uint8(uintVal)
		case *uint16:
			uintVal, err := strconv.ParseUint(args[i], 10, 16)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to uint16: %v", i+1, args[i], err)
			}
			*variable = uint16(uintVal)
		case *uint32:
			uintVal, err := strconv.ParseUint(args[i], 10, 32)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to uint32: %v", i+1, args[i], err)
			}
			*variable = uint32(uintVal)
		case *uint64:
			uintVal, err := strconv.ParseUint(args[i], 10, 64)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to uint64: %v", i+1, args[i], err)
			}
			*variable = uint64(uintVal)
		case *float32:
			floatVal, err := strconv.ParseFloat(args[i], 32)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to float32: %v", i+1, args[i], err)
			}
			*variable = float32(floatVal)
		case *float64:
			floatVal, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				return fmt.Errorf("argument %d ('%s') could not be converted to float64: %v", i+1, args[i], err)
			}
			*variable = floatVal
		default:
			return fmt.Errorf("unsupported variable type at argument %d", i+1)
		}
	}

	return nil
}

func (c *nativeContext) Send(what interface{}, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)
	_, err := c.b.Send(c.Recipient(), what, opts...)
	return err
}

func (c *nativeContext) inheritOpts(opts ...interface{}) []interface{} {
	var (
		ignoreThread bool
	)

	if opts == nil {
		opts = make([]interface{}, 0)
	}

	for _, opt := range opts {
		switch opt.(type) {
		case Option:
			switch opt {
			case IgnoreThread:
				ignoreThread = true
			default:
			}
		}
	}

	switch {
	case !ignoreThread && c.Message() != nil && c.Message().ThreadID != 0:
		opts = append(opts, &Topic{ThreadID: c.Message().ThreadID})
	}

	return opts
}

func (c *nativeContext) SendAlbum(a Album, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)

	_, err := c.b.SendAlbum(c.Recipient(), a, opts...)
	return err
}

func (c *nativeContext) Reply(what interface{}, opts ...interface{}) error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	opts = c.inheritOpts(opts...)
	_, err := c.b.Reply(msg, what, opts...)
	return err
}

func (c *nativeContext) Forward(msg Editable, opts ...interface{}) error {
	_, err := c.b.Forward(c.Recipient(), msg, opts...)
	return err
}

func (c *nativeContext) ForwardTo(to Recipient, opts ...interface{}) error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	_, err := c.b.Forward(to, msg, opts...)
	return err
}

func (c *nativeContext) Edit(what interface{}, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)

	if c.u.InlineResult != nil {
		_, err := c.b.Edit(c.u.InlineResult, what, opts...)
		return err
	}
	if c.u.Callback != nil {
		_, err := c.b.Edit(c.u.Callback, what, opts...)
		return err
	}
	return ErrBadContext
}

func (c *nativeContext) EditCaption(caption string, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)

	if c.u.InlineResult != nil {
		_, err := c.b.EditCaption(c.u.InlineResult, caption, opts...)
		return err
	}
	if c.u.Callback != nil {
		_, err := c.b.EditCaption(c.u.Callback, caption, opts...)
		return err
	}
	return ErrBadContext
}

func (c *nativeContext) EditOrSend(what interface{}, opts ...interface{}) error {
	err := c.Edit(what, opts...)
	if errors.Is(err, ErrBadContext) {
		return c.Send(what, opts...)
	}
	return err
}

func (c *nativeContext) EditOrReply(what interface{}, opts ...interface{}) error {
	err := c.Edit(what, opts...)
	if errors.Is(err, ErrBadContext) {
		return c.Reply(what, opts...)
	}
	return err
}

func (c *nativeContext) EditAfter(d time.Duration, what interface{}, opts ...interface{}) *time.Timer {
	return time.AfterFunc(d, func() {
		if err := c.Edit(what, opts); err != nil {
			if b, ok := c.b.(*Bot); ok {
				b.OnError(err, c)
			}
		}
	})
}

func (c *nativeContext) EditAfterSignal(whatChan chan interface{}, opts ...interface{}) {
	// Wait for the channel to send data
	what := <-whatChan

	// Perform the Edit operation with the received data
	if err := c.Edit(what, opts...); err != nil {
		if b, ok := c.b.(*Bot); ok {
			b.OnError(err, c)
		}
	}
}

func (c *nativeContext) EditAfterSignalWithTimeout(timeout time.Duration, whatChan chan interface{}, timeoutWhat interface{}, opts ...interface{}) {
	select {
	case what := <-whatChan:
		// Perform the Edit operation with the received message
		if err := c.Edit(what, opts...); err != nil {
			if b, ok := c.b.(*Bot); ok {
				b.OnError(err, c)
			}
		}

	case <-time.After(timeout):
		// Perform the Edit operation with the timeout message or data
		if err := c.Edit(timeoutWhat, opts...); err != nil {
			if b, ok := c.b.(*Bot); ok {
				b.OnError(err, c)
			}
		}
	}
}

func (c *nativeContext) Delete() error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	return c.b.Delete(msg)
}

func (c *nativeContext) DeleteAfter(d time.Duration) *time.Timer {
	return time.AfterFunc(d, func() {
		if err := c.Delete(); err != nil {
			if b, ok := c.b.(*Bot); ok {
				b.OnError(err, c)
			}
		}
	})
}

func (c *nativeContext) DeleteAfterSignal(Chan chan bool) {
	select {
	case <-Chan:
		// Perform the Delete operation when the signal is received
		if err := c.Delete(); err != nil {
			// Handle any error that occurs during the delete operation
			if b, ok := c.b.(*Bot); ok {
				b.OnError(err, c)
			}
		}
	}
}

func (c *nativeContext) Notify(action ChatAction) error {
	return c.b.Notify(c.Recipient(), action)
}

func (c *nativeContext) Ship(what ...interface{}) error {
	if c.u.ShippingQuery == nil {
		return errors.New("gogram: context shipping query is nil")
	}
	return c.b.Ship(c.u.ShippingQuery, what...)
}

func (c *nativeContext) Accept(errorMessage ...string) error {
	if c.u.PreCheckoutQuery == nil {
		return errors.New("gogram: context pre checkout query is nil")
	}
	return c.b.Accept(c.u.PreCheckoutQuery, errorMessage...)
}

func (c *nativeContext) Respond(resp ...*CallbackResponse) error {
	if c.u.Callback == nil {
		return errors.New("gogram: context callback is nil")
	}
	return c.b.Respond(c.u.Callback, resp...)
}

func (c *nativeContext) RespondText(text string) error {
	return c.Respond(&CallbackResponse{Text: text})
}

func (c *nativeContext) RespondAlert(text string) error {
	return c.Respond(&CallbackResponse{Text: text, ShowAlert: true})
}

func (c *nativeContext) Answer(resp *QueryResponse) error {
	if c.u.Query == nil {
		return errors.New("gogram: context inline query is nil")
	}
	return c.b.Answer(c.u.Query, resp)
}

func (c *nativeContext) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.store == nil {
		c.store = make(map[string]interface{})
	}

	c.store[key] = value
}

func (c *nativeContext) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.store[key]
}
