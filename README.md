# VnExpress RSS API

1. Fetch RSS From VNExpress and stored in DB every minute
2. Serve Rest API

## Main Features

- CRUD RSS Feed Link(link + name)
- Get RSS Posts with/without filter( filter by rss feed link)
- Cron Job

## Deployment

- Containerize the api
- Deploy in AWS

## Data structure

```xml
<rss version="2.0">
<channel>
<title>Thế giới - VnExpress RSS</title>
<description>VnExpress RSS</description>
<image>
<url>https://s.vnecdn.net/vnexpress/i/v20/logos/vne_logo_rss.png</url>
<title>Tin nhanh VnExpress - Đọc báo, tin tức online 24h</title>
<link>https://vnexpress.net</link>
</image>
<pubDate>Thu, 06 Feb 2025 20:45:46 +0700</pubDate>
<generator>VnExpress</generator>
<link>https://vnexpress.net/rss/the-gioi.rss</link>
<item>
<title>Israel nối gót Mỹ rời Hội đồng Nhân quyền LHQ</title>
<description>
<![CDATA[ <a href="https://vnexpress.net/israel-noi-got-my-roi-hoi-dong-nhan-quyen-lhq-4846612.html"><img src="https://i1-vnexpress.vnecdn.net/2025/02/06/afp-20250123-36vm3wf-v1-highre-6597-7613-1738846835.jpg?w=1200&h=0&q=100&dpr=1&fit=crop&s=EZLAecVUZhLBSg818oeXbA"></a></br>Ngoại trưởng Israel cho biết nước này sẽ rút khỏi Hội đồng Nhân quyền Liên Hợp Quốc, hưởng ứng quyết định tương tự của Mỹ. ]]>
</description>
<pubDate>Thu, 06 Feb 2025 20:30:48 +0700</pubDate>
<link>https://vnexpress.net/israel-noi-got-my-roi-hoi-dong-nhan-quyen-lhq-4846612.html</link>
<guid>https://vnexpress.net/israel-noi-got-my-roi-hoi-dong-nhan-quyen-lhq-4846612.html</guid>
<enclosure type="image/jpeg" length="1200" url="https://i1-vnexpress.vnecdn.net/2025/02/06/afp-20250123-36vm3wf-v1-highre-6597-7613-1738846835.jpg?w=1200&h=0&q=100&dpr=1&fit=crop&s=EZLAecVUZhLBSg818oeXbA"/>
</item>
```

## Development

- [Live reload go apps with air](https://github.com/air-verse/air)
- [Go linter with golangci-lint](https://golangci-lint.run/welcome/install/)
- [Gin gonic document](https://gin-gonic.com/docs/)

```sh
go env GOPATH
ls $(go env GOPATH)/bin
#  Add Go Bin Directory to PATH
# ~/.zshrc
export PATH="$(go env GOPATH)/bin:$PATH"
# init
air init
```

### Linter

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
golangci-lint
```

### Testing

```sh
go install github.com/stretchr/testify
```

```go
// Standard testing
func TestWithoutTestify(t *testing.T) {
    result := SomeFunction()
    if result != expectedValue {
        t.Errorf("Expected %v, got %v", expectedValue, result)
    }
}

// With testify
func TestWithTestify(t *testing.T) {
    result := SomeFunction()
    assert.Equal(t, expectedValue, result, "they should be equal")
}
```

### Api Doc

- [Gin Swagger](https://pkg.go.dev/github.com/swaggo/gin-swagger@v1.6.0#section-readme)

How to update api doc

```sh
go generate ./..
make docs_generate
```
