# DEVSETUP

Under construction :warning:

## Requirements

For compile

- node 16
- golang 19

For format and linting

- [eslint](https://eslint.org/) + [prettier](https://prettier.io/)
- [go for VSCode](https://marketplace.visualstudio.com/items?itemName=golang.Go) + [golangci-lint](https://github.com/golangci/golangci-lint)

## VS Code setup

Recommended `settings.json`.

```json
{
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "eslint.validate": ["typescript"],
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"]
}
```
