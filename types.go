package telebot

import "strconv"

// Sendable is any object that can send itself.
type Sendable interface {
	Send(*Bot, Recipient, *SendOptions) (*Message, error)
}

// Recipient is basically any possible endpoint you can send
// messages to. It's usually a distinct user or a chat.
type Recipient interface {
	// ID of user or group chat, @Username for channel
	Destination() string
}

// User object represents a Telegram user, bot
//
// object represents a group chat if Title is empty.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`
	Username string `json:"username"`
}

// Destination is internal user ID.
func (u *User) Destination() string {
	return strconv.Itoa(u.ID)
}

// Chat object represents a Telegram user, bot or group chat.
//
// Type of chat, can be either “private”, “group”, "supergroup" or “channel”
type Chat struct {
	ID int64 `json:"id"`

	// See telebot.ChatType and consts.
	Type ChatType `json:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Destination is internal chat ID.
func (c *Chat) Destination() string {
	ret := "@" + c.Username
	if c.Type != "channel" {
		ret = strconv.FormatInt(c.ID, 10)
	}
	return ret
}

// IsGroupChat returns true if chat object represents a group chat.
func (c *Chat) IsGroupChat() bool {
	return c.Type != "private"
}

// Update object represents an incoming update.
type Update struct {
	ID      int64    `json:"update_id"`
	Payload *Message `json:"message"`

	// optional
	Callback *Callback `json:"callback_query"`
	Query    *Query    `json:"inline_query"`
}

// Photo object represents a photo (with or without caption).
type Photo struct {
	File

	Width  int `json:"width"`
	Height int `json:"height"`

	Caption string `json:"caption,omitempty"`
}

// Audio object represents an audio file.
type Audio struct {
	File

	// Duration of the recording in seconds as defined by sender.
	Duration int `json:"duration"`

	// Title (optional) as defined by sender or by audio tags.
	Title string `json:"title"`

	// Performer (optional) is defined by sender or by audio tags.
	Performer string `json:"performer"`

	// MIME type (optional) of the file as defined by sender.
	Mime string `json:"mime_type"`

	Caption string `json:"caption,omitempty"`
}

// Voice object represents a voice note.
type Voice struct {
	File

	// Duration of the recording in seconds as defined by sender.
	Duration int `json:"duration"`

	// MIME type (optional) of the file as defined by sender.
	Mime string `json:"mime_type"`

	Caption string `json:"caption,omitempty"`
}

// Document object represents a general file (as opposed to Photo or Audio).
// Telegram users can send files of any type of up to 1.5 GB in size.
type Document struct {
	File

	// Document thumbnail as defined by sender.
	Preview Photo `json:"thumb"`

	// Original filename as defined by sender.
	FileName string `json:"file_name"`

	// MIME type of the file as defined by sender.
	Mime string `json:"mime_type"`

	Caption string `json:"caption,omitempty"`
}

// Sticker object represents a WebP image, so-called sticker.
type Sticker struct {
	File

	Width  int `json:"width"`
	Height int `json:"height"`

	// Sticker thumbnail in .webp or .jpg format.
	Thumbnail Photo `json:"thumb"`

	// Associated emoji
	Emoji string `json:"emoji"`
}

// Video object represents an MP4-encoded video.
type Video struct {
	Audio

	Width  int `json:"width"`
	Height int `json:"height"`

	// Text description of the video as defined by sender.
	Caption string `json:"caption,omitempty"`

	// Video thumbnail.
	Thumbnail Photo `json:"thumb"`
}

// This object represents a video message (available in Telegram apps
// as of v.4.0).
type VideoNote struct {
	File

	// Duration of the recording in seconds as defined by sender.
	Duration int `json:"duration"`

	// Video note thumbnail.
	Thumbnail Photo `json:"thumb"`
}

// Contact object represents a contact to Telegram user
type Contact struct {
	UserID      int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// Location object represents geographic position.
type Location struct {
	Lat float32 `json:"latitude"`
	Lng float32 `json:"longitude"`
}

// Venue object represents a venue location with name, address and
// optional foursquare ID.
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id,omitempty"`
}

// KeyboardButton represents a button displayed in reply-keyboard.
type KeyboardButton struct {
	Text string `json:"text"`

	Contact  bool `json:"request_contact,omitempty"`
	Location bool `json:"request_location,omitempty"`
}

// InlineButton represents a button displayed in the message.
type InlineButton struct {
	Text string `json:"text"`

	URL         string `json:"url,omitempty"`
	Data        string `json:"callback_data,omitempty"`
	InlineQuery string `json:"switch_inline_query,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears
// right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by
	// an Array of KeyboardButton objects.
	InlineKeyboard [][]InlineButton `json:"inline_keyboard,omitempty"`
}

// Callback object represents a query from a callback button in an
// inline keyboard.
type Callback struct {
	ID string `json:"id"`

	// For message sent to channels, Sender may be empty
	Sender *User `json:"from"`

	// Message will be set if the button that originated the query
	// was attached to a message sent by a bot.
	Message *Message `json:"message"`

	// MessageID will be set if the button was attached to a message
	// sent via the bot in inline mode.
	MessageID string `json:"inline_message_id"`

	// Data associated with the callback button. Be aware that
	// a bad client can send arbitrary data in this field.
	Data string `json:"data"`
}

// CallbackResponse builds a response to a Callback query.
//
// See also: https://core.telegram.org/bots/api#answerCallbackQuery
type CallbackResponse struct {
	// The ID of the callback to which this is a response.
	// It is not necessary to specify this field manually.
	CallbackID string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be shown to the user.
	Text string `json:"text,omitempty"`

	// (Optional) If true, an alert will be shown by the client instead
	// of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// (Optional) URL that will be opened by the user's client.
	// If you have created a Game and accepted the conditions via @Botfather
	// specify the URL that opens your game
	// note that this will only work if the query comes from a callback_game button.
	// Otherwise, you may use links like telegram.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url,omitempty"`
}

// ChatMember object represents information about a single chat member.
type ChatMember struct {
	User   User   `json:"user"`
	Status string `json:"status"`
}

// UserProfilePhotos object represent a user's profile pictures.
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has.
	Count int `json:"total_count"`

	// Requested profile pictures (in up to 4 sizes each).
	Photos [][]Photo `json:"photos"`
}
