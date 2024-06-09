# Twitter-like Application Backend

This is a backend implementation for a Twitter-like application built using Go Fiber and PostgreSQL.

## Features

- User authentication (registration and login)
- Tweet management (creation and retrieval)
- Follows: Users can follow and unfollow other users
- Likes: Users can like and unlike tweets
- Retweets: Users can retweet and unretweet tweets

## Prerequisites

Before running the application, make sure you have the following installed:

- Go: [Installation Guide](https://golang.org/doc/install)
- PostgreSQL: [Installation Guide](https://www.postgresql.org/download/)
- Go Fiber: `go get -u github.com/gofiber/fiber/v2`
- PostgreSQL driver for Go: `go get github.com/jackc/pgx/v4/pgxpool`
- Environment configuration: `go get github.com/spf13/viper`

## Setup

1. Clone this repository:

   ```bash
   git clone <repository_url>
   cd twitter-backend
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the PostgreSQL database and configure the connection string in the `.env` file.

4. Run the database migrations:

   ```bash
   go run github.com/golang-migrate/migrate/v4 -path database/migrations -database "<your_database_url>" up
   ```

5. Start the server:

   ```bash
   go run main.go
   ```

## API Endpoints

### Authentication

- `POST /api/register`: Register a new user.
- `POST /api/login`: Login with existing credentials.

### Users

- `PATCH /api/user/change-username`: Change the username of the authenticated user.
- `POST /api/follow/:id`: Follow a user.
- `DELETE /api/unfollow/:id`: Unfollow a user.

### Tweets

- `POST /api/tweets`: Create a new tweet.
- `GET /api/tweets`: Get all tweets.

### Likes

- `POST /api/tweets/:id/like`: Like a tweet.
- `DELETE /api/tweets/:id/unlike`: Unlike a tweet.

### Retweets

- `POST /api/tweets/:id/retweet`: Retweet a tweet.
- `DELETE /api/tweets/:id/unretweet`: Remove a retweet.

## Middleware

- `AuthRequired()`: Authentication middleware to protect routes.

## Configuration

You can configure the application using environment variables. Create a `.env` file in the root directory with the following variables:

```
DB_URL=postgres://username:password@localhost:5432/twitter?sslmode=disable
JWT_SECRET=your_jwt_secret_key
```

## Contributing

Contributions are welcome! If you find any issues or want to add new features, feel free to open a pull request.




