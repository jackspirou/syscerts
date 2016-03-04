# syscerts
Gather local system certificates in Go via a public `SystemRootsPool` method.

### This package does one thing:
Provide a way to gather local system certificates
on multiple OS platforms.

### How does it do it?:
It uses the `crypto/x509` package and provides a single public method called
`SystemRootsPool()` to return a `*x509.CertPool` object.

### How do you use it?:
```Go
// gather CA certs
certpool := syscerts.SystemRootsPool()

// place them in an HTTP client for trusted SSL/TLS connections
tlsConfig := &tls.Config{RootCAs: certpool}
transport := &http.Transport{TLSClientConfig: tlsConfig}
client := &http.Client{Transport: transport}

// make a request
resp, err := client.Do(req)
```

### Why even do this?:
As it turns out, there is already a method in the `crypto/x509` named
`systemRootsPool`.
The `crypto/x509.systemRootsPool` method is the same as
`github.com/jackspirou/syscerts.SystemRootsPool` except that the `crypto/x509`
package keeps its `systemRootsPool` method private.

When `crypto/x509.systemRootsPool` becomes a public method in Go 1.7 this
package will no longer be needed.

You can find more at this Go issue: https://github.com/golang/go/issues/13335
