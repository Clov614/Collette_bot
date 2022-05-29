package eventListener

type MsgGroupEvent struct {
	Anonymous   interface{} `json:"anonymous"`
	Font        int         `json:"font"`
	GroupID     int         `json:"group_id"`
	Message     string      `json:"message"`
	MessageID   int         `json:"message_id"`
	MessageSeq  int         `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfID      int         `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
	SubType string `json:"sub_type"`
	Time    int    `json:"time"`
	UserID  int    `json:"user_id"`
}

type MsgPrivateEvent struct {
	Font        int    `json:"font"`
	Message     string `json:"message"`
	MessageID   int    `json:"message_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	RawMessage  string `json:"raw_message"`
	SelfID      int    `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
	SubType  string `json:"sub_type"`
	TargetID int    `json:"target_id"`
	Time     int    `json:"time"`
	UserID   int    `json:"user_id"`
}

type LoginEvent struct {
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfID        int    `json:"self_id"`
	SubType       string `json:"sub_type"`
	Time          int    `json:"time"`
}

type HeartBeat struct {
	Interval      int    `json:"interval"`
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfID        int    `json:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"PacketReceived"`
			PacketSent      int `json:"PacketSent"`
			PacketLost      int `json:"PacketLost"`
			MessageReceived int `json:"MessageReceived"`
			MessageSent     int `json:"MessageSent"`
			LastMessageTime int `json:"LastMessageTime"`
			DisconnectTimes int `json:"DisconnectTimes"`
			LostTimes       int `json:"LostTimes"`
		} `json:"stat"`
	} `json:"status"`
	Time int `json:"time"`
}