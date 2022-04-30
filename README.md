# ghost-external-links
Ghost CMS: List all external links (href) in posts and articles

## Usage

Tokens can be created in Ghost Admin Panel > Settings > Integrations > Custom Integration

```go
GHOST_ADMIN_KEY="abcde:fgeh" GHOST_URL="https://example.com" GHOST_CONTENT_KEY="abcde" go run cmd/ghost-external-links/main.go 
[...]

2 | https://www.financeads.net/tc.php?t=25796C144333224B
Source: https://etf.capital/robo-advisors-fur-junge-anleger/

2 | https://de.wikipedia.org/wiki/Mt.Gox
Source: https://etf.capital/krypto-boerse-im-vergleich-kosten-und-besonderheiten/

[...]
```

Output format:

```
Count of each external link | External URL
Source: Page/Post URL
```
