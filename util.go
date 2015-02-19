package gocluster

import (
  "crypto/sha1"
  "io"
  "encoding/hex"
)

func ToSha1(content string) string {
  h := sha1.New()
  io.WriteString(h, content)

  hash := hex.EncodeToString(h.Sum(nil))
  return hash
}