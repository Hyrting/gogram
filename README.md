# Gogram

**Gogram** is a Go framework for building Telegram bots, forked from [telebot](https://github.com/tucnak/telebot). This project aims to extend and improve the original package with new features and optimizations, while still staying up-to-date with upstream changes.

## Features

- **SendAfter**: Schedule a message to be sent after a specified duration.
- **SendAfterSignal**: Wait for a message from a channel and send it to the specified recipient when received.
- **DeleteAfter**: Schedule a delete operation for a message after a specified duration.
- **DeleteAfterSignal**: Wait for a signal from a channel and delete the message when the signal is received.
- **EditAfter**: Schedule an edit operation for a message after a specified duration.
- **EditAfterSignal**: Wait for a message from a channel and edit the specified message when received.
- **EditAfterSignalWithTimeout**: Wait for a message from a channel and edit the specified message if received within a timeout; otherwise, perform a different action.
- **Built-in Goroutine Pool**: Efficiently handles multiple updates concurrently by utilizing goroutines.
- **Command Argument Parsing**: The `ParseArgs` command parses all arguments based on the variables passed into it.
- **Rate Limiter (Coming Soon)**: A rate limiter to comply with Telegram's API rate limits, ensuring smooth request handling during high traffic.

## Installation

To install **Gogram**, you can use `go get` to fetch it directly from GitHub:

```bash
go get github.com/hyrting/gogram
```

After installation, you can import the package in your Go project:

```go
import "github.com/hyrting/gogram"
```

## Contributions and Suggestions

I'm open to all feature recommendations, feedback, and suggestions! Feel free to submit ideas or contribute via pull requests. The goal is to continue improving gogram while making it flexible for various bot-building needs.

## Donations

If you'd like to support the project, donations are greatly appreciated:

- **Bitcoin**: `bc1qyxvfvv2c0g0ydhchaa9rqk63x0hz8enxc74kw0`
- **Ethereum**: `0xBabee16687eE49C3CB88Ce939431653dEc56A788`
- **Litecoin**: `LSqrE8Tm5hNJuimkzpMWmgiMoaP8fpKEfd`
- **Ton**: `UQBnrFlTTEr0ezoJyxqsXSTdoraKnp-AGI82yFYDFC4Hj36b`
- **Monero**: `44cCgBhwKzY9Wd9Wme5cbdRQSxhC5SKpLCvZoNnDCzv8PakVVSMbSxCiP5u6eSuwZ6P29pJqrHA3RAkmDizPv2LC14PgczX`
