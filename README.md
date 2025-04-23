# 🌐 http-tester

A simple command-line HTTP client for testing RESTful APIs. Supports `GET`, `POST`, `PUT`, and `DELETE` methods with JSON payloads.

Built with Go and [`urfave/cli`](https://github.com/urfave/cli).

## 🛠️ Features

- 🔍 `GET` requests to fetch resources
- ✉️ `POST` and `PUT` requests with JSON bodies
- ❌ `DELETE` requests to remove resources
- 📦 Standalone binary with no runtime dependencies

## 📦 Usage

### `GET`

```bash
http-tester get https://jsonplaceholder.typicode.com/posts/1
```

### `POST`

```bash
http-tester post https://jsonplaceholder.typicode.com/posts --data '{"title":"foo", "body":"bar", "userId":1}'
```

### `PUT`

```bash
http-tester put https://jsonplaceholder.typicode.com/posts/1 --data '{"id":1,"title":"updated","body":"new body","userId":1}'
```

### `DELETE`

```bash
http-tester delete https://jsonplaceholder.typicode.com/posts/1
```

---

## 🧩 Example Output

```bash
Status: 200 OK

{
  "id": 1,
  "title": "foo",
  "body": "bar",
  "userId": 1
}
```

---

## 🧪 Dependencies

- [Go standard library](https://golang.org/pkg/)
- [urfave/cli v2](https://github.com/urfave/cli)
