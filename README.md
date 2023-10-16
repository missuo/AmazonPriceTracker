# Amazon Price Tracker
Amazon Price Tracker is a web application that allows you to track the prices of products on Amazon. 

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API Endpoint](#api-endpoint)
- [Contributing](#contributing)
- [License](#license)

## Features

- Fetches and tracks the prices of products on Amazon.
- Supports both new and used prices.
- Provides a simple API to retrieve price information.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed on your local machine.
- [Gin](https://github.com/gin-gonic/gin) framework for web application.

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/missuo/AmazonPriceTracker
   cd AmazonPriceTracker
   ```

2. Build and run the application:

   ```sh
   go build main.go
   ./main
   ```

## Usage

Once the application is running, you can access it through a web browser. Here are the available routes:

- `/`: Home route that displays a welcome message.

- `/price`: Price tracking route. You can provide a product link or ID as a query parameter to get the price information.

Example API request:
```
GET /price?link=https://www.amazon.com/dp/<product-id>
```

## API Endpoint

The API endpoint for retrieving product price information is:

```
GET /price?link=<Amazon-product-link>&id=<Amazon-product-id>
```

- `link`: The Amazon product link (e.g., `https://www.amazon.com/dp/<product-id>`).
- `id`: (Optional) The Amazon product ID. If provided, it will be used instead of extracting the ID from the link.

## Contributing

Contributions are welcome! If you would like to contribute to this project, please open an issue or create a pull request.

## License

This project is licensed under the [Apache License](./LICENSE).