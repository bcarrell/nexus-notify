nexus-notify
============

This is a Nexus 5 notifier written in Go.  It'll email you when the Nexus 5 is in the Play Store.

If you'd like to use this, add a `vars.go` file that looks like this:

```
package main

const (
	ApiKey       string = "Mandrill API key"
	EmailAddress string = "destination_email"
	EmailName    string = "destination_name"
)
```

Then `go build` and run the binary as you please.
