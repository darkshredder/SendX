# SendX

## Installation

```bash
go mod tidy
go run .
```

## Usage

```bash
curl --location --request POST 'http://localhost:7771/pagesource' \
--header 'Content-Type: application/json' \
--data-raw '{
   "uri": "https://google.com",
   "retryLimit": 3
}'
```

## Assumptions

- The application is a simple web server that accepts a POST request with a JSON body containing a URI and a retry limit.
- Worker pool is used to process the for downloading page Source.
- Queue is used to store the requests.
- If file is not downloaded successfully. The source uri will fallback to not found page.