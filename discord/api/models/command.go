package apimodels

var (
  SLASH_COMMAND = 1
  USER_COMMAND = 2
  MESSAGE_COMMAND = 3
)

var (
  GUILD_INSTALL = 0
  USER_INSTALL = 1
)

var (
  GUILD = 0
  BOT_DM = 1
  PRIVATE_CHANNEL = 2
)

var (
  Indonesian = "id"
  Danish = "da"
  German = "de"
	EnglishUK = "en-GB"
	EnglishUS = "en-US"
	SpanishES	 = "es-ES"
	Spanish419 = "es-419"
  French = "fr"
  Croatian = "hr"
  Italian = "it"
  Lithuanian = "lt"
  Hungarian = "hu"
  Dutch = "nl"
  Norwegian = "no"
  Polish = "pl"
	Portuguese = "pt-BR"
  Romanian = "ro"
  Finnish = "fi"
	Swedish	 = "sv-SE"
  Vietnamese = "vi"
  Turkish = "tr"
  Czech = "cs"
  Greek = "el"
  Bulgarian = "bg"
  Russian = "ru"
  Ukrainian = "uk"
  Hindi = "hi"
  Thai = "th"
	Chinese = "zh-CN"
  Japanese = "ja"
	ChineseTaiwan = "zh-TW"
  Korean = "ko"
)

var (
  SUB_COMMAND	= 1	
  SUB_COMMAND_GROUP	= 2	
  STRING	= 3	
  INTEGER	= 4
  BOOLEAN	= 5	
  USER = 6	
  CHANNEL	= 7
  ROLE	= 8	
  MENTIONABLE	= 9
  NUMBER	= 10
  ATTACHMENT	= 11
)

var (
  GUILD_TEXT = 0
  DM = 1
  GUILD_VOICE = 2
  GROUP_DM = 3
  GUILD_CATEGORY = 4
  GUILD_ANNOUNCEMENT = 5
  ANNOUNCEMENT_THREAD = 1
  PUBLIC_THREAD = 1
  PRIVATE_THREAD = 1
  GUILD_STAGE_VOICE = 1
  GUILD_DIRECTORY = 1
  GUILD_FORUM = 1
  GUILD_MEDIA = 1
)

type CreateCommandInput struct {
  Id string `json:"id"`
  Type int `json:"type"`
  ApplicationId string `json:"application_id"`
  GuildId string `json:"guild_id"`
  Name string `json:"name"`
  NameLocalizations map[string]string `json:"name_localizations"`
  Description string `json:"description"`
  DescriptionLocalizations map[string]string `json:"description_localizations"`
  Options CommandOption `json:"options"`
  DefaultMemberPermissions string `json:"default_member_permissions"`
  DmPermission bool `json:"dm_permission"`
  DefaultPermission bool `json:"default_permission"`
  Nsfw bool `json:"nsfw"`
  IntegrationTypes []int `json:"integration_types"`
  Contexts []int `json:"contexts"`
  Version string `json:"version"`
}

type CommandOption struct {
  Type int `json:"type"`
  Name string `json:"name"`
  NameLocalizations map[string]string `json:"name_localizations"`
  Description string `json:"description"`
  DescriptionLocalizations map[string]string `json:"description_localizations"`
  Required bool `json:"required"`
  Choices []CommandOptionChoice `json:"choices"`
  Options []CommandOption `json:"options"`
  ChannelTypes []int `json:"channel_types"`
  MinValue string `json:"min_value"`
  MaxValue string `json:"max_value"`
  MinLength string `json:"min_length"`
  MaxLength string `json:"max_length"`
  Autocomplete bool `json:"autocomplete"`
}

type CommandOptionChoice struct {
  Name string `json:"name"`
  NameLocalizations map[string]string `json:"name_localizations"`
  Value string `json:"value"` // Can be string, int or double
}
