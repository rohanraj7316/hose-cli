# Project Title

url shortener it helps to shorten any given url

## Getting Started

below are the steps to make your project up and running

1. `mv start.sh.sample start.sh`

2. `go download`

3. `sh start.sh`

### Prerequisites

you need `go 1.18+`

### API

#### Create

used to create shorted urls by using long urls

**Path**: `/`

**Method**: `POST`

**Request**:

```json
{
  "originalUrl": "https://stackoverflow.com/questions/77551163/in-linux-evnrionment-differentiate-between-between-project-relative-path-and-fiy"
}
```

**Response**:

```json
{
  "redirectionUrl": "http://localhost:8080/J2tb4"
}
```

#### Get

used for redirection

**Path**: `/`

**Method**: `GET`

**Request**: `NA`

**Response**: `status code 308`

#### Top 3 Shorted Domains

used to generate analytics of top 3 domains that's been shorted

**Path**: `/top-3-shorted-domains`

**Method**: `GET`

**Request**: `NA`

**Response**:

```json
{
  "topThreeShortedDomains": {
    "stackoverflow.com": 2,
    "stackoverflow1.com": 2,
    "stackoverflow2.com": 2
  }
}
```
