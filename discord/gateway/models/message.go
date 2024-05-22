package gatewaymodels

type Message struct {
  D	Snowflake `json:"d"`
  ChannelId	Snowflake `json:"channel_id"`
  // Author string `json:"author"`
  Content string `json:"content"`
  Timestamp	string `json:"timestamp"`
  EditedTimestamp	string `json:"edited_timestamp"`
  Tts	bool `json:"tts"`
  MentionEveryone	bool `json:"mention_everyone"`
  // Mentions string `json:"mentions"`
  // MentionRoles string `json:"mention_roles"`
  // MentionChannels string `json:"mention_channels"`
  // Attachments string `json:"attachments"`
  // Embeds string `json:"embeds"`
  // Reactions string `json:"reactions"`
  Nonce string `json:"nonce"`
  Pinned bool `json:"pinned"`
  WebhookId Snowflake `json:"webhook_id"`
  Type int `json:"type"`
  // Activity string `json:"activity"`
  // Application string `json:"application"`
  ApplicationId Snowflake `json:"application_id"`
  // MessageReference string `json:"message_reference"`
  Flags int `json:"flags"`
  // ReferencedMessage string `json:"referenced_message"`
  // InteractionMetadata string `json:"interaction_metadata"`
  // Interaction string `json:"interaction"`
  // Thread string `json:"thread"`
  // Components string `json:"components"`
  // StickerItems string `json:"sticker_items"`
  // Stickers string `json:"stickers"`
  Position int `json:"position"`
  // RoleSubscriptionData string `json:"role_subscription_data"`
  // Resolved string `json:"resolved"`
  // Poll string `json:"poll"`
  // Call string `json:"call"`
}
