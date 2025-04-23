# 🌐 http-tester

A simple command-line HTTP client for testing RESTful APIs. Supports `GET`, `POST`, `PUT`, and `DELETE` methods with JSON payloads.

Built with Go and [`urfave/cli`](https://github.com/urfave/cli).

## 🛠️ Features

- 🔍 `GET` requests to fetch resources
- ✉️ `POST` and `PUT` requests with JSON bodies
- ❌ `DELETE` requests to remove resources
- 📦 Standalone binary with no runtime dependencies

## 🏃‍♂️ Usage

You can run the tool directly with `go run` or use the compiled binary `http-tester`.

### GET

```bash
http-tester get <URL>
```

**Example:**
```bash
http-tester get https://jsonplaceholder.typicode.com/posts/1
```

### POST

Supports two options to provide the JSON payload:

- **`--data`**: raw JSON string (escape double quotes in PowerShell)
- **`--data-file`**: path to a file with JSON content

```bash
# Using --data (bash/zsh)
http-tester post --data '{"title":"foo","body":"bar","userId":1}' https://jsonplaceholder.typicode.com/posts

# Using --data-file
http-tester post --data-file data.json https://jsonplaceholder.typicode.com/posts
```

### PUT

Same flags as POST for JSON payload:

```bash
# Using --data
http-tester put --data '{"title":"updated","body":"baz","userId":1}' https://jsonplaceholder.typicode.com/posts/1

# Using --data-file
http-tester put --data-file update.json https://jsonplaceholder.typicode.com/posts/1
```

### DELETE

```bash
http-tester delete <URL>
```

**Example:**
```bash
http-tester delete https://jsonplaceholder.typicode.com/posts/1
```

---

## 📄 Examples

**Working POST request in PowerShell:**

```powershell
go run main.go post --data '{"title":"foo","body":"bar","userId":1}' https://jsonplaceholder.typicode.com/posts
```

Sample output:
```text
Status: 201 Created

{"id":101,"title":"foo","body":"bar","userId":1}
```

---

## ⚙️ Configuration Options

- `--data` (string): raw JSON data to send in POST/PUT.
- `--data-file` (string): read JSON payload from file.
- Either `--data` or `--data-file` must be provided for POST/PUT.

---

## 🧪 Dependencies

- [Go standard library](https://golang.org/pkg/)
- [urfave/cli v2](https://github.com/urfave/cli)

