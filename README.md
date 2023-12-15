# Orange Squad LiveNet Token Generator API

The Orange Squad LiveNet Token Generator API is the token generation and management API developed specifically for the Orange Squad Live Networking. It provides functionality for creating and managing authentication tokens for users associated with their user IDs (e.g. Discord IDs, which work as their unique identifiers).

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [GET Token](#get-token)
  - [SET Token](#set-token)
  - [REWRITE Token](#rewrite-token)
  - [DELETE Token](#delete-token)
- [License](#license)

## Installation

To use the Orange Squad LiveNet Token Generator API, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/orange-squad-livenet-token-generator.git
   ```

2. Navigate to the project directory:

   ```bash
   cd orange-squad-livenet-token-generator
   ```

3. Run the API:

   ```bash
   go run main.go
   ```

   The server will start listening on port 8080.

## Usage

The API exposes several endpoints to manage tokens associated with user IDs.

### GET Token

Retrieve a token for a given user ID.

```http
GET /get/{id}
```

### SET Token

Generate and set a new token for a given user ID.

```http
GET /set/{id}
```

### REWRITE Token

Re-generate and set a new token for a given user ID.

```http
GET /rewrite/{id}
```

### DELETE Token

Delete the token associated with a given user ID.

```http
GET /delete/{id}
```

All methods use GET so that they can be easily previewed in a browser.

### Error Responses

If an error occurs during the API operations, an error response in JSON format will be returned, containing an "error" field.

```json
{
  "error": "Error message details"
}
```

### Token Responses

Normally, the `/set`, `/get`, and `/rewrite` routes return tokens of the provided ID. In that case, the API will return a JSON response containing a "token" field.

```json
{
  "token": "Token value"
}
```

### Message Responses

The `/delete` route returns a message indicating whether the token was successfully deleted or not. In that case, the API will return a JSON response containing a "message" field.

```json
{
  "message": "Message details"
}
```

## License

This project is licensed under the BSD-3 Clause License - see the [LICENSE](LICENSE) file for details.