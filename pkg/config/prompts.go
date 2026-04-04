package config

import (
	"strings"
)

func wrapBlockCode(t string, code string) string {
	return "```" + t + "\n" + code + "\n```"
}

var COMMIT_PROMPT = strings.ReplaceAll(
	COMMIT_PROMPT_WITH_CUSTOM_TONE,
	"{{tone}}",
	"Use a LOT of emojis, be funny, and expressive. Feel free to be profane, but don't be offensive",
)

var COMMIT_PROMPT_WITH_CUSTOM_TONE = wrapBlockCode("diff", "{{diff}}") + "\n\n" + `
Write short commit messages:
- Important: Follow the conventional commit format
- The first line should be a short summary of the changes
- Remember to mention the files that were changed, and what was changed
- Explain the 'why' behind changes
- Use bullet points for multiple changes
- Tone: {{tone}}
- If there are no changes, or the input is blank - then return a blank string

Think carefully before you write your commit message.

The output format should be:

<Summary of changes>
- <changes>
- <changes>

What you write will be passed directly to git commit -m "[message]"
`
