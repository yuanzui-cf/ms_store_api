# Macrohard Store API

## Overview

The Macrohard Store API project provides a backend service to generate download URLs for UWP apps.

## Requirement

`Rust >= 1.88.0`

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yuanzui-cf/ms_store_api.git
cd ms_store_api
```

2. Build the project:

```bash
cargo build --release
```

3. Run the server:

```bash
./target/release/msapi
```

## Usage

- **Endpoint**: `/{product_id}`
- **Method**: `GET`
- **Description**: Fetches app urls for the given Microsoft Store product ID.

### Example Request

```bash
curl http://localhost:9000/9NBLGGH2JHXJ
```

## Configuration

The server can be configured using:

- **Environment Variables**: Prefix `MSAPI_`.
- **Config File**: `config.toml`.

### Default Configuration

- **App Name**: Macrohard Store API
- **Address**: `0.0.0.0`
- **Port**: `9000`

## Docker Support

### Build Docker Image

```bash
docker build -t ms_store_api .
```

### Run Docker Container

```bash
docker run -p 9000:9000 ms_store_api
```

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request.

## License

This project is licensed under the MIT License.

## Attribution

Parts of the code for generating app URLs were adapted from https://github.com/mjishnu/alt-app-installer-cli.
