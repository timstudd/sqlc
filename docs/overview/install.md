# Installing sqlc

sqlc is distributed as a single binary with zero dependencies.

## macOS

```
brew install sqlc
```

## Ubuntu

```
sudo snap install sqlc
```

## go install

### Go >= 1.17:

```
go install github.com/timstudd/sqlc/cmd/sqlc@latest
```

### Go < 1.17:

```
go get github.com/timstudd/sqlc/cmd/sqlc
```

## Docker

```
docker pull kjconroy/sqlc
```

Run `sqlc` using `docker run`:

```
docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
```

## Downloads

Get pre-built binaries for _v1.10.0_:

- [Linux](https://github.com/timstudd/sqlc/releases/download/v1.10.0/sqlc_1.10.0_linux_amd64.tar.gz)
- [macOS](https://github.com/timstudd/sqlc/releases/download/v1.10.0/sqlc_1.10.0_darwin_amd64.zip)
- [Windows (MySQL only)](https://github.com/timstudd/sqlc/releases/download/v1.10.0/sqlc_1.10.0_windows_amd64.zip)
