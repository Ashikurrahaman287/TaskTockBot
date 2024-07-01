# TaskTockBot

TaskTockBot is a Telegram bot written in Go that helps you manage your tasks. You can add tasks, mark them as done, and receive periodic reminders for pending tasks.

## Features

- **Add Tasks**: Use `/add <task>` to add tasks to your list.
- **Mark Tasks as Done**: Use `/done <task>` to mark tasks as done.
- **Periodic Reminders**: Receive reminders every hour for pending tasks.

## Setup

To run TaskTockBot locally or deploy it to a server, follow these steps:

### Prerequisites

- Go (version 1.16+)
- `go-telegram-bot-api` library (`github.com/go-telegram-bot-api/telegram-bot-api/v5`)
- Telegram Bot Token

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Ashikurrahaman287/TaskTockBot.git
   cd TaskTockBot
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your Telegram Bot Token:
   
   - Create a new Telegram bot and obtain the API token.
   - Set the `botToken` variable in `main.go` with your bot token.

4. Build and run the bot:

   ```bash
   go build -o TaskTockBot main.go
   ./TaskTockBot
   ```

### Usage

- Start the bot by sending `/start`.
- Add a task using `/add <task>`.
- Mark a task as done using `/done <task>`.

## Contributing

Contributions are welcome! If you have ideas for new features, find bugs, or want to improve the code, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

### Instructions for Customization:

- **Features**: Describe the bot's capabilities briefly under the "Features" section.
- **Setup**: Provide clear steps to set up the bot locally or on a server.
- **Usage**: Explain how to interact with the bot and use its commands.
- **Contributing**: Encourage others to contribute to the project and specify how they can do so.
- **License**: Include information about the project's licensing terms.
