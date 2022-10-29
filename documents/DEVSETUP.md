# DEVSETUP

Under construction :warning:

## Requirements

To actually run something on your machine you will need to have installed:

- Node 16
- Golang 19

Then, to keep the code up to the standards enforced, there are also tools for each of the areas:

**Front-end**

[ESlint](https://eslint.org/) and [Prettier](https://prettier.io/)

**Back-end**:

[golangci-lint](https://github.com/golangci/golangci-lint) creates a good combination to plugin in a couple of linters that we enforce, such as:

- [revive](https://github.com/mgechev/revive)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [misspell](https://github.com/client9/misspell/tree/v0.3.4)
- [gocritic](https://github.com/go-critic/go-critic)

The Go linters are currently configured in `.golangci.yml`, per [golangci-lint configuration documentation](https://golangci-lint.run/usage/configuration). Feel free to suggest rules that should also be enforced, if you want inspiration check out [revive rule list](https://github.com/mgechev/revive/blob/master/RULES_DESCRIPTIONS.md).

## VS Code setup

For the front-end, the recommended `settings.json` is the following:

```json
{
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "eslint.validate": ["typescript"]
}
```

For the back-end, VSCode should be configured using the [Go for VSCode](https://marketplace.visualstudio.com/items?itemName=golang.Go) along with the aforementioned linters installed as binaries (or go binaries) in the system - the configuration file `.golangci.yml` will pick them up for you and enforce it in your code.
