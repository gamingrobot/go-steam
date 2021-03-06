// Code generated by protoc-gen-go.
// source: steammessages_player.steamclient.proto
// DO NOT EDIT!

package unified

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// discarding unused import google_protobuf "code.google.com/p/goprotobuf/protoc-gen-go/descriptor"
// discarding unused import steammessages_unified_base_steamclient "steammessages_unified_base.steamclient.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CPlayer_GetGameBadgeLevels_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Request) Reset()         { *m = CPlayer_GetGameBadgeLevels_Request{} }
func (m *CPlayer_GetGameBadgeLevels_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Request) ProtoMessage()    {}

func (m *CPlayer_GetGameBadgeLevels_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CPlayer_GetGameBadgeLevels_Response struct {
	PlayerLevel      *uint32                                      `protobuf:"varint,1,opt,name=player_level" json:"player_level,omitempty"`
	Badges           []*CPlayer_GetGameBadgeLevels_Response_Badge `protobuf:"bytes,2,rep,name=badges" json:"badges,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response) Reset()         { *m = CPlayer_GetGameBadgeLevels_Response{} }
func (m *CPlayer_GetGameBadgeLevels_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response) ProtoMessage()    {}

func (m *CPlayer_GetGameBadgeLevels_Response) GetPlayerLevel() uint32 {
	if m != nil && m.PlayerLevel != nil {
		return *m.PlayerLevel
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response) GetBadges() []*CPlayer_GetGameBadgeLevels_Response_Badge {
	if m != nil {
		return m.Badges
	}
	return nil
}

type CPlayer_GetGameBadgeLevels_Response_Badge struct {
	Level            *int32  `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Series           *int32  `protobuf:"varint,2,opt,name=series" json:"series,omitempty"`
	BorderColor      *uint32 `protobuf:"varint,3,opt,name=border_color" json:"border_color,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) Reset() {
	*m = CPlayer_GetGameBadgeLevels_Response_Badge{}
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response_Badge) ProtoMessage()    {}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetSeries() int32 {
	if m != nil && m.Series != nil {
		return *m.Series
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetBorderColor() uint32 {
	if m != nil && m.BorderColor != nil {
		return *m.BorderColor
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Request struct {
	MinLastPlayed    *uint32 `protobuf:"varint,1,opt,name=min_last_played" json:"min_last_played,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Request) Reset()         { *m = CPlayer_GetLastPlayedTimes_Request{} }
func (m *CPlayer_GetLastPlayedTimes_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Request) ProtoMessage()    {}

func (m *CPlayer_GetLastPlayedTimes_Request) GetMinLastPlayed() uint32 {
	if m != nil && m.MinLastPlayed != nil {
		return *m.MinLastPlayed
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Response struct {
	Games            []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response) Reset()         { *m = CPlayer_GetLastPlayedTimes_Response{} }
func (m *CPlayer_GetLastPlayedTimes_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response) ProtoMessage()    {}

func (m *CPlayer_GetLastPlayedTimes_Response) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayer_GetLastPlayedTimes_Response_Game struct {
	Appid            *int32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	LastPlaytime     *uint32 `protobuf:"varint,2,opt,name=last_playtime" json:"last_playtime,omitempty"`
	Playtime_2Weeks  *int32  `protobuf:"varint,3,opt,name=playtime_2weeks" json:"playtime_2weeks,omitempty"`
	PlaytimeForever  *int32  `protobuf:"varint,4,opt,name=playtime_forever" json:"playtime_forever,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) Reset() {
	*m = CPlayer_GetLastPlayedTimes_Response_Game{}
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response_Game) ProtoMessage()    {}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetAppid() int32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetLastPlaytime() uint32 {
	if m != nil && m.LastPlaytime != nil {
		return *m.LastPlaytime
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetPlaytime_2Weeks() int32 {
	if m != nil && m.Playtime_2Weeks != nil {
		return *m.Playtime_2Weeks
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetPlaytimeForever() int32 {
	if m != nil && m.PlaytimeForever != nil {
		return *m.PlaytimeForever
	}
	return 0
}

type CPlayer_AcceptSSA_Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayer_AcceptSSA_Request) Reset()         { *m = CPlayer_AcceptSSA_Request{} }
func (m *CPlayer_AcceptSSA_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Request) ProtoMessage()    {}

type CPlayer_AcceptSSA_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayer_AcceptSSA_Response) Reset()         { *m = CPlayer_AcceptSSA_Response{} }
func (m *CPlayer_AcceptSSA_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Response) ProtoMessage()    {}

type CPlayer_LastPlayedTimes_Notification struct {
	Games            []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *CPlayer_LastPlayedTimes_Notification) Reset()         { *m = CPlayer_LastPlayedTimes_Notification{} }
func (m *CPlayer_LastPlayedTimes_Notification) String() string { return proto.CompactTextString(m) }
func (*CPlayer_LastPlayedTimes_Notification) ProtoMessage()    {}

func (m *CPlayer_LastPlayedTimes_Notification) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

func init() {
}
