# Server Notify

This Go application checks if a webpage is working by sending an HTTP GET request and checking for a `200 OK` status. It runs in an infinite loop, checking the status every `TIME_PING` minutes. If the page is down (doesn't return a `200 OK`), the application will notify you via either email or Telegram, depending on the first argument passed (`mail` or `telegram`).

## Features

- Infinite loop that checks if a webpage returns a `200 OK` status.
- Notifies via email or Telegram if the webpage is down.
- Configurable check interval using environment variables.

## Setup

To make the application work, you need to configure the following environment variables in a `.env` file.

### Environment Variables

#### General Configuration

- `TIME_PING`: The interval in minutes for checking the webpage.
- `SITES`: The webpage(s) to monitor, separated by commas (e.g., `https://example.com,https://another-site.com`).

#### Email Configuration (if using email notification)

- `MAIL_FROM`: Your email address.
- `MAIL_PASSWORD`: Your email account's password or app-specific password.
- `MAIL_TO`: The recipient email address to notify.

#### Telegram Configuration (if using Telegram notification)

- `TELEGRAM_TOKEN`: Your Telegram bot token.
- `TELEGRAM_CHAT_ID`: The chat ID to send notifications to.

## Usage

You can run the application either by building it with Go or by using Docker.

### Option 1: Run with Docker

1. Make sure you have a `.env` file with all necessary environment variables filled out.
2. Build the Docker image:

   ```bash
   docker build -t webpage-status-checker .
   ```
3. Run the application using Docker, passing either `mail` or `telegram` as the first argument:
    ```bash
    docker run --env-file .env webpage-status-checker mail
    ````
    or
    
    ```bash
    docker run --env-file .env webpage-status-checker telegram
    ```
### Option 2: Run Locally with Go

1. Make sure Go 1.22.1 or higher is installed on your machine.
2. Fill out the `.env` file with the required environment variables.
3. Run the application using Go:
    ```bash
    go run main.go  mail
    ````
    or
    
    ```bash
    go run main.go mail telegram
    ```

## Example

If you want to check the status of `https://example.com` every 5 minutes and receive notifications by email:

1. Set the following values in your `.env` file:
    ```bash
    TIME_PING=5
    SITES=https://example.com
    MAIL_FROM=your_email@example.com
    MAIL_PASSWORD=your_password
    MAIL_TO=recipient@example.com
    ```
2. Run the application with Docker:
    ```bash
    docker run --env-file .env webpage-status-checker mail
    ```
    Or run it locally with Go:

    ```bash
    go run main.go mail
    ```

## Requirements

- Go 1.22.1 or higher (if running locally).
- Docker (if running with Docker).
- Correctly configured environment variables.

## License

This project is licensed under the MIT License.
```vbnet
This `README.md` provides clear instructions on how to configure, build, and run the project, either using Docker or directly with Go, and also explains the required environment variables.
```