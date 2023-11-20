
# PokeAPI Go Proxy

The PokeAPI Go Proxy is a Golang-based service that acts as a proxy for the [PokeAPI](https://pokeapi.co/), caching responses to enhance performance and parsing them to remove direct links to the PokeAPI main service.

## Table of Contents

- [PokeAPI Go Proxy](#pokeapi-go-proxy)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
    - [Endpoint](#endpoint)
    - [Caching](#caching)
  - [Configuration](#configuration)
  - [Error Handling](#error-handling)
  - [Security](#security)
  - [Logging](#logging)
  - [Testing](#testing)
  - [Contributing](#contributing)
  - [License](#license)

## Features

- Caches responses from the PokeAPI to improve performance.
- Parses responses to remove direct links to the PokeAPI main service.
- Error handling for robustness.
- Security measures to protect against vulnerabilities.
- Logging for monitoring and debugging.

## Getting Started

### Prerequisites

- Go (Golang)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ozancakir/pokeapi-go-proxy.git
   cd pokeapi-go-proxy
   ```

2. Run the project:

   ```bash
   go run cmd/main.go
   ```


## Usage

### Endpoint

The proxy exposes all endpoints that can be found in pokeapi v2

```
GET /api/pokemon
```

- It will get results from https://pokeapi.co/api/v2/pokemon while caching the content. 


### Caching

All responses from destination endipint will be cached using sqlite in pokeapi.db

## Configuration

- Modify the environment file (rename `.example.env` to `.env`) to customize urls
- POKEAPI_URL -> pokeapi url for proxy destination // default:[POKEAPI_URL](https://pokeapi.co/api/v2)
- API_PREFIX -> api prefix for your services // default: /api
- PORT -> desired application port
- GIN_MODE -> GIN framework running mode // default: debug [debug,release,test]
- API_KEY -> X-API-KEY header check value, if not provided api will be public

## Error Handling

The proxy handles errors gracefully and provides informative error messages to clients.

## Security

- The proxy is designed with security in mind. If exposed publicly, consider implementing authentication mechanisms.

## Logging

- Logging is implemented to keep track of incoming requests, cache hits, cache misses, and any errors. Logs are available in the `logs` directory.

## Testing

- Thoroughly test the proxy under various scenarios, including different types of requests, error conditions, and high traffic loads.

## Contributing

Contributions are welcome! Please follow our [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize this template further based on your project's specific details and requirements.