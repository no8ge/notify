package util

import (
	"testing"

	types "github.com/no8geo/notify/pkg"
)

func TestWechatSendMarkdown(t *testing.T) {

	data := &types.MarkdownMsg{
		Msgtype: "markdown",
		Markdown: map[string]string{
			"content": "123",
		},
	}

	w := WeChat{
		Key: "5311360b-bfe0-461f-a075-d3c5434ec804",
	}
	err := w.SendMarkdown(data)
	if err != nil {
		t.Errorf("Failed to send the Markdown message: %v", err)
		return
	}

}
