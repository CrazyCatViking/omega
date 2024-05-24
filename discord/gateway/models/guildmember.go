package gatewaymodels

type GuildMember struct {
  User User `json:"user"`
  Nick string `json:"nick"`
  Avatar string `json:"avatar"`
  Roles	[]string `json:"roles"`
  JoinedAt string `json:"joined_at"`
  PremiumSince string `json:"premium_since"`
  Deaf bool `json:"deaf"`
  Mute bool `json:"mute"`
  Flags	int `json:"flags"`
  Pending bool `json:"pending"`
  Permissions string `json:"permissions"`
  CommunicationDisabledUntil string `json:"communication_disabled_until"`
  // AvatarDecoration_data string `json:"avatar_decoration_data"`
}
