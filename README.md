# chatgpt
A tiny ChatGPT3 CLI. It uses the most expensive model with maximal token count,
so one query can cost up to $0.08 (as of Jan 2023).

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
enter to submit & "q" to stop
> What's the capital of China?
> Beijing is the capital of China.
> q
```