# Project Title

CLI for encryption and decryption

## Installation

```bash
curl -O https://raw.githubusercontent.com/rohanraj7316/hose-cli/refs/heads/main/install.sh && chmod +x install.sh && ./install.sh
```

## Usage

### Encrypt

```bash
hose encrypt -p "Rohan Raj" -s "15760b7b91427cb951011634a426e3c7" -k "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqsalSBFQaNMCeMrEGkhDfRRfHJAAGWEu+sx+DuOKeXIB21AgbNxvv0Qm3jxVUPlRbr0wCLs+tsA67oj2dx6GNFoRznT9fEKuBvXHzqiDejjP5HmgqFgVnJgXH+2++1VUtuRcU6fHtZoWddvnlDKL3RGLLDl13ObVgsrG2nlC2a+++xvdavASnaz6TbbqLbn511U+05nnkX+vuso5GGYAMhqUf0QyDAiR0BEgZy2VX4MBngKfYpvIRwNNog7DQvm4OH9524PLz0rfxlkZT0xC403kPqd9sNHHdvJ4qnjHlPQG6aQQkAR6Potk67mGWNyDvctobPppTUsF2BYpCMhPEQIDAQAB" --json
```

output:

```json
{"api_encryption_key":"YrbnkevvI+sSFr2CaUHbhvxgpaYkcwPZszWlxtVXlZvY1G++YyBahxxXzuxbknwxNfsO6nxoH4hiv60X9Syy3lXJoz7cLNP33YHI5CSaNdm/XvXOy7c/CtcSBbpDFkRAd9Ff4cjTLnPCdbKSk480m5J0rKuKYk9ljK42f1zre2KqHBhv5wU4xF0WB45duX+t/qGeNUYq28D9sKc6qmwUi6/L3OKTV/G8BJSK94zLIk4ROh8QBItqUXkEplpjRFZscDcSUMMj5r2qvLR3s6SNHG+5FvFhljvLogVwNfZREJEJJ2jg139KasMZ00dXVzRkHI6NCkYukIdmS9ce/CZaBQ==","encrypted_payload":"WOu8OygUp+aB1tE9AZq8nA==.e12zAdbW8mMDPW8HrVmjLg==.TUTAAyidksSB"}
```

### Decrypt

```bash
 hose decrypt -e "7nOo+Z6MJFyAVe7SsKXLzg==.gr4ZHNGuf/7tOT8qF3mfsQ==.Ix12QGPZoVI4" -s "15760b7b91427cb951011634a426e3c7" --json
```

output:

```json
{"decrypted_payload":"Rohan Raj"}
```
