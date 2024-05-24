package apimodels

type CreateMessageInput struct {
  Content string `json:"content"`
  Nonce string `json:"nonce"`
  Tts bool `json:"tts"`
  // Embeds string `json:"embeds"`
  // AllowedMentions string `json:"allowed_mentions"`
  // MessageReference string `json:"message_reference"`
  // Components string `json:"components"`
  StickerIds []string `json:"sticker_ids"`
  // Files string `json:"files"`
  PayloadJson string `json:"payload_json"`
  // Attachments string `json:"attachments"`
  Flags []int `json:"flags"`
  EnforceNonce bool `json:"enforce_nonce"`
  // Poll string `json:"poll"`
}
