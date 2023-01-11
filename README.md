# chatgpt
A tiny ChapGPT3 CLI.

Install or update:

```text
go install github.com/clfs/chatgpt@latest
```

Uninstall:

```text
rm -i $(which chatgpt)
```

Usage:

```text
$ export GPT3_KEY="sk-..." # this should go in, e.g., ~/.zshenv
$ chatgpt
type "exit", "quit", or "q" to exit; hit enter to submit
> What's the capital of China?
> Beijing is the capital of China.
> q
```